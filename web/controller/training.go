package controller

import (
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/logic"
	"ahutoj/web/middlewares"
	"ahutoj/web/utils"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func AddTraining(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.ListAll)
	err := ctx.ShouldBindWith(req, binding.JSON)
	if err != nil {
		// 请求参数有误 直接返回响应
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	fmt.Printf("req:%+v\n", req)

	resp, err := logic.AddTraining(req, ctx)
	if err != nil {
		logger.Errorf("call AddTraining failed,req=%+v,err=%s", *req, err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}

func RegisterTraining(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.RegisterTrainingReq)
	err := ctx.ShouldBindWith(req, binding.JSON)
	if err != nil {
		// 请求参数有误 直接返回响应
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	req.UID = middlewares.GetUid(ctx)
	if req.UID == "" {
		response.ResponseError(ctx, constanct.AUTH_Token_EmptyCode)
	}
	fmt.Printf("req:%+v\n", req)

	resp, err := logic.RegisterTraining(ctx, req)
	if err != nil {
		logger.Errorf("call DoResiger failed,req=%+v,err=%s", *req, err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
	}
	response.ResponseOK(ctx, resp)
}

func EditTraining(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.EditListReq)
	err := ctx.ShouldBindWith(req, binding.JSON)
	if err != nil {
		// 请求参数有误 直接返回响应
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	fmt.Printf("req:%+v\n", req)

	resp, err := logic.EditTraining(req, ctx)
	if err != nil {
		logger.Errorf("call DoResiger failed,req=%+v,err=%s", *req, err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
	}
	response.ResponseOK(ctx, resp)
}

func DeleteTraining(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.DelListReq)
	err := ctx.ShouldBindWith(req, binding.JSON)
	if err != nil {
		// 请求参数有误 直接返回响应
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	fmt.Printf("req:%+v\n", req)

	resp, err := logic.DeleteTraining(req, ctx)
	if err != nil {
		logger.Errorf("call DoResiger failed,req=%+v,err=%s", *req, err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
	}
	response.ResponseOK(ctx, resp)
}

func GetListTraining(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.TrainingListReq)
	if err := ctx.ShouldBindWith(req, binding.Query); err != nil {
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp, err := logic.GetTrainingList(ctx, req)
	if err != nil {
		logger.Errorf("call GetTrainingList failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}

func GetTraining(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.TrainingReq)
	err := ctx.ShouldBindWith(req, binding.Query)
	if err != nil {
		logger.Errorf("call ShouldBindwith failed,err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}

	LIDStr := ctx.Param("id")
	if LIDStr == "" {
		logger.Errorf("call Param failed, err=%s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}

	req.LID, err = strconv.ParseInt(LIDStr, 10, 64)
	if err != nil {
		logger.Errorf("call GetContest fialed,err=%s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	logger.Infof("req:%+v", utils.Sdump(req))
	resp, err := logic.GetTraining(ctx, req)
	if err != nil {
		logger.Errorf("call GetTraining failed,err = %s", err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}

func GetRankTraining(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.GetTrainingRankReq)
	LIDStr := ctx.Param("id")
	if LIDStr == "" {
		logger.Errorf("lid is empty")
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}

	var err error
	req.LID, err = strconv.ParseInt(LIDStr, 10, 64)
	if err != nil {
		logger.Errorf("call GetTraining fialed,err=%s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}

	if err = ctx.ShouldBindWith(&req, binding.Query); err != nil {
		logger.Errorf("call Param failed, err")
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}

	resp, err := logic.GetRankTraining(ctx, req)
	if err != nil {
		logger.Errorf("call GetRankTraining failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	response.ResponseOK(ctx, resp)
}

func GetTrainUserInfo(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.ListUserReq)
	err := ctx.ShouldBindWith(req, binding.Query)
	if err != nil {
		// 请求参数有误 直接返回响应
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	fmt.Printf("req:%+v\n", req)

	resp, err := logic.GetTrainUserInfo(ctx, req)
	if err != nil {
		logger.Errorf("call GetTrainUserInfo failed,req=%+v,err=%s", *req, err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}
