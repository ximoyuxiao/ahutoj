package service

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
		response.ResponseError(ctx, constanct.GetResCode(constanct.Contest, constanct.Service, constanct.Parsesparameters))
		return
	}
	resp, err := logic.AddContest(ctx, req)
	if err != nil {
		logger.Errorf("call AddContest failed, err = %s", err.Error())
		response.ResponseServerError(ctx, constanct.GetResCode(constanct.Contest, constanct.Service, constanct.ServerBusy))
		return
	}
	response.ResponseOK(ctx, resp)
}

func EditContest(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.EditContestReq)
	if err := ctx.ShouldBindBodyWith(req, binding.JSON); err != nil {
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.GetResCode(constanct.Contest, constanct.Service, constanct.Parsesparameters))
		return
	}
	resp, err := logic.EditContest(ctx, req)
	if err != nil {
		logger.Errorf("call EditContest failed, err = %s", err.Error())
		response.ResponseServerError(ctx, constanct.GetResCode(constanct.Contest, constanct.Service, constanct.ServerBusy))
		return
	}
	response.ResponseOK(ctx, resp)
}

func DeleteContest(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.DeleteContestReq)
	if err := ctx.ShouldBindBodyWith(req, binding.JSON); err != nil {
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.GetResCode(constanct.Contest, constanct.Service, constanct.Parsesparameters))
		return
	}
	resp, err := logic.DeleteContest(ctx, req)
	if err != nil {
		logger.Errorf("call DeleteContest failed, err = %s", err.Error())
		response.ResponseServerError(ctx, constanct.GetResCode(constanct.Contest, constanct.Service, constanct.ServerBusy))
		return
	}
	response.ResponseOK(ctx, resp)
}

func GetListContest(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.ContestListReq)
	if err := ctx.ShouldBindWith(req, binding.Query); err != nil {
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.GetResCode(constanct.Contest, constanct.Service, constanct.Parsesparameters))
		return
	}
	resp, err := logic.GetListContest(ctx, req)
	if err != nil {
		logger.Errorf("call GetListContest failed, err = %s", err.Error())
		response.ResponseServerError(ctx, constanct.GetResCode(constanct.Contest, constanct.Service, constanct.ServerBusy))
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
		response.ResponseError(ctx, constanct.GetResCode(constanct.Contest, constanct.Service, constanct.Parsesparameters))
	}
	CIDStr := ctx.Param("id")
	if CIDStr == "" {
		logger.Errorf("call Param failed, err=%s", err.Error())
		response.ResponseError(ctx, constanct.GetResCode(constanct.Contest, constanct.Service, constanct.Parsesparameters))
		return
	}

	req.CID, err = strconv.ParseInt(CIDStr, 10, 64)
	if err != nil {
		logger.Errorf("call GetContest fialed,err=%s", err.Error())
		response.ResponseError(ctx, constanct.GetResCode(constanct.Contest, constanct.Service, constanct.Parsesparameters))
		return
	}
	logger.Infof("req:%+v", utils.Sdump(req))
	resp, err := logic.GetContest(ctx, req)
	if err != nil {
		logger.Errorf("call GetContest failed, err = %s", err.Error())
		response.ResponseServerError(ctx, constanct.GetResCode(constanct.Contest, constanct.Service, constanct.ServerBusy))
		return
	}
	response.ResponseOK(ctx, resp)
}

func GteRankContest(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.GetContestRankReq)
	CIDStr := ctx.Param("id")
	if CIDStr == "" {
		logger.Errorf("cid is empty")
		response.ResponseError(ctx, constanct.GetResCode(constanct.Contest, constanct.Service, constanct.Parsesparameters))
		return
	}

	var err error
	req.CID, err = strconv.ParseInt(CIDStr, 10, 64)
	if err != nil {
		logger.Errorf("call GetContest fialed,err=%s", err.Error())
		response.ResponseError(ctx, constanct.GetResCode(constanct.Contest, constanct.Service, constanct.Parsesparameters))
		return
	}

	if err = ctx.ShouldBindWith(&req, binding.Query); err != nil {
		logger.Errorf("call Param failed, err")
		response.ResponseError(ctx, constanct.GetResCode(constanct.Contest, constanct.Service, constanct.Parsesparameters))
		return
	}

	resp, err := logic.GteRankContest(ctx, req)
	if err != nil {
		logger.Errorf("call AddContest failed, err = %s", err.Error())
		response.ResponseServerError(ctx, constanct.GetResCode(constanct.Contest, constanct.Service, constanct.ServerBusy))
		return
	}
	response.ResponseOK(ctx, resp)
}
