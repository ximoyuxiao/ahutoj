package main

import (
	rediscache "ahutoj/web/cache/redis"
	mysqldao "ahutoj/web/dao/mysqlDao"
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
	err = mysqldao.InitMysql(nil)
	logger := utils.GetLogInstance()
	if err != nil {
		logger.Errorf("init mysql error mysqlConf:%+v, err=%s", utils.Sdump(utils.GetConfInstance().MySQLConfig), err.Error())
		os.Exit(1)
	}

	//初始化Redis数据库
	err = rediscache.InitRedisPool()
	if err != nil {
		logger.Errorf("init redis error redisConf=%+v, err=%s", utils.Sdump(utils.GetConfInstance().RedisConfig), err.Error())
		os.Exit(1)
	}
	//初始化JWT
	middlewares.InitJwt()

	middlewares.InitSnowflake(utils.GetConfInstance().StartTime, utils.GetConfInstance().MachineID)

	routers.InitServer()
	return nil
}
