package service

import (
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/logic"
	"ahutoj/web/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Login(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.LoginReq)
	if err := ctx.ShouldBindWith(req, binding.JSON); err != nil {
		// 请求参数有误，直接返回响应
		logger.Errorf("call ShouldBindWith failed, err =%s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp, err := logic.CheckLogin(req, ctx)
	if err != nil {
		logger.Errorf("call CheckLogin failed,req=%+v,err=%s", utils.Sdump(req), err.Error())
		response.ResponseError(ctx, resp.(response.Response).StatusCode)
	}
	response.ResponseOK(ctx, resp)
}

func Register(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.User)
	//1、 获取参数
	err := ctx.ShouldBindWith(req, binding.JSON)
	if err != nil {
		logger.Errorf("call ShouldBindWith failed, err =%s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	logger.Infof("req:%+v\n", req)

	//2、 处理业务逻辑
	resp, err := logic.DoResiger(ctx, req)
	if err != nil {
		logger.Errorf("call DoResiger failed,req=%+v,err=%s", *req, err.Error())
		response.ResponseError(ctx, constanct.ServerBusyCode)
	}

	//3、 构建响应值，将处理结果返回
	response.ResponseOK(ctx, resp)
}

func Logout(ctx *gin.Context) {
	response.ResponseOK(ctx, response.CreateResponse(constanct.SuccessCode))
}
