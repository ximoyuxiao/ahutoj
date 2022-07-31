package mysqldao

import (
	"ahutoj/web/dao"
	"context"
)

func SelectTraningByLid(ctx context.Context, lid int64) (dao.List, error) {
	db := GetDB(ctx)
}

func SelectTraningList(ctx context.Context, offset, limit int64) (dao.Permission, error) {
	db := GetDB(ctx)
}

func InsertTraning(ctx context.Context, traning dao.List) error {
	db := GetDB(ctx)
}

func DeleteTraning(ctx context.Context, lid int64) error {
	db := GetDB(ctx)
}

func UpdateTraning(ctx context.Context, traning dao.List) error {
	db := GetDB(ctx)
}
func SaveTraning(ctx context.Context, traning dao.List) error {
	db := GetDB(ctx)
}
