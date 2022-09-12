package models

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/mapping"

	"github.com/gin-gonic/gin"
)

type PermissionType int32

const (
	Administrator   PermissionType = 0
	Contest_creator PermissionType = 1
	Problem_edit    PermissionType = 2
	Source_browser  PermissionType = 3
)

func GetPermission(ctx *gin.Context, uid string) (dao.Permission, error) {
	permission, err := mysqldao.SelectPermissionByUid(ctx, uid)
	if err != nil {
		return dao.Permission{}, err
	}
	return permission, err
}
func AddPermission(ctx *gin.Context, permission *dao.Permission) error {
	return mysqldao.InsertPermission(ctx, permission)
}

func EditPermission(ctx *gin.Context, permission *dao.Permission) error {
	return mysqldao.SavePermission(ctx, permission)
}

func DeletePermission(ctx *gin.Context, uid string) error {
	return mysqldao.DeletePermission(ctx, &uid)
}

func GetPermissionList(ctx *gin.Context, offset, size int) ([]dao.Permission, error) {
	return mysqldao.SelectPermissionList(ctx, offset, size)

}
func PermisionReqToDao(req request.PermissionReq) dao.Permission {
	Permission := dao.Permission{}
	Permission.Uid = req.Uid
	if req.Administrator {
		Permission.Administrator = "Y"
	} else {
		Permission.Administrator = "N"
	}

	if req.Contest_creator {
		Permission.Contest_creator = "Y"
	} else {
		Permission.Contest_creator = "N"
	}

	if req.Problem_edit {
		Permission.Problem_edit = "Y"
	} else {
		Permission.Problem_edit = "N"
	}

	if req.Source_browser {
		Permission.Source_browser = "Y"
	} else {
		Permission.Source_browser = "N"
	}
	return Permission
}

func PermissionDaoToResp(permission dao.Permission) response.Permission {
	return response.Permission{
		Uid:           permission.Uid,
		PermissionMap: mapping.PermissionToBitMap(permission),
	}
}

func CheckUserPermission(ctx *gin.Context, uid string, checkPermission PermissionType) bool {
	permission, err := mysqldao.SelectPermissionByUid(ctx, uid)
	if err != nil {
		return false
	}
	if permission.Administrator == "Y" {
		return true
	}
	switch checkPermission {
	case Administrator:
		return permission.Administrator == "Y"
	case Contest_creator:
		return permission.Contest_creator == "Y"
	case Problem_edit:
		return permission.Problem_edit == "Y"
	case Source_browser:
		return permission.Source_browser == "Y"
	default:
		return false
	}
}
