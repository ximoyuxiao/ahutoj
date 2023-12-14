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

func GetPermission(ctx *gin.Context, UID string) (dao.Permission, error) {
	permission, err := mysqldao.SelectPermissionByUID(ctx, UID)
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

func DeletePermission(ctx *gin.Context, UID string) error {
	return mysqldao.DeletePermission(ctx, &UID)
}

func GetPermissionList(ctx *gin.Context, offset, size int) ([]dao.Permission, error) {
	return mysqldao.SelectPermissionList(ctx, offset, size)

}
func GetPermissionCount(ctx *gin.Context) (int64, error) {
	return mysqldao.SelectPermissionCount(ctx)
}
func PermisionReqToDao(req request.PermissionReq) dao.Permission {
	Permission := dao.Permission{}
	Permission.UID = req.UID
	if req.Administrator {
		Permission.SuperAdmin = "Y"
	} else {
		Permission.SuperAdmin = "N"
	}

	if req.Contest_creator {
		Permission.ContestAdmin = "Y"
	} else {
		Permission.ContestAdmin = "N"
	}

	if req.Problem_edit {
		Permission.ProblemAdmin = "Y"
	} else {
		Permission.ProblemAdmin = "N"
	}

	if req.Source_browser {
		Permission.SourceAdmin = "Y"
	} else {
		Permission.SourceAdmin = "N"
	}
	Permission.ListAdmin = "N"
	return Permission
}

func PermissionDaoToResp(permission dao.Permission) response.Permission {
	return response.Permission{
		UID:           permission.UID,
		PermissionMap: mapping.PermissionToBitMap(permission),
	}
}

func CheckUserPermission(ctx *gin.Context, UID string, checkPermission PermissionType) bool {
	permission, err := mysqldao.SelectPermissionByUID(ctx, UID)
	if err != nil {
		return false
	}
	if permission.SuperAdmin == "Y" {
		return true
	}
	switch checkPermission {
	case Administrator:
		return permission.SuperAdmin == "Y"
	case Contest_creator:
		return permission.ContestAdmin == "Y"
	case Problem_edit:
		return permission.ProblemAdmin == "Y"
	case Source_browser:
		return permission.SourceAdmin == "Y"
	default:
		return false
	}
}
