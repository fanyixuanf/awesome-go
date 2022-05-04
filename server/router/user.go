package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user").
		Use(middleware.JWTAuth()).
		Use(middleware.OperationRecord())
	{
		UserRouter.POST("setUserInfo", v1.SetUserInfo)     // 更新用户信息
	}
}
