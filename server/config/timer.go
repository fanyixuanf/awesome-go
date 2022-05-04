/*
@Time : 2022/3/27 17:25
@Author : Yuxue.fan<fanyixuanf@gmail.com>
@File : timer
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package config

type Timer struct {
	Start  bool     `mapstructure:"start" json:"start" yaml:"start"` // 是否启用
	Spec   string   `mapstructure:"spec" json:"spec" yaml:"spec"`    // CRON表达式
	Detail []Detail `mapstructure:"detail" json:"detail" yaml:"detail"`
}

type Detail struct {
	TableName    string `mapstructure:"tableName" json:"tableName" yaml:"tableName"`          // 需要清理的表名
	CompareField string `mapstructure:"compareField" json:"compareField" yaml:"compareField"` // 需要比较时间的字段
	Interval     string `mapstructure:"interval" json:"interval" yaml:"interval"`             // 时间间隔
}
