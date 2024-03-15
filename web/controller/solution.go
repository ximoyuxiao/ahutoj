package controller

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/logic"
	"ahutoj/web/middlewares"
	"ahutoj/web/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"strconv"
)

func SolutionOperator(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.SolutionReq)
	err := ctx.ShouldBindWith(req, binding.JSON)
	if err != nil {
		// 请求参数有误 直接返回响应
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	UID := req.Uid
	if UID != middlewares.GetUid(ctx) {
		logger.Errorf("Failed to get user information, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
	}
	if req.ActionType == constanct.ADDCODE {
		//判断内容是否为空
		if req.Text == "" {
			logger.Errorf("add solution failed, because text is null")
			response.ResponseError(ctx, constanct.ServerErrorCode)
			return
		}
		resp, err := logic.AddSoulution(ctx, req)
		if err != nil {
			logger.Errorf("add solution failed")
			response.ResponseError(ctx, constanct.SOLUTION_ADD_FAILED)
			return
		}
		//响应
		response.ResponseOK(ctx, resp)

	} else if req.ActionType == constanct.EDITCODE {
		err := logic.EditSolution(ctx, req)
		if err != nil {
			logger.Errorf("call EditSolution failed, err = %s", err.Error())
			response.ResponseError(ctx, constanct.SOLUTION_EDIT_FAILED)
			return
		}
		response.ResponseOK(ctx, response.CreateResponse(constanct.SuccessCode))
	} else if req.ActionType == constanct.DELETECODE {
		// 检查id不为空
		if req.Sid == 0 {
			logger.Errorf("user '%v' delete solution failed, because solutionIDStr is null.", req)
			response.ResponseError(ctx, constanct.ServerErrorCode)
			return
		}
		// 执行删除题解操作
		err = logic.DeleteSolution(ctx, req)
		if err != nil {
			logger.Errorf("user '%v' delete solution failed.beceuse %v", utils.Sdump(req), err)
			response.ResponseError(ctx, constanct.SOLUTION_DELETE_FAILED)
			return
		}
		//成功
		response.ResponseOK(ctx, response.CreateResponse(constanct.SuccessCode))
	} else {
		logger.Errorf("Unknown request parameters")
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
}

func GetSolution(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.GetSolutionReq)
	SIDstr := ctx.Param("id")
	SID, err := strconv.Atoi(SIDstr)
	req.SID = int64(SID)
	if err != nil {
		logger.Errorf("call Atoi failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	db := mysqldao.GetDB(ctx)
	var solution dao.Solution
	err = db.Where(req.SID).Find(&solution).Error
	UID := middlewares.GetUid(ctx)
	resp := response.SoultionResp{
		SolutionList: response.SolutionResponseElement{
			Data:       logic.GetSubCommentList(ctx, int64(req.SID)),
			Sid:        &solution.SID,
			Text:       &solution.Text,
			Title:      &solution.Title,
			Uid:        &solution.UID,
			IsFavorite: logic.MyFavorite(ctx, int(req.SID), UID),
		},
	}
	if err != nil {
		logger.Errorf("call AddPermission failed,req=%+v, err=%s", *req, err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}

func GetSolutions(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.GetSolutionListReq)
	if err := ctx.ShouldBindWith(req, binding.Query); err != nil {
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp, err := logic.GetSolutiontList(ctx, req)
	if err != nil {
		logger.Errorf("call GetSoulutions failed,req=%+v, err=%s", *req, err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}
