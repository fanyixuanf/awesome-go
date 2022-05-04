/**
 * @Author: Yuxue.Fan<yuxue.fan@yidatec.com>
 * @Description:
 * @File: main.go
 * @Version: 1.0.0
 * @Date: 2022/2/17 9:27
 */

package main

import (
	"gin-vue-admin/core"
	"gin-vue-admin/global"
	"gin-vue-admin/initialize"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8088
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

func main () {
	// 初始化

	// db
	initialize.Gorm()
	// redis
	initialize.Redis()
	// 关闭链接
	db, _ := global.GVA_DB.DB()
	defer db.Close()
	defer global.GVA_REDIS.Close()
	// 定时任务
	initialize.Timer()
	// Run
	core.Run()
}
