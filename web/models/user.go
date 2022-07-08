package models

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/utils"
	"context"
)

//通过uid判断用户是否存在
func IsUserExistByUid(ctx context.Context, user *dao.User) bool {
	count, err := mysqldao.SelectUserCountByUid(ctx, user.Uid)
	if err != nil {
		return false
	}
	return count > 0
}

//通过uid获得用户信息
func FindUserByUid(ctx context.Context, user *dao.User) error {
	return mysqldao.SelectUserByUid(ctx, user)
}

//判断用户密码是否相同
func EqualPassWord(ctx context.Context, user *dao.User, password string) bool {
	md5Password, err := utils.MD5EnCode(user.Uid, password)
	if err != nil {
		return false
	}
	return md5Password == user.Pass
}

func CreateUser(ctx context.Context, user *dao.User) error {
	logger := utils.GetLogInstance()
	err := mysqldao.InserUserTable(ctx, *user)
	if err != nil {
		logger.Error("call InserUserTable failed,user= %+v, err=%s", utils.Sdump(user), err.Error())
	}
	return err
}
