/*
@Time : 2022/3/27 17:28
@Author : Yuxue.fan<fanyixuanf@gmail.com>
@File : timer
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package initialize

import (
	timerCfg "gin-vue-admin/config"
	"gin-vue-admin/global"
	"go.uber.org/zap"
)

func Timer() {
	if global.GVA_CONFIG.Timer.Start {
		for i := range global.GVA_CONFIG.Timer.Detail {
			go func(detail timerCfg.Detail) {
				global.GVA_Timer.AddTaskByFunc("ClearDB", global.GVA_CONFIG.Timer.Spec, func() {

					global.GVA_LOG.Info("Timer start:", zap.String("timer start:",  global.GVA_CONFIG.Timer.Detail[i].TableName))
				})
			}(global.GVA_CONFIG.Timer.Detail[i])
		}
	}
}
