package controller

import (
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/logic"
	"ahutoj/web/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func AddPermission(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.AddPermissionReq)
	if err := ctx.ShouldBindWith(req, binding.JSON); err != nil {
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp, err := logic.AddPermission(ctx, req)
	if err != nil {
		logger.Errorf("call AddPermission failed,req=%+v, err=%s", *req, err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}

func EditPermission(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.EditPermissionReq)
	if err := ctx.ShouldBindWith(req, binding.JSON); err != nil {
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp, err := logic.EditPermission(ctx, req)
	if err != nil {
		logger.Errorf("call EditPermission failed,req=%+v, err=%s", *req, err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}

func DeletePermission(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.DeletePermissionReq)
	if err := ctx.ShouldBindWith(req, binding.JSON); err != nil {
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp, err := logic.DeletePermission(ctx, req)
	if err != nil {
		logger.Errorf("call DeletePermission failed,req=%+v, err=%s", *req, err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}

func GetListPermission(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.PermissionListReq)
	if err := ctx.ShouldBindWith(req, binding.Query); err != nil {
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp, err := logic.GetPermissionList(ctx, req)
	if err != nil {
		logger.Errorf("call GetPermissionList failed,req=%+v, err=%s", *req, err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}

func GetPermission(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	UID := ctx.Param("id")

	if UID == "" {
		logger.Errorf("")
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp, err := logic.GetPermission(ctx, UID)
	if err != nil {
		logger.Errorf("call GetPermission failed,UID=%+v, err=%s", UID, err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}
