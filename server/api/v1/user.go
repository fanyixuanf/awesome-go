/**
 * @Author: Yuxue.Fan<yuxue.fan@yidatec.com>
 * @Description:
 * @File: user.go
 * @Version: 1.0.0
 * @Date: 2022/2/17 16:01
 */

package v1

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/global/response"
	"gin-vue-admin/middleware"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"time"
)

// Login
//@Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.RegisterAndLoginStruct true "用户名, 密码, 验证码"
// @Success 200 {object} response.Response{data=resp.LoginResponse,msg=string} "返回包括用户信息,token,过期时间"
// @Router /base/login [post]
func Login(c *gin.Context) {
	var L request.RegisterAndLoginStruct
	_ = c.ShouldBindJSON(&L)
	UserVerify := utils.Rules{
		"CaptchaId": {utils.NotEmpty()},
		"Captcha":   {utils.NotEmpty()},
		"Username":  {utils.NotEmpty()},
		"Password":  {utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(L, UserVerify)
	if UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		global.GVA_LOG.Error("数据验证失败, err:", zap.Any("err", UserVerifyErr))
		return
	}
	if store.Verify(L.CaptchaId, L.Captcha, true) {
		U := &model.SysUser{Username: L.Username, Password: L.Password}
		if err, user := service.Login(U); err != nil {
			response.FailWithMessage(fmt.Sprintf("用户名密码错误或%v", err), c)
		} else {
			tokenNext(c, *user)
		}
	} else {
		response.FailWithMessage("验证码错误", c)
	}
}

// 登录以后签发jwt
func tokenNext(c *gin.Context, user model.SysUser) {
	j := &middleware.JWT{
		SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey), // 唯一签名
	}
	clams := request.CustomClaims{
		UUID:        uuid.UUID(user.UUID),
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
		BufferTime:  60 * 60 * 24, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,       // 签名生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*7, // 过期时间 7天
			Issuer:    "Tools",                        // 签名的发行者
		},
	}
	token, err := j.CreateToken(clams)
	if err != nil {
		response.FailWithMessage("获取token失败", c)
		global.GVA_LOG.Error("获取token失败, err:", zap.Any("err", err))
		return
	}
	if !global.GVA_CONFIG.System.UseMultipoint {
		response.OkWithData(resp.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: clams.StandardClaims.ExpiresAt * 1000,
		}, c)
		return
	}
	err, jwtStr := service.GetRedisJWT(user.Username)
	if err == redis.Nil {
		if err := service.SetRedisJWT(token, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			global.GVA_LOG.Error("设置登录状态失败, err:", zap.Any("err", err))
			return
		}
		if err := service.SetRedisJWT(time.Now().Format("2006-01-02 15:04:05"), user.Username+"_login"); err != nil {
			response.FailWithMessage("设置登录时间失败", c)
			global.GVA_LOG.Error("redis set login date failed, err:", zap.Any("err", err))
			return
		}
		response.OkWithData(resp.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: clams.StandardClaims.ExpiresAt * 1000,
		}, c)
	} else if err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		var blackJWT model.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := service.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			global.GVA_LOG.Error("redis JsonInBlacklist failed, err:", zap.Any("err", err))
			return
		}
		if err := service.SetRedisJWT(token, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			global.GVA_LOG.Error("redis SetRedisJWT failed, err:", zap.Any("err", err))
			return
		}

		if err := service.SetRedisJWT(time.Now().Format("2006-01-02 15:04:05"), user.Username+"_login"); err != nil {
			response.FailWithMessage("设置登录时间失败", c)
			global.GVA_LOG.Error("redis set login date failed, err:", zap.Any("err", err))
			return
		}

		response.OkWithData(resp.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: clams.StandardClaims.ExpiresAt * 1000,
		}, c)
	}
}

// @Tags Base
// @Summary 用户注册
// @Produce  application/json
// @Param data body request.RegisterStruct true "用户名, 密码"
// @Success 200 {object} response.Response{data=resp.SysUserResponse,msg=string} "用户注册账号,返回包括用户信息"
// @Router /api/base/register [post]
func Register(c *gin.Context) {
	var R request.RegisterStruct
	_ = c.ShouldBindJSON(&R)
	UserVerify := utils.Rules{
		"Username":  {utils.NotEmpty()},
		"Password":  {utils.NotEmpty()},
		"Captcha":   {utils.NotEmpty()},
		"CaptchaId": {utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(R, UserVerify)
	if UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	if store.Verify(R.CaptchaId, R.Captcha, true) {
		user := &model.SysUser{Username: R.Username, NickName: R.NickName, Password: R.Password, HeaderImg: R.HeaderImg, AuthorityId: R.AuthorityId}
		err, userReturn := service.Register(*user)
		if err != nil {
			response.FailWithDetailed(response.ERROR, resp.SysUserResponse{User: userReturn}, fmt.Sprintf("%v", err), c)
		} else {
			response.OkDetailed(resp.SysUserResponse{User: userReturn}, "注册成功", c)
		}
	} else {
		response.FailWithMessage("验证码错误", c)
	}
}

// SetUserInfo
// @Task User
// @Summary 更新用户用户信息
// @Produce application/json
// @Param data body request.ChangePasswordStruct true "用户名, 旧密码, 新密码, 头像, 个人简介"
// @Success 200 {object} response.Response{data=string,msg=string} "用户注册账号,返回包括用户信息"
// @Router /api/user/setUserInfo [post]
func SetUserInfo(c *gin.Context) {
	var Req request.ChangePasswordStruct
	_ = c.ShouldBindJSON(&Req)
	UserVerify := utils.Rules{
		"Username":  {},
		"Password":  {},
		"NewPassword": {},
		"HeaderImg": {},
		"UserProfile": {},
	}
	UserVerifyErr := utils.Verify(Req, UserVerify)
	if UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	u := &model.SysUser{Username: Req.Username, Password: Req.Password, HeaderImg: Req.HeaderImg, UserProfile: Req.UserProfile}
	err, _ :=service.SetUserInfo(u, Req.NewPassword)
	fmt.Printf("%v\n", u)
	if err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败，原密码与当前账户不符", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}
