package main

import (
	mysqldao "ahutoj/web/dao/mysqlDao"
	origininfo "ahutoj/web/service/useranalytics/originInfo"
	"ahutoj/web/utils"
	"fmt"
	"os"
	"time"
)

var configPath = "./config.yaml"

func main() {
	if len(os.Args) >= 2 {
		configPath = os.Args[1]
	}
	err := initUserAnalytics(configPath)
	if err != nil {
		panic("load faild")
	}
}

func initUserAnalytics(configPath string) error {
	err := utils.ConfigInit(configPath)
	if err != nil {
		return fmt.Errorf("call InitAppConfig failed")
	}

	err = mysqldao.InitMysql(nil)
	logger := utils.GetLogInstance()
	if err != nil {
		logger.Errorf("init mysql error mysqlConf:%+v, err=%s", utils.Sdump(utils.GetConfInstance().MySQLConfig), err.Error())
		os.Exit(1)
	}
	go GetInnerInfo()
	go GetOriginInfo()
	InitServer()
	return nil
}

func GetInnerInfo() {
	time.Sleep(2 * time.Hour) // 每两个小时更新一次用户信息
}

func GetOriginInfo() {
	origininfo.GetCodeForecesInfo()
	time.Sleep(2 * time.Hour)
}

func InitServer() {
}
