package controller

import (
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/logic"
	"ahutoj/web/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func GetObject(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.GetObjectReq)
	var err error
	if err = ctx.ShouldBindWith(req, binding.JSON); err != nil {
		// 请求参数有误，直接返回响应
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp, err := logic.GetObject(ctx, req)
	if err != nil {
		logger.Errorf("call GetObject failed,req=%+v,err=%s", utils.Sdump(req), err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}
func GetObjects(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.GetObjectsReq)
	var err error
	if err = ctx.ShouldBindWith(req, binding.JSON); err != nil {
		// 请求参数有误，直接返回响应
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp, err := logic.GetObjects(ctx, req)
	if err != nil {
		logger.Errorf("call GetObjects failed,req=%+v,err=%s", utils.Sdump(req), err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}
func FGetObject(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.FGetObjectReq)
	var err error
	if err = ctx.ShouldBindWith(req, binding.JSON); err != nil {
		// 请求参数有误，直接返回响应
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp, err := logic.FGetObject(ctx, req)
	if err != nil {
		logger.Errorf("call FGetObject failed,req=%+v,err=%s", utils.Sdump(req), err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}
func FPutObject(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.FPutObjectReq)
	var err error
	if err = ctx.ShouldBindWith(req, binding.JSON); err != nil {
		// 请求参数有误，直接返回响应
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp, err := logic.FPutObject(ctx, req)
	if err != nil {
		logger.Errorf("call FPutObject failed,req=%+v,err=%s", utils.Sdump(req), err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}

//func CreateObject(ctx *gin.Context) {
//	logger := utils.GetLogInstance()
//	req := new(request.CreateObjectReq)
//	var err error
//	if err = ctx.ShouldBindWith(req, binding.JSON); err != nil {
//		// 请求参数有误，直接返回响应
//		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
//		response.ResponseError(ctx, constanct.InvalidParamCode)
//		return
//	}
//	resp, err := logic.CreateObject(ctx, req)
//	if err != nil {
//		logger.Errorf("call CreateObject failed,req=%+v,err=%s", utils.Sdump(req), err.Error())
//		response.ResponseError(ctx, constanct.ServerErrorCode)
//	}
//	response.ResponseOK(ctx, resp)
//}

func ModifyObject(ctx *gin.Context) {
	response.ResponseError(ctx, constanct.NotimplementedCode)
}

func DeleteObject(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.DeleteObjectReq)
	var err error
	if err = ctx.ShouldBindWith(req, binding.JSON); err != nil {
		// 请求参数有误，直接返回响应
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp, err := logic.DeleteObject(ctx, req)
	if err != nil {
		logger.Errorf("call DeleteObject failed,req=%+v,err=%s", utils.Sdump(req), err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
	}
	response.ResponseOK(ctx, resp)
}

func GetObjectInfo(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.GetObjectInfoReq)
	var err error
	if err = ctx.ShouldBindWith(req, binding.JSON); err != nil {
		// 请求参数有误，直接返回响应
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp, err := logic.GetObjectInfo(ctx, req)
	if err != nil {
		logger.Errorf("call GetObjectInfo failed,req=%+v,err=%s", utils.Sdump(req), err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}

func GetBuckets(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	var err error
	resp, err := logic.GetBucket(ctx)
	if err != nil {
		logger.Errorf("call GetBucket failed,err=%s", err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}

func CreateBucket(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.CreateBucketreq)
	var err error

	if err = ctx.ShouldBindWith(req, binding.JSON); err != nil {
		// 请求参数有误，直接返回响应
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}

	resp, err := logic.CreateBucket(ctx, req)
	if err != nil {
		logger.Errorf("call CreateBucket failed,req=%+v,err=%s", utils.Sdump(req), err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}

func RemoveBucket(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.RemoveBucketreq)
	var err error

	if err = ctx.ShouldBindWith(req, binding.JSON); err != nil {
		// 请求参数有误，直接返回响应
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp, err := logic.RemoveBucket(ctx, req)
	if err != nil {
		logger.Errorf("call RemoveBucket failed,req=%+v,err=%s", utils.Sdump(req), err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}

//func UnzipObject(ctx *gin.Context) {
//	logger := utils.GetLogInstance()
//	req := new(request.UnzipReq)
//	var err error
//	if err = ctx.ShouldBindWith(req, binding.JSON); err != nil {
//		// 请求参数有误，直接返回响应
//		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
//		response.ResponseError(ctx, constanct.InvalidParamCode)
//		return
//	}
//	resp, err := logic.UnzipObject(ctx, req)
//	if err != nil {
//		logger.Errorf("call UnzipObject failed,req=%+v,err=%s", utils.Sdump(req), err.Error())
//		response.ResponseError(ctx, constanct.ServerErrorCode)
//		return
//	}
//	response.ResponseOK(ctx, resp)
//}
