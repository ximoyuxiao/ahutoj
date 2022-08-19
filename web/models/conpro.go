package models

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"context"
)

func AddConProblemToDb(ctx context.Context, conPro dao.ConPro) error {
	return mysqldao.InsertConProblem(ctx, conPro)
}

func GetConProblemFromDB(ctx context.Context, cid int64) ([]dao.ConPro, error) {
	return mysqldao.SelectConProblemByCid(ctx, cid)
}
