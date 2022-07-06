package service

import (
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/logic"
	"ahutoj/web/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func LoginSerivce(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.LoginReq)
	if err := ctx.ShouldBindWith(req, binding.Form); err != nil {
		// 请求参数有误，直接返回响应
		logger.Errorf("call ShouldBindWith failed, err =%s", err.Error())
		response.ResponseError(ctx, 201)
		return
	}
	resp, err := logic.CheckLogin(req, ctx)
	if err != nil {
		logger.Errorf("call CheckLogin failed,req=%+v,err=%s", utils.Sdump(req), err.Error())
		response.ResponseError(ctx, 202)
	}
	logger.Debugf("loginResp=%+v", resp)
	response.ResponseOK(ctx, resp)
}
