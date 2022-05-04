/*
@Time : 2022/4/9 21:25
@Author : Yuxue.fan<fanyixuanf@gmail.com>
@File : cron_job.go
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package model

import "gorm.io/gorm"

type CronJob struct {
	gorm.Model
	Name string `json:"name" gorm:"comment:名称"`
	JobKey string `json:"job_key"`
	Url string `json:"url"`
	CronExpression string `json:"cron_expression"`
	Msg string `json:"msg"`
	Desc string `json:"desc"`
	Status int `json:"status"`
	TimerId int `json:"timer_id"`
}
