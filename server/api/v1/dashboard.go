/**
 * @Author: Yuxue.Fan<yuxue.fan@yidatec.com>
 * @Description:
 * @File: dashboard
 * @Version: 1.0.0
 * @Date: 2022/2/23 14:28
 */

package v1

import (
	"encoding/json"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

func GetInitData(c *gin.Context) {
	var req request.DashBoardReq
	_ = c.ShouldBindJSON(&req)
	dashboardReq := utils.Rules{}
	err := utils.Verify(req, dashboardReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		global.GVA_LOG.Error("dashboard request, err:", zap.Any("err", err))
		return
	}
	claims := c.MustGet("claims").(*request.CustomClaims)
	err, loginDate := service.GetRedisJWT(claims.Username + "_login")
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		global.GVA_LOG.Error("dashboard get login date, err:", zap.Any("err", err))
		return
	}

	server, err := service.GetServerInfo()
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败，%v", err), c)
		return
	}
	data := getIpAddress()

	response.OkWithData(resp.DashboardRsp{
		LoginDate: loginDate,
		Ip: c.ClientIP(),
		Server: *server,
		Other: data,
	}, c)
}

func getIpAddress() resp.OtherRsp {
	res, err := http.Get("https://ipinfo.io/json")
	if err != nil {
		global.GVA_LOG.Error("dashboard getIpAddress failed, err:", zap.Any("err", err))
		return resp.OtherRsp{}
	}

	defer  res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		global.GVA_LOG.Error("dashboard read body failed, err:", zap.Any("err", err))
		return resp.OtherRsp{}
	}
	var data resp.OtherRsp
	json.Unmarshal(body, &data)
	return data
}

