package models

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"context"
)

func GetSubmitByCIDFromDB(ctx context.Context, CID, fb int64) ([]dao.Submit, error) {
	return mysqldao.SelectSubmitByCID(ctx, CID, fb)
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
