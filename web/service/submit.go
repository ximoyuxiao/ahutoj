package service

import (
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/utils"
	"strconv"

	"ahutoj/web/logic"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func AddCommit(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.AddSubmitReq)
	if err := ctx.ShouldBindWith(req, binding.JSON); err != nil {
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}

	resp, err := logic.AddSubmit(ctx, req)
	if err != nil {
		logger.Errorf("call AddSubmit failed, req=%+v, err=%s", utils.Sdump(req), err)
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}

func RejudgeCommit(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.RejudgeSubmitReq)
	if err := ctx.ShouldBindWith(req, binding.JSON); err != nil {
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}

	resp, err := logic.RejudgeSubmit(ctx, req)
	if err != nil {
		logger.Errorf("call RejudgeSubmit failed, req=%+v, err=%s", utils.Sdump(req), err)
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}

func StatusList(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.SubmitListReq)
	if err := ctx.ShouldBindWith(req, binding.Query); err != nil {
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}

	resp, err := logic.GetSubmits(ctx, req)
	if err != nil {
		logger.Errorf("call GetSubmitList failed, req=%+v, err=%s", utils.Sdump(req), err)
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}

func GetCommit(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	var err error
	req := new(request.GetSubmitReq)
	cidStr := ctx.Param("id")
	req.SID, err = strconv.ParseInt(cidStr, 10, 64)
	if err != nil {
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}

	resp, err := logic.GetSubmit(ctx, req)
	if err != nil {
		logger.Errorf("call GetSubmit failed, req=%+v, err=%s", utils.Sdump(req), err)
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}
