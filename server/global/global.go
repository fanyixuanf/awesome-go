/**
 * @Author: Yuxue.Fan<yuxue.fan@yidatec.com>
 * @Description:
 * @File: global
 * @Version: 1.0.0
 * @Date: 2022/2/18 10:24
 */

package global

import (
	"gin-vue-admin/utils/timer"
	"go.uber.org/zap"

	"gin-vue-admin/config"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GVA_DB     *gorm.DB
	GVA_REDIS  *redis.Client
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	GVA_LOG    *zap.Logger
	GVA_Timer  timer.Timer = timer.NewTimerTask()
)
