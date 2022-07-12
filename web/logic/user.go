package logic

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/middlewares"
	"ahutoj/web/models"
	"ahutoj/web/utils"

	"github.com/gin-gonic/gin"
)

func CheckLogin(req *request.LoginReq, c *gin.Context) (interface{}, error) {
	logger := utils.GetLogInstance()
	user := dao.User{
		Uid: req.Uid,
	}
	if req.Uid == "" {
		return response.CreateResponse(constanct.UIDEmpty), nil
	}
	if req.Pass == "" {
		return response.CreateResponse(constanct.PassEmpty), nil
	}
	if ok := models.IsUserExistByUid(c, &user); !ok {
		return response.CreateResponse(constanct.UIDNotExistCode), nil
	}
	if err := models.FindUserByUid(c, &user); err != nil {
		return response.CreateResponse(constanct.MySQLErrorCode), err
	}
	if ok := models.EqualPassWord(c, &user, req.Pass); !ok {
		return response.CreateResponse(constanct.PassWordErrorCode), nil
	}
	token, err := middlewares.GetToken(c, user.Uid)
	if err != nil {
		logger.Errorf("call GetToken failed, err=%s", err.Error())
		return response.CreateResponse(constanct.TokenBuildErrorCode), nil
	}
	permission, err := mysqldao.SelectPermissionByUid(c, user.Uid)
	if err != nil {
		return response.CreateResponse(constanct.MySQLErrorCode), err
	}
	return response.LoginResp{
		Response: response.Response{
			StatusCode: constanct.SuccessCode,
			StatusMsg:  constanct.SuccessCode.Msg(),
		},
		Token: token,
		Uname: user.Uname,
		Permission: response.Permission{
			Administrator:   permission.Administrator == "Y",
			Problem_edit:    permission.Problem_edit == "Y",
			Source_browser:  permission.Source_browser == "Y",
			Contest_creator: permission.Contest_creator == "Y",
		},
	}, nil
}
func DoResiger(c *gin.Context, req *request.User) (interface{}, error) {
	logger := utils.GetLogInstance()
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
	//检测用户合法性

	//查看用户账号是否存在
	exist := models.IsUserExistByUid(c, &user)
	if exist {
		return response.CreateResponse(constanct.UIDExistCOde), nil
	}
	// 2、密码加密处理（MD5)
	user.Pass, _ = utils.MD5EnCode(req.Uid, req.Pass)
	// 3、创建用户
	err := models.CreateUser(c, &user)
	if err != nil {
		logger.Errorf("call CreateUser failed,err=%s", err.Error())
		return response.CreateResponse(constanct.MySQLErrorCode), err
	}
	// 4、返回注册成功的信息给用户
	return response.CreateResponse(constanct.SuccessCode), nil
}
func GetUserInfo(c *gin.Context, req *string) (interface{}, error) {
	user := dao.User{
		Uid: *req,
	}
	exist := models.IsUserExistByUid(c, &user)
	if !exist {
		return response.CreateResponse(constanct.UIDNotExistCode), nil
	}
	models.FindUserByUid(c, &user)
	return response.CreateUserResp(&user), nil
}
