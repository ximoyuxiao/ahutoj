package models

import (
	rediscache "ahutoj/web/cache/redis"
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/middlewares"
	"ahutoj/web/utils"
	"context"
)

func GetSubmitByCIDFromDB(ctx context.Context, CID, fb int64) ([]dao.Submit, error) {
	return mysqldao.SelectSubmitByCID(ctx, CID, fb)
}
func GetSubmitByLIDFromDB(ctx context.Context, LID, fb int64) ([]dao.Submit, error) {
	return mysqldao.SelectSubmitByCID(ctx, LID, fb)
}
func CreateSubmit(ctx context.Context, submit dao.Submit) error {
	return mysqldao.InsertSubmit(ctx, submit)
}

func GetSubmitList(ctx context.Context, submit dao.Submit, offset, limit int) ([]dao.Submit, error) {
	return mysqldao.SelectSubmitList(ctx, submit, offset, limit)
}

func RejudgeSubmit(ctx context.Context, submit dao.Submit) error {
	return mysqldao.RejudgeSubmit(ctx, submit)
}

func GetSubmitListCount(ctx context.Context, submit dao.Submit) (int64, error) {
	return mysqldao.SelectCountSubmit(ctx, submit)
}
func GetUserStatusInfo(ctx context.Context, submit dao.Submit, lastTime int64) ([]dao.Submit, error) {
	return mysqldao.SelectSubmitByUID(ctx, submit, lastTime)
}

func UpdateSubmit(ctx context.Context, submit dao.Submit) error {
	return mysqldao.UpdateSubmit(ctx, submit)
}

func GetOriginJudgeSubmit(ctx context.Context) ([]dao.Submit, error) {
	return mysqldao.SelectSubmitIsOriginJudge(ctx)
}

func FindLastSIDByUID(ctx context.Context, UID string) (dao.Submit, error) {
	return mysqldao.FindLastSIDByUID(ctx, UID)
}

func EqualLastSource(ctx context.Context, UID string, PID string, Source string) bool {
	logger := utils.GetLogInstance()
	SourceMD5, err := utils.MD5EnCodeStr(Source)
	if err != nil {
		logger.Errorf("call MD5EnCodeStr failed. Source:%s", Source)
		return false
	}
	Lastsource := rediscache.GetLastSource(ctx, UID, PID)
	if Lastsource == SourceMD5 {
		return true
	}
	rediscache.SetLastSource(ctx, UID, PID, SourceMD5)
	return false
}

func CommitRabitMQ(ctx context.Context, submit dao.Submit) error {
	logger := utils.GetLogInstance()
	if submit.SID == 0 {
		return nil
	}
	rabitmq := middlewares.GetRabbitMq()
	produce := middlewares.NewProducer(rabitmq)
	if submit.IsOriginJudge {
		err := produce.SendMessage(constanct.ORIGINJUDGE, submit)
		if err != nil {
			logger.Errorf("call SendMessage(%s) failed, submit=%v, err=%s", constanct.ORIGINJUDGE, submit, err.Error())
			return err
		}
	} else {
		err := produce.SendMessage(constanct.INNERJUDGE, submit)
		if err != nil {
			logger.Errorf("call SendMessage(%s) failed, submit=%v, err=%s", constanct.INNERJUDGE, submit, err.Error())
			return err
		}
	}
	return nil
}
