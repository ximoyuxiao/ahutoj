package service

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/logic"
	"ahutoj/web/middlewares"
	"ahutoj/web/models"
	"ahutoj/web/utils"

	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
)

func UserInfo(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := ctx.Query("uid")
	logger.Infof("req:%+v", req)
	if req == "" {
		req = middlewares.GetUid(ctx)
	}
	resp, err := logic.GetUserInfo(ctx, &req)
	if err != nil {
		logger.Errorf("call GetUserInfo failed,req=%+v,err=%s", req, err.Error())
		response.ResponseError(ctx, constanct.ServerBusyCode)
	}

	response.ResponseOK(ctx, resp)
}

func EditUserInfo(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := &request.UserEditReq{}
	err := ctx.ShouldBindWith(req, binding.JSON)
	if err != nil {
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	usr := req.ToUser(middlewares.GetUid(ctx))

	// call db
	if !models.IsUserExistByUid(ctx, usr) {
		response.ResponseError(ctx, constanct.UIDNotExistCode)
		return
	}

	// erase illegal data
	{
		usr.Pass = ""
		usr.Vjid = ""
		usr.Vjpwd = ""
	}

	err = mysqldao.UpdateUserByUid(ctx, usr)
	if err != nil {
		logger.Errorf("update user info failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.MySQLErrorCode)
		return
	}

	response.ResponseOK(ctx, response.CreateResponse(constanct.SuccessCode))
}

func EditUserPass(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := &request.UserEditPassReq{}
	err := ctx.ShouldBindWith(req, binding.JSON)
	if err != nil {
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}

	// call db
	usr := &dao.User{
		Uid: middlewares.GetUid(ctx),
	}
	err = mysqldao.SelectUserByUid(ctx, usr)
	if err != nil {
		logger.Errorf("query user failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.MySQLErrorCode)
		return
	}

	if !models.EqualPassWord(ctx, usr, req.OldPwd) {
		logger.Errorf("user old_pwd error!!")
		response.ResponseError(ctx, constanct.PassWordErrorCode)
		return
	}

	usr.Pass, _ = utils.MD5EnCode(usr.Uid, req.Pwd)

	err = mysqldao.UpdateUserByUid(ctx, usr)
	if err != nil {
		logger.Errorf("update user passwd failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.MySQLErrorCode)
		return
	}
	response.ResponseOK(ctx, response.CreateResponse(constanct.SuccessCode))
}

func VjudgeBind(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := &request.UserEditVjudgeReq{}
	err := ctx.ShouldBindWith(req, binding.JSON)
	if err != nil {
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}

	usr := req.ToUser(middlewares.GetUid(ctx))
	// usr.Vjpwd, _ = utils.MD5EnCode(req.Vjid, req.Vjpwd)

	// call db
	err = mysqldao.UpdateUserByUid(ctx, usr)
	if err != nil {
		logger.Errorf("update mysql error =%s", err.Error())
		response.ResponseError(ctx, constanct.MySQLErrorCode)
		return
	}

	response.ResponseOK(ctx, response.CreateResponse(constanct.SuccessCode))
}
