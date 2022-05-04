/**
 * @Author: Yuxue.Fan<yuxue.fan@yidatec.com>
 * @Description:
 * @File: server.go
 * @Version: 1.0.0
 * @Date: 2022/2/17 10:39
 */

package core

import (
	"fmt"
	"gin-vue-admin/initialize"
	"time"
)

type server interface {
	ListenAndServe() error
}

func Run() {
	Router := initialize.Routers()
	//address := fmt.Sprintf(":%q", "8088")
	s := initServer(":8088", Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	fmt.Println(s.ListenAndServe().Error())

}