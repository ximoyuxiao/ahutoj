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

func AddTraining(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.List)
	err := ctx.ShouldBindWith(req, binding.JSON)
	if err != nil {
		// 请求参数有误 直接返回响应
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	fmt.Printf("req:%+v\n", req)

	req2 := new(request.ListProblem)
	err2 := ctx.ShouldBindWith(req2, binding.JSON)
	if err2 != nil {
		// 请求参数有误 直接返回响应
		logger.Errorf("call ShouldBindWith failed, err2 = %s", err2.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	fmt.Printf("req:%+v\n", req2)

	resp, err := logic.AddTraining(req, req2, ctx)
	if err != nil {
		logger.Errorf("call DoResiger failed,req=%+v,err=%s", *req, err.Error())
		response.ResponseError(ctx, constanct.ServerBusyCode)
	}
	response.ResponseOK(ctx, resp)
}

func EditTraining(ctx *gin.Context) {

}

func DeleteTraining(ctx *gin.Context) {

}

func GetListTraining(ctx *gin.Context) {

}

func GetTraining(ctx *gin.Context) {

}

func GetRankTraining(ctx *gin.Context) {

}
