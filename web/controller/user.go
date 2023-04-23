package controller

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
		response.ResponseError(ctx, constanct.ServerErrorCode)
	}

	response.ResponseOK(ctx, resp)
}

func UserStatusInfo(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.UserStatusInfoReq)
	if err := ctx.ShouldBindBodyWith(req, binding.JSON); err != nil {
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}

	resp, err := logic.GetUserStatusInfo(ctx, *req)
	if err != nil {
		logger.Errorf("call GetUserStatusInfo failed, req=%+v, err=%s", utils.Sdump(req), err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
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
		response.ResponseOK(ctx, constanct.InvalidParamCode)
		return
	}
	usr := req.ToUser(middlewares.GetUid(ctx))

	// call db
	if !models.IsUserExistByUID(ctx, usr) {
		response.ResponseOK(ctx, constanct.USER_EDITINFO_UIDNotExistCode)
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
		response.ResponseOK(ctx, constanct.USER_CFBIND_FAILED)
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
		UID: middlewares.GetUid(ctx),
	}
	err = mysqldao.SelectUserByUID(ctx, usr)
	if err != nil {
		logger.Errorf("query user failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.USER_EDITPASS_FAILED)
		return
	}

	if !models.EqualPassWord(ctx, usr, req.OldPwd) {
		logger.Errorf("user old_pwd error!!")
		response.ResponseError(ctx, constanct.USER_EDITPASS_PasswordCode)
		return
	}
	if req.Pwd == "" {
		response.ResponseError(ctx, constanct.USER_EDITPASS_PasswordEmptyCode)
		return
	}
	usr.Pass, _ = utils.MD5EnCode(usr.UID, req.Pwd)

	err = mysqldao.UpdateUserByUID(ctx, usr)
	if err != nil {
		logger.Errorf("update user passwd failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.USER_EDITPASS_FAILED)
		return
	}
	response.ResponseOK(ctx, response.CreateResponse(constanct.SuccessCode))
}

func VjudgeBind(ctx *gin.Context) {
	// logger := utils.GetLogInstance()
	// req := &request.UserEditVjudgeReq{}
	// err := ctx.ShouldBindWith(req, binding.JSON)
	// if err != nil {
	// 	logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
	// 	response.ResponseError(ctx, constanct.InvalidParamCode)
	// 	return
	// }

	// usr := req.ToUser(middlewares.GetUid(ctx))
	// // usr.Vjpwd, _ = utils.MD5EnCode(req.Vjid, req.Vjpwd)

	// // call db
	// err = mysqldao.UpdateUserByUID(ctx, usr)
	// if err != nil {
	// 	logger.Errorf("update mysql error =%s", err.Error())
	// 	response.ResponseError(ctx, constanct.MySQLErrorCode)
	// 	return
	// }

	response.ResponseOK(ctx, response.CreateResponse(constanct.NotimplementedCode))
}

func AddUsersRange(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.AddUsersRangeReq)
	err := ctx.ShouldBindWith(req, binding.JSON)
	if err != nil {
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp, err := logic.AddUsersRange(ctx, *req)
	if err != nil {
		logger.Errorf("call AddUsers err=%s, req=%+v", err.Error(), utils.Sdump(req))
		response.ResponseError(ctx, constanct.ServerErrorCode)
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
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp, err := logic.AddUsers(ctx, *req)
	if err != nil {
		logger.Errorf("call AddUsers err=%s, req=%+v", err.Error(), utils.Sdump(req))
		response.ResponseError(ctx, constanct.ServerErrorCode)
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
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp, err := logic.CodeForceBind(ctx, *req)
	if err != nil {
		logger.Errorf("call CodeForceBind err=%s, req=%+v", err.Error(), utils.Sdump(req))
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}

func EditImage(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	file, err := ctx.FormFile("file")
	if err != nil {
		logger.Errorf("call FormFile filed, err=%s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	logger.Infof("upfile:%s", file.Filename)
	if !checkImageFile(file.Filename) {
		logger.Errorf("chekfile failed filename:%s", file.Filename)
		response.ResponseError(ctx, constanct.USER_EDITIMAG_TPYECODE)
		return
	}

	headPath := utils.GetConfInstance().HeadPath
	//SaveUploadedFile上传表单文件到指定的路径
	err = CheckAndCreatDir(ctx, headPath)
	if err != nil {
		logger.Errorf("call CheckAndCreatDir failed headPath:%s", headPath)
		response.ResponseError(ctx, constanct.USER_EDITIMAG_FAILED)
		return
	}
	suffix := getFileSuffix(file.Filename)
	name := middlewares.GetUid(ctx)
	headURL := headPath + "UID_" + name + "." + suffix
	//更新用户信息
	user := dao.User{
		UID:     middlewares.GetUid(ctx),
		HeadURL: headURL,
	}
	err = mysqldao.UpdateUserByUID(ctx, &user)
	if err != nil {
		logger.Errorf("update Image Failed,err:%v", err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	err = ctx.SaveUploadedFile(file, headURL)
	if err != nil {
		logger.Errorf("update Image Failed,please,headURL=%v,err:%v", headURL, err.Error())
		response.ResponseError(ctx, constanct.USER_EDITIMAG_SAVECODE)
	}
	response.ResponseOK(ctx, struct {
		response.Response
		ImageURL string `json:"ImageURL"`
	}{
		Response: response.CreateResponse(constanct.SuccessCode),
		ImageURL: "image/head/" + "UID_" + name + "." + suffix,
	},
	)
}

func AdminChangePassWord(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.PasswordResetReq)
	err := ctx.ShouldBindWith(req, binding.JSON)
	if err != nil {
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp, err := logic.ResetPassword(ctx, req)
	if err != nil {
		logger.Errorf("call ResetPassword err=%s, req=%+v", err.Error(), utils.Sdump(req))
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}
