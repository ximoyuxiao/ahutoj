package mysqldao

import (
	"ahutoj/web/dao"
	"context"
)

func SelectPermissionByUid(ctx context.Context, uid string) (dao.Permission, error) {
	db := GetDB(ctx)
	permission := dao.Permission{}
	err := db.Table("permission").Where("uid=?", uid).Find(&permission).Error
	return permission, err
}
