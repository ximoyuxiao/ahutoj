package models

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	redisdao "ahutoj/web/dao/redisDao"
	"ahutoj/web/utils"
	"context"

	"github.com/bytedance/gopkg/util/logger"
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
	SourceMD5, err := utils.MD5EnCodeStr(Source)
	if err != nil {
		logger.Errorf("call MD5EnCodeStr failed. Source:%s", Source)
		return false
	}
	Lastsource := redisdao.GetLastSource(ctx, UID, PID)
	if Lastsource == SourceMD5 {
		return true
	}
	redisdao.SetLastSource(ctx, UID, PID, SourceMD5)
	return false
}
