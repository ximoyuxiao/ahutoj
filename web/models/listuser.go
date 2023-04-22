package models

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"context"
)

func SaveTraningUser(ctx context.Context, listUser dao.ListUser) error {
	return mysqldao.InsertListUser(ctx, listUser)
}

func FindTrainUserInfo(ctx context.Context, listUser dao.ListUser) ([]dao.ListUser, error) {
	return mysqldao.SelectListUser(ctx, listUser)
}
