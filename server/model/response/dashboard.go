/**
 * @Author: Yuxue.Fan<yuxue.fan@yidatec.com>
 * @Description:
 * @File: dashboard
 * @Version: 1.0.0
 * @Date: 2022/2/23 16:05
 */

package response

import "gin-vue-admin/utils"

type DashboardRsp struct {
	LoginDate string `json:"login_date"`
	Ip string `json:"ip"`
	Address string `json:"address"`
	Server utils.Server `json:"server"`
	Other OtherRsp `json:"other"`
}

type OtherRsp struct {
	IP string `json:"ip"`
	City string `json:"city"`
	Region string `json:"region"`
	Country string `json:"country"`
	Loc string `json:"loc"`
	Org string `json:"org"`
	Timezone string `json:"timezone"`
	Readme string `json:"readme"`
}
