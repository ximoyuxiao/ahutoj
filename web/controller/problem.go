package controller

import (
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/logic"
	"ahutoj/web/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func AddProblem(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.Problem)
	err := ctx.ShouldBindWith(req, binding.JSON)
	if err != nil {
		// 请求参数有误 直接返回响应
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	fmt.Printf("req:%+v\n", req)
	resp, err := logic.AddProblem(req, ctx)
	if err != nil {
		logger.Errorf("call AddProblem failed,req=%+v,err=%s", *req, err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}

func GetProblemList(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.ProblemListReq)
	err := ctx.ShouldBindWith(req, binding.Query)
	if err != nil {
		// 请求参数有误 直接返回响应
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	err = ctx.ShouldBindJSON(req)
	// err = ctx.ShouldBindBodyWith(req, binding.JSON)
	if err != nil && err.Error() != "EOF" {
		logger.Errorf("call ShouldBindBodyWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp, err := logic.GetProblemList(ctx, req)
	if err != nil {
		logger.Errorf("call GetProblemList failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}

func GetProblem(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	PID := ctx.Param("id")
	resp, err := logic.GetProblemInfo(ctx, PID)
	if err != nil {
		logger.Errorf("call GetProblemInfo failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}

func EditProblem(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.EditProblemReq)
	err := ctx.ShouldBindWith(req, binding.JSON)
	if err != nil {
		//请求参数有误 直接返回响应
		logger.Errorf("call ShouldBindWith failed, err =%s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp, err := logic.EditProblem(req, ctx)
	if err != nil {
		logger.Errorf("call DoResiger failed,req=%+v,err=%s", *req, err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}

func DeleteProblem(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.DeleteProblemReq)
	err := ctx.BindJSON(req)
	if err != nil {
		logger.Errorf("call ShouldBindWith failed, err =%s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp, err := logic.DeleteProblem(ctx, req)
	if err != nil {
		logger.Errorf("call DoResiger failed,req=%+v,err=%s", *req, err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}
