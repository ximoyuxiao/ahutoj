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
	req := ctx.Query("UID")
	if req == "" {
		req = middlewares.GetUid(ctx)
	}
	logger.Infof("req:%+v", req)
	resp, err := logic.GetUserInfo(ctx, &req)
	if err != nil {
		logger.Errorf("call GetUserInfo failed,req=%+v,err=%s", req, err.Error())
		response.ResponseError(ctx, constanct.ServerBusyCode)
		return
	}

	response.ResponseOK(ctx, resp)
}
func UserStatusInfo(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.UserStatusInfoReq)
	if err := ctx.ShouldBindBodyWith(req, binding.JSON); err != nil {
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.ParametersInvlidCode)
		return
	}

	resp, err := logic.GetUserStatusInfo(ctx, *req)
	if err != nil {
		logger.Errorf("call GetUserStatusInfo failed, req=%+v, err=%s", utils.Sdump(req), err.Error())
		response.ResponseError(ctx, constanct.ServerBusyCode)
		return
	}
	response.ResponseOK(ctx, resp)
}
func EditUserInfo(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := &request.UserEditReq{}
	err := ctx.ShouldBindWith(req, binding.JSON)
	if err != nil {
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.ParametersInvlidCode)
		return
	}
	usr := req.ToUser(middlewares.GetUid(ctx))

	// call db
	if !models.IsUserExistByUID(ctx, usr) {
		response.ResponseError(ctx, constanct.GetResCode(constanct.User, constanct.Service, constanct.UIDNotExist))
		return
	}

	// erase illegal data
	{
		usr.Pass = ""
		usr.Vjid = ""
		usr.Vjpwd = ""
	}

	err = mysqldao.UpdateUserByUID(ctx, usr)
	if err != nil {
		logger.Errorf("update user info failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.ParametersInvlidCode)
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
		response.ResponseError(ctx, constanct.ParametersInvlidCode)
		return
	}

	// call db
	usr := &dao.User{
		UID: middlewares.GetUid(ctx),
	}
	err = mysqldao.SelectUserByUID(ctx, usr)
	if err != nil {
		logger.Errorf("query user failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.ParametersInvlidCode)
		return
	}

	if !models.EqualPassWord(ctx, usr, req.OldPwd) {
		logger.Errorf("user old_pwd error!!")
		response.ResponseError(ctx, constanct.GetResCode(constanct.User, constanct.Service, constanct.PasswordError))
		return
	}

	usr.Pass, _ = utils.MD5EnCode(usr.UID, req.Pwd)

	err = mysqldao.UpdateUserByUID(ctx, usr)
	if err != nil {
		logger.Errorf("update user passwd failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.ParametersInvlidCode)
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
		response.ResponseError(ctx, constanct.ParametersInvlidCode)
		return
	}

	usr := req.ToUser(middlewares.GetUid(ctx))
	// usr.Vjpwd, _ = utils.MD5EnCode(req.Vjid, req.Vjpwd)

	// call db
	err = mysqldao.UpdateUserByUID(ctx, usr)
	if err != nil {
		logger.Errorf("update mysql error =%s", err.Error())
		response.ResponseError(ctx, constanct.ParametersInvlidCode)
		return
	}

	response.ResponseOK(ctx, response.CreateResponse(constanct.SuccessCode))
}

func AddUsersRange(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.AddUsersRangeReq)
	err := ctx.ShouldBindWith(req, binding.JSON)
	if err != nil {
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.ParametersInvlidCode)
		return
	}
	resp, err := logic.AddUsersRange(ctx, *req)
	if err != nil {
		logger.Errorf("call AddUsers err=%s, req=%+v", err.Error(), utils.Sdump(req))
		response.ResponseError(ctx, constanct.ServerBusyCode)
		return
	}
	response.ResponseOK(ctx, resp)
}

func AddUsers(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.AddUsersReq)
	err := ctx.ShouldBindWith(req, binding.JSON)
	if err != nil {
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.ParametersInvlidCode)
		return
	}
	resp, err := logic.AddUsers(ctx, *req)
	if err != nil {
		logger.Errorf("call AddUsers err=%s, req=%+v", err.Error(), utils.Sdump(req))
		response.ResponseError(ctx, constanct.ServerBusyCode)
		return
	}
	response.ResponseOK(ctx, resp)
}
func CodeForceBind(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.CodeForceBindReq)
	err := ctx.ShouldBindWith(req, binding.JSON)
	if err != nil {
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.ParametersInvlidCode)
		return
	}
	resp, err := logic.CodeForceBind(ctx, *req)
	if err != nil {
		logger.Errorf("call CodeForceBind err=%s, req=%+v", err.Error(), utils.Sdump(req))
		response.ResponseError(ctx, constanct.ServerBusyCode)
		return
	}
	response.ResponseOK(ctx, resp)
}
