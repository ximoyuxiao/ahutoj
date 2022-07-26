package logic

import (
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/response"
	"ahutoj/web/models"
	"ahutoj/web/utils"

	"github.com/gin-gonic/gin"
)

func GetPermission(ctx *gin.Context, uid string) (interface{}, error) {
	logger := utils.GetLogInstance()
	permission, err := models.GetPermission(ctx, uid)

	if err != nil {
		logger.Errorf("call GetPermission failed , uid=%d, err=%s", uid, err.Error())
		return nil, err
	}
	return response.PermissionResp{
		Response: response.CreateResponse(constanct.SuccessCode),
		Permission: response.Permission{
			Administrator:   permission.Administrator == "Y",
			Problem_edit:    permission.Problem_edit == "Y",
			Source_browser:  permission.Source_browser == "Y",
			Contest_creator: permission.Contest_creator == "Y",
		},
	}, nil
}
