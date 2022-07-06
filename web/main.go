package main

import (
	mysqldao "ahutoj/web/dao/mysqlDao"
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

	routers.InitServer()

}

func initAPP(ConfigPath string) error {
	//初始化配置文件
	err := utils.ConfigInit(ConfigPath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	//初始化日志服务
	utils.LogInit()
	//初始化MySQL数据库
	err = mysqldao.InitMysql()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//初始化JWT略

	return nil
}
