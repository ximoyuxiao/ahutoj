package main

import (
	rediscache "ahutoj/web/cache/redis"
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/middlewares"
	"ahutoj/web/models"
	"ahutoj/web/routers"
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
	rbtcfg := utils.GetConfInstance().RabbitMQ
	middlewares.NewRabbitMQ(rbtcfg.Host, rbtcfg.Port, rbtcfg.Username, rbtcfg.Password, 12)
	osscfg := utils.GetConfInstance().OssConfig
	middlewares.NewOss(osscfg.Host, osscfg.Port, osscfg.AccessKeyID, osscfg.SecretAccessKey, osscfg.UseSSL)
	//初始化持久化服务
	go persistence()
	routers.InitServer()
	return nil
}
func persistence() {
	logger := utils.GetLogInstance()

	Ce := middlewares.GetConsumer(constanct.JUDGECE)
	Result := middlewares.GetConsumer(constanct.JUDGERESULT)
	for {
		ces, err := Ce.ConsumeMessage() //发生其他错误时，可以重新尝试获取
		if err != nil {
			logger.Errorf("call ConsumeMessage failed,queue:%v,err:%v", constanct.JUDGECE, err.Error())
			time.Sleep(20 * time.Second)
			continue
		}
		results, err := Result.ConsumeMessage()
		if err != nil {
			logger.Errorf("call ConsumeMessage failed,queue:%v,err:%v", constanct.JUDGERESULT, err.Error())
			time.Sleep(20 * time.Second) //防止日志被顶上去了~
			continue
		}
		//必须消费成功,否则直接退出协程(应等待排查错误)
		for {
			select {
			case msg := <-ces:
				ceinfo := dao.CeInfo{}
				err := json.Unmarshal(msg.Body, &ceinfo)
				if err != nil {
					logger.Errorf("call json.Unmarshal failed,err:%v", err.Error())
					return
				}
				if err := CEinfoToDataBase(context.Background(), &ceinfo); err != nil {
					logger.Errorf("call CEinfoToDataBase failed,err:%v", err.Error())
					return
				}
			case msg := <-results:
				submit := dao.Submit{}
				err = json.Unmarshal(msg.Body, &submit)
				if err != nil {
					logger.Errorf("call json.Unmarshal failed,err:%v", err.Error())
					return
				}
				logger.Debugf("submit:%v", utils.Sdump(submit))
				if err := SubmitToDataBase(context.Background(), &submit); err != nil {
					logger.Errorf("call SubmitToDataBase failed,err:%v", err.Error())
					return
				}
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
		err := mysqldao.IncUserSolved(ctx, submit.UID)
		if err != nil {
			logger.Errorf("call IncUserSolved failed,submit=%v, err=%v", utils.Sdump(submit), err.Error())
			return err
		}
		err = mysqldao.IncProblemSolved(ctx, submit.PID)
		if err != nil {
			logger.Errorf("call IncProblemSolved failed,submit=%v, err=%v", utils.Sdump(submit), err.Error())
			return err
		}
		if submit.CID > 0 {
			err = mysqldao.IncConProSolved(ctx, submit.CID, submit.PID)
			if err != nil {
				logger.Errorf("call IncConProSolved failed,submit=%v, err=%v", utils.Sdump(submit), err.Error())
				return err
			}
		}
	}
	return err
}
