package mysqldao

import (
	"ahutoj/web/dao"
	"context"
)

func SelectUserByUID(ctx context.Context, user *dao.User) error {
	db := GetDB(ctx)
	// select * from User where UID = ''
	err := db.Where("UID=?", user.UID).Find(user).Error
	return err
}

func SelectUserCountByUID(ctx context.Context, UID string) (count int64, err error) {
	db := GetDB(ctx)
	err = db.Table(dao.User{}.TableName()).Where("UID=?", UID).Count(&count).Error
	return count, err
}

func InsertUserTable(ctx context.Context, user dao.User) error {
	db := GetDB(ctx)
	err := db.Create(&user).Error
	return err
}

func UpdateUserByUID(ctx context.Context, user *dao.User) error {
	db := GetDB(ctx)
	err := db.Where("UID=?", user.UID).Updates(&user).Error
	return err
}
