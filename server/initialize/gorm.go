/**
 * @Author: Yuxue.Fan<yuxue.fan@yidatec.com>
 * @Description:
 * @File: gorm.go
 * @Version: 1.0.0
 * @Date: 2022/2/17 10:17
 */

package initialize

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"

	"gorm.io/driver/mysql"
)

func Gorm() {
	GormMysql()
}

// GormMysql 初始化Mysql数据库
func GormMysql() {
	m := global.GVA_CONFIG.Mysql
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	gormConfig := config(m.LogMode)
	//var err error
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig); err != nil {
		global.GVA_LOG.Error("mysql connect ping failed, err:", zap.Any("err", err))
		os.Exit(0)
	} else {
		global.GVA_DB = db
		global.GVA_LOG.Info("mysql connect success")
		//GormDBTables(global.GVA_DB)
		//sqlDB, _ := global.GVA_DB.DB()
		//sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		//sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	}
}

// config 根据配置决定是否开启日志
func config(mod bool) (c *gorm.Config) {
	if mod {
		c = &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	} else {
		c = &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	}
	return
}

// GormDBTables 注册数据库表专用
func GormDBTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.SysUser{},
	)
	if err != nil {
		global.GVA_LOG.Info("register table failed")
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
