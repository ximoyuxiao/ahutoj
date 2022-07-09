package logic

import (
	"ahutoj/web/dao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/middlewares"
	"ahutoj/web/models"
	"ahutoj/web/utils"

	"github.com/gin-gonic/gin"
)

func CheckLogin(req *request.LoginReq, c *gin.Context) (interface{}, error) {
	user := dao.User{
		Uid: req.Uid,
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
	ok := models.EqualPassWord(c, &user, req.Pass)
	if !ok {
		return response.Response{
			StatusCode: constanct.PassWordErrorCode,
			StatusMsg:  constanct.PassWordErrorCode.Msg(),
		}, nil
	}
	token, err := middlewares.GetToken(c, user.Uid)
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
func DoResiger(c *gin.Context, req *request.User) (interface{}, error) {
	// 1、查看用户账号是否存在
	user := dao.User{
		Uid:     req.Uid,
		Uname:   req.Uname,
		Pass:    req.Pass,
		School:  req.School,
		Classes: req.Classes,
		Major:   req.Major,
		Vjid:    req.Vjid,
		Vjpwd:   req.Vjpwd,
		Email:   req.Email,
	}
	exist := models.IsUserExistByUid(c, &user)
	if exist {
		return response.CreateResponse(constanct.UIDExistCOde), nil
	}
	// 2、密码加密处理（MD5)
	user.Pass, _ = utils.MD5EnCode(req.Uid, req.Pass)
	// 3、创建用户
	err := models.CreateUser(c, &user)
	if err != nil {
		//日志报错
		utils.GetLogInstance().Errorf("call CreateUser failed,err=%s", err.Error())
		return response.CreateResponse(constanct.MySQLErrorCode), err
	}
	// 4、返回注册成功的信息给用户
	return response.CreateResponse(constanct.SuccessCode), nil
}
