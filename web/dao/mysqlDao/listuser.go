package mysqldao

import (
	"ahutoj/web/dao"
	"context"
)

func InsertListUser(ctx context.Context, listUser dao.ListUser) error {
	db := GetDB(ctx)
	return db.Create(listUser).Error
}
