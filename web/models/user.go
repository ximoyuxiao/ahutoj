package models

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/utils"
	"context"
	"fmt"
)

func IsUserExistByUid(ctx context.Context, user *dao.User) bool {
	count, err := mysqldao.SelectUserCountByUid(ctx, user.Uid)
	if err != nil {
		return false
	}
	return count > 0
}
func FindUserByUid(ctx context.Context, user *dao.User) error {
	return mysqldao.SelectUserByUid(ctx, user)
}

func EqualPassWord(ctx context.Context, user *dao.User, password string) bool {
	md5Password, err := utils.MD5EnCode(user.Uid, password)
	if err != nil {
		return false
	}
	fmt.Println(md5Password)
	fmt.Println(user.Pass)
	return md5Password == user.Pass
}
