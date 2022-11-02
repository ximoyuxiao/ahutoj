package main

import (
	mysqldao "ahutoj/web/dao/mysqlDao"
	redisdao "ahutoj/web/dao/redisDao"
	"ahutoj/web/middlewares"
	originjudge "ahutoj/web/originJudge"
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
	fmt.Println("error server down!")
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
	logger := utils.GetLogInstance()
	if err != nil {
		logger.Errorf("init mysql error mysqlConf:%+v, err=%s", utils.Sdump(utils.GetConfInstance().MySQLConfig), err.Error())
		os.Exit(1)
	}

	//初始化Redis数据库
	err = redisdao.InitRedisPool()
	if err != nil {
		logger.Errorf("init redis error redisConf=%+v, err=%s", utils.Sdump(utils.GetConfInstance().RedisConfig), err.Error())
		os.Exit(1)
	}
	//初始化JWT
	middlewares.InitJwt()

	middlewares.InitSnowflake(utils.GetConfInstance().StartTime, utils.GetConfInstance().MachineID)
	// 初始化 重判题目的协程
	if utils.GetConfInstance().UseOriginJudge {
		go originjudge.InitOriginThread()
	}
	routers.InitServer()
	return nil
}
