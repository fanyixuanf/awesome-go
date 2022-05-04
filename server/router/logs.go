/*
@Time : 2022/4/4 16:42
@Author : Yuxue.fan<fanyixuanf@gmail.com>
@File : logs.go
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

func InitLogsRouter(Router *gin.RouterGroup) {
	LogsRouter := Router.Group("logs").
		Use(middleware.JWTAuth()).
		Use(middleware.OperationRecord())
	{
		LogsRouter.POST("getLogsList", v1.GetLogsList)
	}
}
