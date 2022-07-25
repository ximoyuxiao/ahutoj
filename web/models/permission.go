package models

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"

	"github.com/gin-gonic/gin"
)

func GetPermission(ctx *gin.Context, uid string) (dao.Permission, error) {
	permission, err := mysqldao.SelectPermissionByUid(ctx, uid)
	if err != nil {
		return dao.Permission{}, err
	}
	return permission, err
}
