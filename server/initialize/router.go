/**
 * @Author: Yuxue.Fan<yuxue.fan@yidatec.com>
 * @Description:
 * @File: router
 * @Version: 1.0.0
 * @Date: 2022/2/17 12:37
 */

package initialize

import (
	_ "gin-vue-admin/docs"
	"gin-vue-admin/middleware"
	"gin-vue-admin/router"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Routers() *gin.Engine {
	var Router = gin.Default()

	// Router.Use(middleware.LoadTls())  // 打开就能玩https了

	// 跨域
	Router.Use(middleware.Cors())
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 方便统一添加路由组前缀 多服务器上线使用
	ApiGroup := Router.Group("")
	// base
	router.InitBaseRouter(ApiGroup)
	// dashboard
	router.InitDashboardRouter(ApiGroup)
	// user
	router.InitUserRouter(ApiGroup)
	// logs
	router.InitLogsRouter(ApiGroup)
	// cron job
	router.InitCronJobRouter(ApiGroup)

	return Router
}