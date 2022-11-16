package originJudged

import (
	"ahutoj/web/dao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/models"
	"ahutoj/web/utils"
	"context"
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
	for {
		/*1、从数据库 当中 提取外部判题*/
		submits, _ := models.GetOriginJudgeSubmit(context.Background())
		/*2、得到后批量更新状态*/
		for _, submit := range submits {
			logger.Errorf("submit size:%d", len(submits))
			submit.Result = constanct.OJ_JUDGE
			models.UpdateSubmit(context.Background(), submit)
			originJudge := GetOriginJudgeFunc(OJPlatform(submit.OJPlatform))
			if originJudge == nil {
				logger.Errorf("not existe plateform,OJPlatform:%d", submit.OJPlatform)
				continue
			}
			// 执行一个协程。
			go originJudge.Judge(context.Background(), submit, submit.OriginPID)
		}
		time.Sleep(5 * time.Second)
	}
}
