package main

import (
	originJudged "ahutoj/originJudge/originjudged"
	mysqldao "ahutoj/web/dao/mysqlDao"
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
	err = mysqldao.InitMysql()
	logger := utils.GetLogInstance()
	if err != nil {
		logger.Errorf("init mysql error mysqlConf:%+v, err=%s", utils.Sdump(utils.GetConfInstance().MySQLConfig), err.Error())
		os.Exit(1)
	}

	// //初始化Redis数据库
	// err = redisdao.InitRedisPool()
	// if err != nil {
	// 	logger.Errorf("init redis error redisConf=%+v, err=%s", utils.Sdump(utils.GetConfInstance().RedisConfig), err.Error())
	// 	os.Exit(1)
	// }
	if utils.GetConfInstance().UseOriginJudge {
		originJudged.InitOriginThread()
	}
	return nil
}
