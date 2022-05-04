/*
@Time : 2022/4/9 21:58
@Author : Yuxue.fan<fanyixuanf@gmail.com>
@File : cron_job.go
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package request

type AddCronJob struct {
	Name string `json:"name"`
	JobKey string `json:"job_key"`
	Url string `json:"url"`
	CronExpression string `json:"cron_expression"`
	Msg string `json:"msg"`
	Desc string `json:"desc"`
}

type ChangeStatusCronJob struct {
	Id int `json:"id"`
}

type CronJobListReq struct {
	PageInfo
}