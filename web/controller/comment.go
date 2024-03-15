package controller

import (
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/logic"
	"ahutoj/web/middlewares"
	"ahutoj/web/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func CommentOperator(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.CommentReq)
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
			logger.Errorf("add comment failed, because text is null")
			response.ResponseError(ctx, constanct.ServerErrorCode)
			return
		}
		err := logic.AddComment(ctx, req)
		if err != nil {
			logger.Errorf("add comment failed")
			response.ResponseError(ctx, constanct.COMMENT_ADD_FAILED)
			return
		}
		//响应
		response.ResponseOK(ctx, response.CreateResponse(constanct.SuccessCode))

	} else if req.ActionType == constanct.DELETECODE {
		// 检查id不为空
		if req.CID == 0 {
			logger.Errorf("user '%v' delete solution failed, because solutionIDStr is null.", req)
			response.ResponseError(ctx, constanct.ServerErrorCode)
			return
		}
		// 执行删除题解操作
		err = logic.DeleteComment(ctx, req)
		if err != nil {
			logger.Errorf("user '%v' delete comment failed.beceuse %v", req, err)
			response.ResponseError(ctx, constanct.COMMENT_DELETE_FAILED)
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

func GetComments(ctx *gin.Context) {
	// GetCommentList
	logger := utils.GetLogInstance()
	req := new(request.CommentListReq)
	if err := ctx.ShouldBindWith(req, binding.Query); err != nil {
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp := logic.GetCommentList(ctx, req)
	response.ResponseOK(ctx, resp)
}
