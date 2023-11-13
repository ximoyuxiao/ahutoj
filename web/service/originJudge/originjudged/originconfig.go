package originJudged

import (
	"ahutoj/web/dao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/middlewares"
	"ahutoj/web/models"
	"ahutoj/web/utils"
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type OJPlatform int64
type OriginFunc interface {
	Judge(ctx context.Context, submit dao.Submit, PID string) error
}
type OriginJudgeUser struct {
	Status   UserStatus
	Cookies  map[string]string
	ID       string
	Password string
}
type OriginJudge struct {
	PID    string     // 平台的题目ID
	Submit dao.Submit // 用户提交代码
	CEInfo string     // CE信息
}

func (p *OriginJudge) CommitResult() error {
	rmq := middlewares.GetRabbitMq()
	pro := middlewares.NewProducer(rmq)
	if p.Submit.Result == constanct.OJ_JUDGE {
		p.Submit.Result = constanct.OJ_FAILED
	}
	err := pro.SendMessage(constanct.JUDGERESULT, p.Submit)
	if err != nil {
		return err
	}
	if p.CEInfo != "" && p.Submit.Result == constanct.OJ_CE {
		return pro.SendMessage(constanct.JUDGECE, dao.CeInfo{
			SID:  p.Submit.SID,
			Info: p.CEInfo,
		})
	}
	return nil
}

var ojMap = map[OJPlatform]OriginFunc{
	Cfoj:      CodeForceJudge{},
	ATcoderoj: AtCoderJudge{},
}

func GetOriginJudgeFunc(oj OJPlatform) OriginFunc {
	originJudgeObj, ok := ojMap[oj]
	if !ok {
		return nil
	}
	return originJudgeObj
}

func InitOriginThread() {
	rand.Seed(time.Now().Unix())
	logger := utils.GetLogInstance()
	rabbit := middlewares.GetRabbitMq()
	if rabbit == nil {
		return
	}
	comm := middlewares.NewConsumer(rabbit, constanct.ORIGINJUDGE)
	for {
		/*1、从数据库 当中 提取外部判题*/

		msgs, err := comm.ConsumeMessage()
		if err != nil {
			logger.Errorf("call ConsumeMessage failed, err=%v", err.Error())
			return
		}
		/*2、得到后批量更新状态*/
		logger.Info("submit size:%d", len(msgs))
		for msg := range msgs {
			submit := dao.Submit{}
			json.Unmarshal(msg.Body, &submit)
			submit.Result = constanct.OJ_JUDGE
			models.UpdateSubmit(context.Background(), submit)
			originJudge := GetOriginJudgeFunc(OJPlatform(submit.OJPlatform))
			if originJudge == nil {
				logger.Errorf("not existe plateform,OJPlatform:%d", submit.OJPlatform)
				continue
			}
			fmt.Println(submit)
			// 执行一个协程。
			go originJudge.Judge(context.Background(), submit, submit.OriginPID)
		}
		time.Sleep(5 * time.Second)
	}
}
