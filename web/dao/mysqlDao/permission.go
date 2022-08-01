package mysqldao

import (
	"ahutoj/web/dao"
	"context"
)

func SelectPermissionByUid(ctx context.Context, uid string) (dao.Permission, error) {
	db := GetDB(ctx)
	permission := dao.Permission{}
	err := db.Table(permission.TableName()).Where("uid=?", uid).Find(&permission).Error
	return permission, err
}
func SelectPermissionList(ctx context.Context, offset, size int64) ([]dao.Permission, error) {
	db := GetDB(ctx)
	ret := make([]dao.Permission, 0)
	err := db.Table(dao.Permission{}.TableName()).Offset(int(offset)).Limit(int(size)).Find(&ret).Error
	return ret, err
}

func InsertPermission(ctx context.Context, permission *dao.Permission) error {
	db := GetDB(ctx)
	err := db.Table(permission.TableName()).Create(permission).Error
	return err
}

func DeletePermission(ctx context.Context, uid *string) error {
	db := GetDB(ctx)
	err := db.Table(dao.Permission{}.TableName()).Delete(uid).Error
	return err
}

func UpdatePermission(ctx context.Context, permission *dao.Permission) error {
	db := GetDB(ctx)
	err := db.Table(dao.Permission{}.TableName()).Where("uid=?", permission.Uid).Updates(permission).Error
	return err
}

func SavePermission(ctx context.Context, permission *dao.Permission) error {
	db := GetDB(ctx)
	err := db.Table(dao.Permission{}.TableName()).Where("uid=?", permission.Uid).Save(permission).Error
	return err
}
