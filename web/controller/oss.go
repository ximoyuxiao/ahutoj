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

func GetObject(ctx *gin.Context) {
	response.ResponseError(ctx, constanct.NotimplementedCode)
}
func GetObjects(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.GetObjectsReq)
	var err error
	if err := ctx.ShouldBindWith(req, binding.JSON); err != nil {
		// 请求参数有误，直接返回响应
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp, err := logic.GetObjects(ctx, req)
	if err != nil {
		logger.Errorf("call GetNotice failed,req=%+v,err=%s", utils.Sdump(req), err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}
func CreateObject(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.UpObjectReq)
	if err := ctx.ShouldBindWith(req, binding.JSON); err != nil {
		// 请求参数有误，直接返回响应
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp, err := logic.CreateObject(ctx, req)
	if err != nil {
		logger.Errorf("call GetNotice failed,req=%+v,err=%s", utils.Sdump(req), err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
	}
	response.ResponseOK(ctx, resp)
}

func ModifyObject(ctx *gin.Context) {
	response.ResponseError(ctx, constanct.NotimplementedCode)
}

func DeleteObject(ctx *gin.Context) {
	response.ResponseError(ctx, constanct.NotimplementedCode)
}

func GetObjectInfo(ctx *gin.Context) {
	response.ResponseError(ctx, constanct.NotimplementedCode)
}
