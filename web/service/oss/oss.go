package main

import (
	"ahutoj/web/routers"
	"ahutoj/web/utils"
	"fmt"
	"os"
)

func main() {
	ConfigPath := "./config.yaml"
	if len(os.Args) >= 2 {
		ConfigPath = os.Args[1]
	}
	initAPP(ConfigPath)
}

func initAPP(ConfigPath string) {
	err := utils.ConfigInit(ConfigPath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	//初始化日志服务
	utils.LogInit()
	routers.InitOssServer()
}
