package mysqldao

import (
	"ahutoj/web/dao"
	"context"
)

func InsertListUser(ctx context.Context, listUser dao.ListUser) error {
	db := GetDB(ctx)
	return db.Create(listUser).Error
}

func SelectListUser(ctx context.Context, listUser dao.ListUser) (dao.ListUser, error) {
	db := GetDB(ctx)
	ret := dao.ListUser{}
	err := db.Where(listUser).Find(&ret).Error
	return ret, err
}
