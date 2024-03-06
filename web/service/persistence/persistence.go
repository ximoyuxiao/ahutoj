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
)

func main() {
	ConfigPath := "./config.yaml"
	if len(os.Args) >= 2 {
		ConfigPath = os.Args[1]
	}
	initPersistence(ConfigPath)

	go DealCeInfo()
	go DealSubmit()

	<-make(chan struct{})
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
			err := json.Unmarshal(msg.Body, &ceinfo)
			if err != nil {
				logger.Errorf("json.Unmarshal failed,err:%v", err.Error())
				continue
			}
			if err := CEinfoToDataBase(context.Background(), &ceinfo); err != nil {
				logger.Errorf("CEinfoToDataBase failed,err:%v", err.Error())
			}
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
			err = json.Unmarshal(msg.Body, &submit)
			if err != nil {
				logger.Errorf("json.Unmarshal failed,err:%v", err.Error())
			}
			logger.Debugf("submit:%v", utils.Sdump(submit))
			if err := SubmitToDataBase(context.Background(), &submit); err != nil {
				logger.Errorf("SubmitToDataBase failed,err:%v", err.Error())
			}
		}
	}
}

func CEinfoToDataBase(ctx context.Context, ceinfo *dao.CeInfo) error {
	logger := utils.GetLogInstance()
	err := mysqldao.InsertCeInfo(ctx, *ceinfo)
	logger.Infof(utils.Sdump(ceinfo))
	if err != nil {
		logger.Errorf("call InsertCEinfo failed,ceinfo=%v, err=%v", utils.Sdump(ceinfo), err.Error())
		return err
	}
	return nil
}

func SubmitToDataBase(ctx context.Context, submit *dao.Submit) error {
	logger := utils.GetLogInstance()
	err := models.UpdateSubmit(ctx, *submit)
	if err != nil {
		logger.Errorf("call UpdateSubmit failed,submit=%v, err=%v", utils.Sdump(submit), err.Error())
		return err
	}
	if submit.Result == constanct.OJ_AC {
		mysqldao.IncUserSolved(ctx, submit.UID)
		if submit.CID > 0 {
			mysqldao.IncConProSolved(ctx, submit.CID, submit.PID)
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
	middlewares.NewRabbitMQ(rbtcfg.Host, rbtcfg.Port, rbtcfg.Username, rbtcfg.Password, 2)
	return nil
}
