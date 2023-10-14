package main

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/middlewares"
	"ahutoj/web/models"
	"ahutoj/web/utils"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func main() {
	ConfigPath := "./config.yaml"
	if len(os.Args) >= 2 {
		ConfigPath = os.Args[1]
	}
	initPersistence(ConfigPath)

	go DealCeInfo()
	go DealSubmit()
	for {
		time.Sleep(time.Hour)
	}
}

func GetConsumer(queueName string) *middlewares.Consumer {
	rmq := middlewares.GetRabbitMq()
	return middlewares.NewConsumer(rmq, queueName)
}

func DealCeInfo() {
	logger := utils.GetLogInstance()
	comm := GetConsumer(constanct.JUDGECE)
	for {
		msgs, err := comm.ConsumeMessage()
		if err != nil {
			logger.Errorf("call ConsumeMessage failed,err:%v", err.Error())
			os.Exit(1)
		}
		for msg := range msgs {
			ceinfo := dao.CeInfo{}
			json.Unmarshal(msg.Body, &ceinfo)
			CEinfoToDataBase(context.Background(), &ceinfo)
		}
	}
}

func DealSubmit() {
	logger := utils.GetLogInstance()
	comm := GetConsumer(constanct.JUDGERESULT)
	for {
		msgs, err := comm.ConsumeMessage()
		if err != nil {
			logger.Errorf("call ConsumeMessage failed,err:%v", err.Error())
			os.Exit(1)
		}
		for msg := range msgs {
			submit := dao.Submit{}
			body := string(msg.Body)
			err = json.Unmarshal([]byte(body), &submit)
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(utils.Sdump(submit))
			SubmitToDataBase(context.Background(), &submit)
		}
	}
}

func CEinfoToDataBase(ctx context.Context, ceinfo *dao.CeInfo) error {
	logger := utils.GetLogInstance()
	err := mysqldao.InsertCeInfo(ctx, *ceinfo)
	logger.Infof(utils.Sdump(ceinfo))
	if err != nil {
		logger.Errorf("call InsertCEinfo failed,ceinfo=%v, err=%v", utils.Sdump(ceinfo), err.Error())
	}
	return err
}

func SubmitToDataBase(ctx context.Context, submit *dao.Submit) error {
	logger := utils.GetLogInstance()
	err := models.UpdateSubmit(context.Background(), *submit)
	if err != nil {
		logger.Errorf("call UpdateSubmit failed,submit=%v, err=%v", utils.Sdump(submit), err.Error())
		return err
	}
	if submit.Result == constanct.OJ_AC {
		mysqldao.IncUserSolved(context.Background(), submit.UID)
		if submit.CID > 0 {
			mysqldao.IncConProSolved(context.Background(), submit.CID, submit.PID)
		}
	}
	if err != nil {
		logger.Errorf("call InsertCEinfo failed,submit=%v, err=%v", utils.Sdump(submit), err.Error())
	}
	return err
}

func initPersistence(ConfigPath string) error {
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
	rbtcfg := utils.GetConfInstance().RabbitMQ
	middlewares.NewRabbitMQ(rbtcfg.Host, rbtcfg.Port, rbtcfg.Username, rbtcfg.Password, 1)
	return nil
}
