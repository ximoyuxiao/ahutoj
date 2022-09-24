package models

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/utils"
	"context"
)

//通过UID判断用户是否存在
func IsUserExistByUID(ctx context.Context, user *dao.User) bool {
	count, err := mysqldao.SelectUserCountByUID(ctx, user.UID)
	if err != nil {
		return false
	}
	return count > 0
}

//通过UID获得用户信息
func FindUserByUID(ctx context.Context, user *dao.User) error {
	return mysqldao.SelectUserByUID(ctx, user)
}

//判断用户密码是否相同
func EqualPassWord(ctx context.Context, user *dao.User, password string) bool {
	md5Password, err := utils.MD5EnCode(user.UID, password)
	if err != nil {
		return false
	}
	return md5Password == user.Pass
}

func CreateUser(ctx context.Context, user *dao.User) error {
	logger := utils.GetLogInstance()
	// 2、密码加密处理（MD5)
	user.Pass, _ = utils.MD5EnCode(user.UID, user.Pass)
	err := mysqldao.InsertUserTable(ctx, *user)
	if err != nil {
		logger.Error("call InsertUserTable failed,user= %+v, err=%s", utils.Sdump(user), err.Error())
	}
	return err
}
