package main

import (
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/utils"
	"fmt"
	"os"
)

var configPath = "./config.yaml"

func main() {
	if len(os.Args) >= 2 {
		configPath = os.Args[1]
	}
	err := initUserAnalytics(configPath)
	panic(err)
}

func initUserAnalytics(configPath string) error {
	err := utils.ConfigInit(configPath)
	if err != nil {
		return fmt.Errorf("call InitAppConfig failed")
	}
	mysqldao.InitMysql(nil)

	go GetInnerInfo()
	go GetOriginInfo()
	InitServer()
	return nil
}

func GetInnerInfo()  {}
func GetOriginInfo() {}
func InitServer()    {}
