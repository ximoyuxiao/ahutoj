package main

import (
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/middlewares"
	originJudged "ahutoj/web/service/originJudge/originjudged"
	"ahutoj/web/utils"
	"fmt"
	"os"
)

func main() {
	ConfigPath := "./config.yaml"
	if len(os.Args) >= 2 {
		ConfigPath = os.Args[1]
	}
	InitOriginJudged(ConfigPath)
	fmt.Println("error server down!")
}

func InitOriginJudged(ConfigPath string) error {
	err := utils.ConfigInit(ConfigPath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	//初始化日志服务
	utils.LogInit() //判题日志

	//初始化MySQL数据库
	err = mysqldao.InitMysql(nil)
	logger := utils.GetLogInstance()
	if err != nil {
		logger.Errorf("init mysql error mysqlConf:%+v, err=%s", utils.Sdump(utils.GetConfInstance().MySQLConfig), err.Error())
		os.Exit(1)
	}
	rbtcfg := utils.GetConfInstance().RabbitMQ
	_, err = middlewares.NewRabbitMQ(rbtcfg.Host, rbtcfg.Port, rbtcfg.Username, rbtcfg.Password, 1)
	if err != nil {
		logger.Errorf("init RabbitMQ error mysqlConf:%+v, err=%s", utils.Sdump(utils.GetConfInstance().RabbitMQ), err.Error())
		os.Exit(1)
	}
	if utils.GetConfInstance().UseOriginJudge {
		originJudged.InitOriginThread()
	}
	return nil
}
