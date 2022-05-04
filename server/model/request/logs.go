/*
@Time : 2022/4/4 17:08
@Author : Yuxue.fan<fanyixuanf@gmail.com>
@File : logs.go
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package request

type LogsReq struct {
	StartDate string `json:"start_date"`
	EndDate string `json:"end_date"`
	PageInfo
}