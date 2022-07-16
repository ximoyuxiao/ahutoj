package service

import (
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/response"
	"ahutoj/web/logic"
	"ahutoj/web/utils"

	"github.com/gin-gonic/gin"
)

func UserInfo(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := ctx.Query("uid")
	logger.Infof("req:%+v", req)

	resp, err := logic.GetUserInfo(ctx, &req)
	if err != nil {
		logger.Errorf("call GetUserInfo failed,req=%+v,err=%s", req, err.Error())
		response.ResponseError(ctx, constanct.ServerBusyCode)
	}

	response.ResponseOK(ctx, resp)
}

func EditUserInfo(ctx *gin.Context) {

}

func EditUserPass(ctx *gin.Context) {

}

func VjudgeBind(ctx *gin.Context) {

}
