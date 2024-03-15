package mysqldao

import (
	"ahutoj/web/dao"
	"context"
	"fmt"

	"gorm.io/gorm"
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
func UpdateEmailByUID(ctx context.Context, UID string, email string) error {
	db := GetDB(ctx)
	err := db.Table(dao.User{}.TableName()).Where("UID=?", UID).Update("Email", email).Error
	return err
}
func SelectUserList(ctx context.Context) ([]dao.User, error) {
	db := GetDB(ctx)
	users := make([]dao.User, 0)
	err := db.Table(dao.User{}.TableName()).Select(&users).Error
	return users, err
}

func IncUserSubmited(ctx context.Context, UID string) error {
	db := GetDB(ctx)
	user := dao.User{}
	return db.Table(user.TableName()).Where("UID=?", UID).UpdateColumn("Submited", gorm.Expr("Submited+1")).Error
}

func IncUserSolved(ctx context.Context, UID string) error {
	db := GetDB(ctx)
	user := dao.User{}
	fmt.Println(UID)
	return db.Table(user.TableName()).Where("UID=?", UID).UpdateColumn("Solved", gorm.Expr("Solved+1")).Error
}
