/**
 * @Author: Yuxue.Fan<yuxue.fan@yidatec.com>
 * @Description:
 * @File: config.go
 * @Version: 1.0.0
 * @Date: 2022/2/21 8:55
 */

package core

import (
	"flag"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/utils"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

var config string

func init() {
	flag.StringVar(&config, "c", "", "choose config file.")
	flag.Parse()
	if config == "" { // 优先级: 命令行 > 环境变量 > 默认值
		if configEnv := os.Getenv(utils.ConfigEnv); configEnv == "" {
			config = utils.ConfigFile
			fmt.Printf("您正在使用config的默认值,config的路径为%v\n", utils.ConfigFile)
		} else {
			config = configEnv
			fmt.Printf("您正在使用GVA_CONFIG环境变量,config的路径为%v\n", config)
		}
	} else {
		fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
	}
	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
		fmt.Println(err)
	}
	global.GVA_VP = v
}