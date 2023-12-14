package mysqldao

import (
	"ahutoj/web/dao"
	"context"
)

func SelectPermissionByUID(ctx context.Context, UID string) (dao.Permission, error) {
	db := GetDB(ctx)
	permission := dao.Permission{}
	err := db.Table(permission.TableName()).Where("UID=?", UID).Find(&permission).Error
	return permission, err
}
func SelectPermissionList(ctx context.Context, offset, size int) ([]dao.Permission, error) {
	db := GetDB(ctx)
	ret := make([]dao.Permission, 0)
	err := db.Table(dao.Permission{}.TableName()).Offset(offset).Limit(size).Find(&ret).Error
	return ret, err
}

func InsertPermission(ctx context.Context, permission *dao.Permission) error {
	db := GetDB(ctx)
	err := db.Table(permission.TableName()).Create(permission).Error
	return err
}

func DeletePermission(ctx context.Context, UID *string) error {
	db := GetDB(ctx)
	err := db.Table(dao.Permission{}.TableName()).Where("UID=?", *UID).Delete(UID).Error
	return err
}

func UpdatePermission(ctx context.Context, permission *dao.Permission) error {
	db := GetDB(ctx)
	err := db.Table(dao.Permission{}.TableName()).Where("UID=?", permission.UID).Updates(permission).Error
	return err
}

func SavePermission(ctx context.Context, permission *dao.Permission) error {
	db := GetDB(ctx)
	err := db.Table(dao.Permission{}.TableName()).Where("UID=?", permission.UID).Save(permission).Error
	return err
}

func SelectPermissionCount(ctx context.Context) (int64, error) {
	db := GetDB(ctx)
	var count int64
	err := db.Table(dao.Permission{}.TableName()).Count(&count).Error
	return count, err
}
