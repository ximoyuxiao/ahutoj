package logic

import (
	"ahutoj/web/dao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/middlewares"
	"ahutoj/web/models"

	"github.com/gin-gonic/gin"
)

func CheckLogin(req *request.LoginReq, c *gin.Context) (interface{}, error) {
	user := dao.User{
		Uid: req.Username,
	}
	if ok := models.IsUserExistByUid(c, &user); !ok {
		return response.Response{
			StatusCode: constanct.UIDNotExistCode,
			StatusMsg:  constanct.UIDNotExistCode.Msg(),
		}, nil
	}
	if err := models.FindUserByUid(c, &user); err != nil {
		return response.Response{
			StatusCode: constanct.MySQLErrorCode,
			StatusMsg:  constanct.MySQLErrorCode.Msg(),
		}, err
	}
	ok := models.EqualPassWord(c, &user, req.Password)
	if !ok {
		return response.Response{
			StatusCode: constanct.PassWordErrorCode,
			StatusMsg:  constanct.PassWordErrorCode.Msg(),
		}, nil
	}
	token, err := middlewares.GetToken(user.Uid)
	if err != nil {
		return response.Response{
			StatusCode: constanct.TokenBuildErrorCode,
			StatusMsg:  constanct.TokenBuildErrorCode.Msg(),
		}, nil
	}
	return response.UserResp{
		Response: response.Response{
			StatusCode: constanct.SuccessCode,
			StatusMsg:  constanct.SuccessCode.Msg(),
		},
		Token: token,
	}, nil
}
