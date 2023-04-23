package controller

import (
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/logic"
	"ahutoj/web/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func GetNotice(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.GetNoticeReq)
	var err error
	req.ID, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logger.Errorf("call Atoi failed,req=%+v,err=%s", ctx.Param("id"), err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
	}
	resp, err := logic.GetNotice(ctx, req)
	if err != nil {
		logger.Errorf("call GetNotice failed,req=%+v,err=%s", utils.Sdump(req), err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
	}
	response.ResponseOK(ctx, resp)
}

func GetNoticeList(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	// req := new(request.GetNoticeReq)
	// var err error
	// req.ID, err = strconv.Atoi(ctx.Param("id"))
	// if err != nil {
	// 	logger.Errorf("call Atoi failed,req=%+v,err=%s", ctx.Param("id"), err.Error())
	// 	response.ResponseError(ctx, constanct.InvalidParamCode)
	// }
	resp, err := logic.GetNoticeList(ctx)
	if err != nil {
		logger.Errorf("call GetNoticeList failed,err=%s", err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
	}
	response.ResponseOK(ctx, resp)
}

func CreateNotice(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.CreateNoticeReq)
	if err := ctx.ShouldBindWith(req, binding.JSON); err != nil {
		// 请求参数有误，直接返回响应
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp, err := logic.CreateNotice(ctx, req)
	if err != nil {
		logger.Errorf("call CheckLogin failed,req=%+v,err=%s", utils.Sdump(req), err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
	}
	response.ResponseOK(ctx, resp)
}

func UpdateNotice(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.UpdateNoticeReq)
	var err error
	req.ID, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logger.Errorf("call Atoi failed,param=%v, err = %s", ctx.Param("id"), err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	if err := ctx.ShouldBindWith(req, binding.JSON); err != nil {
		// 请求参数有误，直接返回响应
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp, err := logic.UpdateNotice(ctx, req)
	if err != nil {
		logger.Errorf("call CheckLogin failed,req=%+v,err=%s", utils.Sdump(req), err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
	}
	response.ResponseOK(ctx, resp)
}

func DeleteNotice(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.DeleteNoticeReq)
	var err error
	req.ID, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logger.Errorf("call Atoi failed,param=%v, err = %s", ctx.Param("id"), err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp, err := logic.DeleteNotice(ctx, req)
	if err != nil {
		logger.Errorf("call DeleteNotice failed,req=%+v,err=%s", utils.Sdump(req), err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
	}
	response.ResponseOK(ctx, resp)
}
