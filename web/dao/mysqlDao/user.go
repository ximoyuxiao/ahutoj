package mysqldao

import (
	"ahutoj/web/dao"
	"context"
)

func SelectUserByUid(ctx context.Context, user *dao.User) error {
	db := GetDB(ctx)
	// select * from User where uid = ''
	err := db.Where("uid=?", user.Uid).Find(user).Error
	return err
}

func SelectUserCountByUid(ctx context.Context, uid string) (count int64, err error) {
	db := GetDB(ctx)
	err = db.Table(dao.User{}.TableName()).Where("uid=?", uid).Count(&count).Error
	return count, err
}

func InsertUserTable(ctx context.Context, user dao.User) error {
	db := GetDB(ctx)
	err := db.Create(&user).Error
	return err
}

func UpdateUserByUid(ctx context.Context, user *dao.User) error {
	db := GetDB(ctx)
	err := db.Where("uid=?", user.Uid).Updates(&user).Error
	return err
}
