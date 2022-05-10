/**
 * @Author: Yuxue.Fan<yuxue.fan@yidatec.com>
 * @Description:
 * @File: user
 * @Version: 1.0.0
 * @Date: 2022/2/18 14:43
 */

package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/utils"
	uuid2 "github.com/google/uuid"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func Login(u *model.SysUser) (err error, userInter *model.SysUser) {
	var user model.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GVA_READDB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	return err, &user
}

func FindUserByUuid(uuid string) (err error, user *model.SysUser) {
	var u model.SysUser
	if err = global.GVA_READDB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil{
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}

func Register(u model.SysUser) (err error, userInter model.SysUser) {
	var user model.SysUser
	if !errors.Is(global.GVA_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}
	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = utils.MD5V([]byte(u.Password))
	u.UUID = uuid2.UUID(uuid.NewV4())
	err = global.GVA_DB.Create(&u).Error
	return err, u
}

func SetUserInfo(u *model.SysUser, newPassword string) (err error, us *model.SysUser) {
	var user model.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GVA_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Updates(map[string]interface{}{
		"password": utils.MD5V([]byte(newPassword)),
		"header_img": u.HeaderImg,
		"user_profile": u.UserProfile,
	}).Error
	return err, u
	return err, u
}