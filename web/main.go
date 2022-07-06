package main

import (
	mysqldao "ahutoj/web/dao/mysqlDao"
	redisdao "ahutoj/web/dao/redisDao"
	"ahutoj/web/middlewares"
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
	logger := utils.GetLogInstance()
	if err != nil {
		logger.Errorf("init mysql error mysqlConf:%+v, err=%s", utils.Sdump(utils.GetInstance().MySQLConfig), err.Error())
		os.Exit(1)
	}

	//初始化Redis数据库
	err = redisdao.InitRedis()
	if err != nil {
		logger.Errorf("init redis error redisConf=%+v, err=%s", utils.Sdump(utils.GetInstance().RedisConfig), err.Error())
		os.Exit(1)
	}
	//初始化JWT略
	middlewares.InitJwt()
	return nil
}
