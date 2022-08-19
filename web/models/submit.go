package models

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"context"
)

func GetSubmitByCidFromDB(ctx context.Context, cid int64) ([]dao.Submit, error) {
	return mysqldao.SelectSubmitByCid(ctx, cid)
}
