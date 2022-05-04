/*
@Time : 2022/4/9 22:08
@Author : Yuxue.fan<fanyixuanf@gmail.com>
@File : cron_job.go
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitCronJobRouter(Router *gin.RouterGroup) {
	CronJobRouter := Router.Group("cronjob").
		Use(middleware.JWTAuth()).
		Use(middleware.OperationRecord())
	{
		CronJobRouter.POST("addCronJob", v1.AddCronJob)
		CronJobRouter.POST("startCronJob", v1.StartCronJob)
		CronJobRouter.POST("stopCronJob", v1.StopCronJob)
		CronJobRouter.POST("cronJobList", v1.CronJobList)
	}
}