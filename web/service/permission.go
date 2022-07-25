package service

import (
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/response"
	"ahutoj/web/logic"
	"ahutoj/web/utils"

	"github.com/gin-gonic/gin"
)

func EditPermission(ctx *gin.Context) {

}

func DeletePermission(ctx *gin.Context) {

}

func AddPermission(ctx *gin.Context) {

}

func GetListPermission(ctx *gin.Context) {

}

func GetPermission(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	uid := ctx.Param("id")

	if uid == "" {
		logger.Errorf("")
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp, _ := logic.GetPermission(ctx, uid)
	response.ResponseOK(ctx, resp)
}
