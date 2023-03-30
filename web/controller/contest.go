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

func AddContest(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.AddContestReq)
	if err := ctx.ShouldBindBodyWith(req, binding.JSON); err != nil {
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp, err := logic.AddContest(ctx, req)
	if err != nil {
		logger.Errorf("call AddContest failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}

func EditContest(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.EditContestReq)
	if err := ctx.ShouldBindBodyWith(req, binding.JSON); err != nil {
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp, err := logic.EditContest(ctx, req)
	if err != nil {
		logger.Errorf("call EditContest failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}

func DeleteContest(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.DeleteContestReq)
	if err := ctx.ShouldBindBodyWith(req, binding.JSON); err != nil {
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp, err := logic.DeleteContest(ctx, req)
	if err != nil {
		logger.Errorf("call DeleteContest failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}

func GetListContest(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.ContestListReq)
	if err := ctx.ShouldBindWith(req, binding.Query); err != nil {
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp, err := logic.GetListContest(ctx, req)
	if err != nil {
		logger.Errorf("call GetListContest failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}

func GetContest(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.GetContestReq)
	var err error
	err = ctx.ShouldBindWith(req, binding.Query)
	if err != nil {
		logger.Errorf("call ShouldBindWith failed, err=%s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
	}
	CIDStr := ctx.Param("id")
	if CIDStr == "" {
		logger.Errorf("call Param failed, err=%s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}

	req.CID, err = strconv.ParseInt(CIDStr, 10, 64)
	if err != nil {
		logger.Errorf("call GetContest fialed,err=%s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	logger.Infof("req:%+v", utils.Sdump(req))
	resp, err := logic.GetContest(ctx, req)
	if err != nil {
		logger.Errorf("call GetContest failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}

func GteRankContest(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.GetContestRankReq)
	CIDStr := ctx.Param("id")
	if CIDStr == "" {
		logger.Errorf("CID is empty")
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}

	var err error
	req.CID, err = strconv.ParseInt(CIDStr, 10, 64)
	if err != nil {
		logger.Errorf("call GetContest fialed,err=%s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}

	if err = ctx.ShouldBindWith(&req, binding.Query); err != nil {
		logger.Errorf("call Param failed, err")
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}

	resp, err := logic.GetRankContest(ctx, req)
	if err != nil {
		logger.Errorf("call AddContest failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}
