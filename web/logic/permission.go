package logic

import (
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/models"
	"ahutoj/web/utils"

	"github.com/gin-gonic/gin"
)

func GetPermission(ctx *gin.Context, UID string) (interface{}, error) {
	logger := utils.GetLogInstance()
	permission, err := models.GetPermission(ctx, UID)
	permission.UID = UID
	if err != nil {
		logger.Errorf("call GetPermission failed , UID=%s, err=%s", UID, err.Error())
		return response.CreateResponse(constanct.ADMIN_GET_FAILED), err
	}
	return response.PermissionResp{
		Response:   response.CreateResponse(constanct.SuccessCode),
		Permission: models.PermissionDaoToResp(permission),
	}, nil
}

func EditPermission(ctx *gin.Context, req *request.EditPermissionReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	permission := models.PermisionReqToDao(req.PermissionReq)
	if req.UID == "admin" {
		return response.CreateResponse(constanct.ADMIN_EDIT_ADMIN), nil
	}
	err := models.EditPermission(ctx, &permission)
	if err != nil {
		logger.Errorf("Call EditPermission Failed,err=%s", err.Error())
		return response.CreateResponse(constanct.ADMIN_EDIT_FAILED), err
	}
	return response.CreateResponse(constanct.SuccessCode), nil
}

func DeletePermission(ctx *gin.Context, req *request.DeletePermissionReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	for _, UID := range req.UIDs {
		err := models.DeletePermission(ctx, UID)
		if err != nil {
			logger.Errorf("Call DeletePermission Failed,err=%s", err.Error())
			return response.CreateResponse(constanct.ADMIN_DELETE_FAILED), err
		}
	}
	return response.CreateResponse(constanct.SuccessCode), nil
}

func AddPermission(ctx *gin.Context, req *request.AddPermissionReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	if req.UID == "" {
		return response.CreateResponse(constanct.ADMIN_ADD_UIDEmpty), nil
	}
	permission := models.PermisionReqToDao(req.PermissionReq)
	err := models.AddPermission(ctx, &permission)
	if err != nil {
		logger.Errorf("call AddPermission Failed,err=%s", err.Error())
		return nil, err
	}
	return response.CreateResponse(constanct.SuccessCode), nil
}

func GetPermissionList(ctx *gin.Context, req *request.PermissionListReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	var ret response.PermissionListResp
	var size int = 20
	if req.Limit > 20 {
		size = req.Limit
	}
	var offset int = 0
	if req.Page > 0 {
		offset = size * req.Page
	}
	permissions, err := models.GetPermissionList(ctx, offset, size)
	if err != nil {
		logger.Errorf("call GetPermissionList Failed,err=%s", err.Error())
		return response.CreateResponse(constanct.ADMIN_LIST_FAILED), err
	}
	ret.Conut = len(permissions)
	ret.Data = make([]response.Permission, 0, len(permissions))
	for _, permission := range permissions {
		ret.Data = append(ret.Data, models.PermissionDaoToResp(permission))
	}
	return ret, nil
}
