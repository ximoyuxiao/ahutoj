package logic

import (
	"ahutoj/web/dao"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/middlewares"
	"ahutoj/web/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CheckLogin(req *request.LoginReq, c *gin.Context) (interface{}, error) {
	user := dao.User{
		Uid: req.Username,
	}
	if ok := models.IsUserExistByUid(c, &user); !ok {
		return response.Response{
			StatusCode: 205,
			StatusMsg:  "账号不存在",
		}, nil
	}
	if err := models.FindUserByUid(c, &user); err != nil {
		return response.Response{
			StatusCode: 205,
			StatusMsg:  "账号不存在",
		}, err
	}
	fmt.Printf("%+v", user)
	ok := models.EqualPassWord(c, &user, req.Password)
	if !ok {
		return response.Response{
			StatusCode: 204,
			StatusMsg:  "密码错误",
		}, nil
	}
	token, err := middlewares.GetToken(user.Uid)
	if err != nil {
		return response.Response{
			StatusCode: 205,
			StatusMsg:  "Token创建失败",
		}, nil
	}
	return response.UserResp{
		Response: response.Response{
			StatusCode: 200,
			StatusMsg:  "登录成功",
		},
		Token: token,
	}, nil
}
