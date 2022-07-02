package service

import (
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/logic"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func LoginSerivce(ctx *gin.Context) {
	req := new(request.LoginReq)
	if err := ctx.ShouldBindWith(req, binding.Form); err != nil {
		// 请求参数有误，直接返回响应
		response.ResponseError(ctx, 201)
		return
	}
	resp, err := logic.CheckLogin(req, ctx)
	if err != nil {
		response.ResponseError(ctx, 202)
	}
	response.ResponseOK(ctx, resp)
}
