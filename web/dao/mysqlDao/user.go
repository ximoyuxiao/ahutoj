package mysqldao

import (
	"ahutoj/web/dao"
	"context"
)

func SelectUserByUid(ctx context.Context, user *dao.User) error {
	db := GetDB(ctx)
	//select * from User where uid = ''
	err := db.Table("User").Where("uid=?", user.Uid).Find(user).Error
	return err
}

func SelectUserCountByUid(ctx context.Context, uid string) (count int64, err error) {
	db := GetDB(ctx)
	err = db.Table("User").Where("uid=?", uid).Count(&count).Error
	return count, err
}
