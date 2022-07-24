package service

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

func AddService(ctx *gin.Context) {
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
		logger.Errorf("call DoResiger failed,req=%+v,err=%s", *req, err.Error())
		response.ResponseError(ctx, constanct.ServerBusyCode)
	}
	response.ResponseOK(ctx, resp)
}

func GetListService(ctx *gin.Context) {

}

func GetService(ctx *gin.Context) {

}

func EditService(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.Problem)
	err := ctx.ShouldBindWith(req, binding.JSON)
	if err != nil {
		//请求参数有误 直接返回响应
		logger.Errorf("call ShouldBindWith failed, err =%s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	logger.Info("req:%+v\n", req)
	resp, err := logic.EditProblem(req, ctx)
	if err != nil {
		logger.Errorf("call DoResiger failed,req=%+v,err=%s", *req, err.Error())
		response.ResponseError(ctx, constanct.ServerBusyCode)
	}
	response.ResponseOK(ctx, resp)
}

func DeleteService(ctx *gin.Context) {

}