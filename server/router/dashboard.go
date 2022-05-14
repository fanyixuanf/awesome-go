/**
 * @Author: Yuxue.Fan<yuxue.fan@yidatec.com>
 * @Description:
 * @File: dashboard
 * @Version: 1.0.0
 * @Date: 2022/2/23 14:23
 */

package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDashboardRouter(router *gin.RouterGroup) (R gin.IRoutes) {
	dashboardRouter := router.Group("dashboard").
		Use(middleware.JWTAuth())
		//Use(middleware.OperationRecord())
	{
		dashboardRouter.POST("init", v1.GetInitData)
	}
	return dashboardRouter
}
