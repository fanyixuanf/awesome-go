/*
@Time : 2022/4/4 16:46
@Author : Yuxue.fan<fanyixuanf@gmail.com>
@File : logs.go
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model/request"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetLogsList(c *gin.Context) {
	var Req request.LogsReq
	_ = c.ShouldBindJSON(&Req)
	var Info request.SysOperationRecordSearch
	Info.Page = Req.Page
	Info.PageSize = Req.PageSize
	Info.Body = Req.StartDate
	Info.Resp = Req.EndDate
	err, data, total :=service.GetSysOperationRecordInfoList(Info)
	if err != nil {
		global.GVA_LOG.Error("获取数据失败, err:", zap.Any("err", Req))
	}
	response.OkDetailed(map[string]interface{}{"list":data, "total":total }, "查询成功", c)
}
