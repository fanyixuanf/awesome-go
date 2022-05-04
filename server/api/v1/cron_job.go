/*
@Time : 2022/4/9 22:05
@Author : Yuxue.fan<fanyixuanf@gmail.com>
@File : cron_job.go
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func AddCronJob(c *gin.Context) {
	var Req request.AddCronJob
	_ = c.ShouldBindJSON(&Req)
	ReqVerify := utils.Rules{
		"Name": {utils.NotEmpty()},
		"JobKey": {utils.NotEmpty()},
		"Url": {utils.NotEmpty()},
		"CronExpression": {utils.NotEmpty()},
		"Msg": {utils.NotEmpty()},
		"Desc": {utils.NotEmpty()},
	}
	ReqVerifyErr := utils.Verify(Req, ReqVerify)
	if ReqVerifyErr != nil {
		response.FailWithMessage(ReqVerifyErr.Error(), c)
		return
	}
	cronjob := &model.CronJob{
		Name:           Req.Name,
		JobKey:         Req.JobKey,
		Url:            Req.Url,
		CronExpression: Req.CronExpression,
		Msg:            Req.Msg,
		Desc:           Req.Desc,
		Status:         0,
	}
	err, job := service.AddCronJob(*cronjob)
	if err != nil {
		global.GVA_LOG.Error("添加失败:", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkDetailed(job, "添加成功", c)
}

func StartCronJob(c *gin.Context) {
	var Req request.ChangeStatusCronJob
	_ = c.ShouldBindJSON(&Req)
	ReqVerify := utils.Rules{
		"id" :{utils.NotEmpty()},
	}
	ReqVerifyErr := utils.Verify(Req, ReqVerify)
	if ReqVerifyErr != nil {
		global.GVA_LOG.Error("参数校验错误:", zap.Error(ReqVerifyErr))
		response.FailWithMessage(ReqVerifyErr.Error(), c)
	}
	var id request.ChangeStatusCronJob
	id.Id = Req.Id
	err, data := service.ChangeCronJobStatus(id, 1, 0)
	if err != nil {
		global.GVA_LOG.Error("sql error: ", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 启动cron job
	runCronJob(data, c)
}

func StopCronJob(c *gin.Context) {
	var Req request.ChangeStatusCronJob
	_ = c.ShouldBindJSON(&Req)
	ReqVerify := utils.Rules{
		"id" :{utils.NotEmpty()},
	}
	ReqVerifyErr := utils.Verify(Req, ReqVerify)
	if ReqVerifyErr != nil {
		global.GVA_LOG.Error("参数校验错误:", zap.Error(ReqVerifyErr))
		response.FailWithMessage(ReqVerifyErr.Error(), c)
	}
	var id request.ChangeStatusCronJob
	id.Id = Req.Id
	err, data := service.ChangeCronJobStatus(id, 0, 0)
	if err != nil {
		global.GVA_LOG.Error("sql error: ", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	global.GVA_Timer.StopTask(data.JobKey)
	response.OkDetailed(data, "暂停成功", c)
}

func CronJobList(c *gin.Context) {
	var Req request.CronJobListReq
	_ = c.ShouldBindJSON(&Req)
	ReqVerify := utils.Rules{
		"Page": {utils.NotEmpty()},
		"PageSize": {utils.NotEmpty()},
	}
	ReqVerifyErr := utils.Verify(Req, ReqVerify)
	if ReqVerifyErr != nil {
		global.GVA_LOG.Error("参数校验错误", zap.Error(ReqVerifyErr))
		response.FailWithMessage(ReqVerifyErr.Error(), c)
		return
	}
	var list  request.CronJobListReq
	list.Page = Req.Page
	list.PageSize = Req.PageSize
	err, data, total := service.GetCronJobList(list)
	if err != nil {
		global.GVA_LOG.Error("获取数据失败, err:", zap.Any("err", err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkDetailed(map[string]interface{}{"list":data, "total":total }, "查询成功", c)
}

func runCronJob(data model.CronJob, c *gin.Context) {
	if data.TimerId != 0 {
		global.GVA_Timer.StartTask(data.JobKey)
		response.OkDetailed(data.TimerId, "启动成功", c)
		return
	}
	id, err := global.GVA_Timer.AddTaskByFunc(data.JobKey, data.CronExpression, func() {
		res, err := http.Get(data.Url)
		if err != nil || res.StatusCode != http.StatusOK {
			global.GVA_LOG.Error("发送消息失败", zap.Error(err))
		}
		global.GVA_LOG.Info("Timer start:", zap.Any("start", map[string]interface{}{"start": 1}))
	})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var cid  request.ChangeStatusCronJob
	cid.Id = int(data.ID)
	service.ChangeCronJobStatus(cid, 1, int(id))
	response.OkDetailed(id, "启动成功", c)
}
