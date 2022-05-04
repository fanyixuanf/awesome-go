/**
 * @Author: Yuxue.Fan<yuxue.fan@yidatec.com>
 * @Description:
 * @File: base.go
 * @Version: 1.0.0
 * @Date: 2022/2/17 16:01
 */

package router

import (
"gin-vue-admin/api/v1"
"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("register", v1.Register)
		BaseRouter.POST("login", v1.Login)
		BaseRouter.POST("captcha", v1.Captcha)
	}
	return BaseRouter
}