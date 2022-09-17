package models

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/utils"
	"context"
)

func GetSubmitByCIDFromDB(ctx context.Context, CID, page, limit int) ([]dao.Submit, error) {
	temp := dao.Submit{
		CID: CID,
	}
	offset, size := utils.GetPageInfo(page, limit)
	return mysqldao.SelectSubmitList(ctx, temp, offset, size)
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
