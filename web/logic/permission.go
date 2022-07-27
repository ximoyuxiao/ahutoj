package logic

import (
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/models"
	"ahutoj/web/utils"

	"github.com/gin-gonic/gin"
)

func GetPermission(ctx *gin.Context, uid string) (interface{}, error) {
	logger := utils.GetLogInstance()
	permission, err := models.GetPermission(ctx, uid)

	if err != nil {
		logger.Errorf("call GetPermission failed , uid=%d, err=%s", uid, err.Error())
		return nil, err
	}
	return response.PermissionResp{
		Response:   response.CreateResponse(constanct.SuccessCode),
		Permission: models.PermissionDaoToResp(permission),
	}, nil
}

func EditPermission(ctx *gin.Context, req *request.EditPermissionReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	permission := models.PermisionReqToDao(req.PermissionReq)
	err := models.EditPermission(ctx, &permission)
	if err != nil {
		logger.Errorf("Call EditPermission Failed,err=%s", err.Error())
		return nil, err
	}
	return response.CreateResponse(constanct.SuccessCode), nil
}

func DeletePermission(ctx *gin.Context, req *request.DeletePermissionReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	for _, uid := range req.Uids {
		err := models.DeletePermission(ctx, uid)
		if err != nil {
			logger.Errorf("Call DeletePermission Failed,err=%s", err.Error())
			return nil, err
		}
	}
	return response.CreateResponse(constanct.SuccessCode), nil
}

func AddPermission(ctx *gin.Context, req *request.AddPermissionReq) (interface{}, error) {
	logger := utils.GetLogInstance()
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
	var size int64 = 20
	if req.Limit > 20 {
		size = req.Limit
	}
	var offset int64 = 0
	if req.Page > 0 {
		offset = size * req.Page
	}
	permissions, err := models.GetPermissionList(ctx, offset, size)
	if err != nil {
		logger.Errorf("call GetPermissionList Failed,err=%s", err.Error())
		return nil, err
	}
	ret.Conut = len(permissions)
	ret.Data = make([]response.Permission, 0, len(permissions))
	for _, permission := range permissions {
		ret.Data = append(ret.Data, models.PermissionDaoToResp(permission))
	}
	return ret, nil
}
