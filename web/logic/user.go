package logic

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/mapping"
	"ahutoj/web/middlewares"
	"ahutoj/web/models"
	originjudge "ahutoj/web/originJudge"
	"ahutoj/web/utils"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func CheckLogin(req *request.LoginReq, c *gin.Context) (interface{}, error) {
	logger := utils.GetLogInstance()
	user := dao.User{
		UID: req.UID,
	}
	if req.UID == "" {
		return response.CreateResponseStr(constanct.GetResCode(constanct.User, constanct.Logic, constanct.ParametersFormatError), "UID不能为空", response.ERROR), nil
	}
	if req.Pass == "" {
		return response.CreateResponseStr(constanct.GetResCode(constanct.User, constanct.Logic, constanct.ParametersFormatError), "Pass不能为空", response.ERROR), nil
	}
	if ok := models.IsUserExistByUID(c, &user); !ok {
		return response.CreateResponseStr(constanct.GetResCode(constanct.User, constanct.Logic, constanct.UIDNotExist), "UID不存在", response.ERROR), nil
	}
	if err := models.FindUserByUID(c, &user); err != nil {
		return response.CreateResponseStr(constanct.GetResCode(constanct.User, constanct.Logic, constanct.DataEmpty), "找不到指定用户", response.ERROR), err
	}
	if ok := models.EqualPassWord(c, &user, req.Pass); !ok {
		return response.CreateResponseStr(constanct.GetResCode(constanct.User, constanct.Logic, constanct.PasswordError), "密码错误", response.ERROR), nil
	}
	token, err := middlewares.GetToken(c, user.UID)
	if err != nil {
		logger.Errorf("call GetToken failed, err=%s", err.Error())

		return response.CreateResponseStr(constanct.GetResCode(constanct.User, constanct.Logic, constanct.TokenBuildError), "Token创建错误", response.ERROR), nil
	}
	permission, err := mysqldao.SelectPermissionByUID(c, user.UID)
	if err != nil {
		return response.CreateResponseStr(constanct.GetResCode(constanct.User, constanct.Logic, constanct.MysqlQuery), "Sql查询错误", response.ERROR), err
	}
	return response.LoginResp{
		Response: response.Response{
			StatusCode: constanct.SuccessCode,
			// StatusMsg:  constanct.SuccessCode.Msg(),
		},
		Token: token,
		Uname: user.Uname,
		Permission: response.Permission{
			PermissionMap: mapping.PermissionToBitMap(permission),
		},
	}, nil
}
func DoResiger(c *gin.Context, req *request.User) (interface{}, error) {
	logger := utils.GetLogInstance()
	user := dao.User{
		UID:     req.UID,
		Uname:   req.Uname,
		Pass:    req.Pass,
		School:  req.School,
		Classes: req.Classes,
		Adept:   req.Adept,
		Major:   req.Major,
		Vjid:    req.Vjid,
		Vjpwd:   req.Vjpwd,
		Email:   req.Email,
	}
	//检测用户合法性

	//查看用户账号是否存在
	exist := models.IsUserExistByUID(c, &user)
	if exist {
		return response.CreateResponse(constanct.GetResCode(constanct.User, constanct.Logic, constanct.UIDExist)), nil
	}
	// 创建用户
	err := models.CreateUser(c, &user)
	if err != nil {
		logger.Errorf("call CreateUser failed,err=%s", err.Error())
		return response.CreateResponseStr(constanct.GetResCode(constanct.User, constanct.Logic, constanct.MysqlAdd), "用户创建错误", response.ERROR), err
	}
	// 获取token
	token, err := middlewares.GetToken(c, req.UID)
	if err != nil {
		logger.Errorf("call GetToken failed, err=%s", err.Error())
		return response.CreateResponseStr(constanct.GetResCode(constanct.User, constanct.Logic, constanct.TokenBuildError), "Token创建错误", response.ERROR), nil
	}
	permission, err := mysqldao.SelectPermissionByUID(c, user.UID)
	if err != nil {
		return response.CreateResponseStr(constanct.GetResCode(constanct.User, constanct.Logic, constanct.MysqlQuery), "Sql查询错误", response.ERROR), err
	}
	// 4、返回注册成功的信息给用户
	return response.RegisterResp{
		Response: response.Response{
			StatusCode: constanct.SuccessCode,
			// StatusMsg:  constanct.SuccessCode.Msg(),
		},
		Token: token,
		Uname: user.Uname,
		Permission: response.Permission{
			PermissionMap: mapping.PermissionToBitMap(permission),
		},
	}, nil
}
func GetUserInfo(c *gin.Context, req *string) (interface{}, error) {
	user := dao.User{
		UID: *req,
	}
	exist := models.IsUserExistByUID(c, &user)
	if !exist {
		return response.CreateResponseStr(constanct.GetResCode(constanct.User, constanct.Logic, constanct.UIDNotExist), "UID不存在", response.ERROR), nil
	}
	models.FindUserByUID(c, &user)
	return response.CreateUserResp(&user), nil
}

func UpdateUserInfo(ctx *gin.Context, req request.UserEditReq) (interface{}, error) {
	return response.CreateResponseStr(constanct.GetResCode(constanct.User, constanct.Logic, constanct.Notimplemented), "接口未实现", response.ERROR), nil
}

func UpdateUserPass(ctx *gin.Context, req request.UserEditPassReq) (interface{}, error) {
	return response.CreateResponseStr(constanct.GetResCode(constanct.User, constanct.Logic, constanct.Notimplemented), "接口未实现", response.ERROR), nil
}

func UpdateUserVjudge(ctx *gin.Context, req request.UserEditVjudgeReq) (interface{}, error) {
	return response.CreateResponseStr(constanct.GetResCode(constanct.User, constanct.Logic, constanct.Notimplemented), "接口未实现", response.ERROR), nil
}

func AddUsersRange(ctx *gin.Context, req request.AddUsersRangeReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	resp := response.AddUsersResp{}
	resp.CreateNumber = 0
	resp.Data = make([]response.UsersItem, 0)
	if req.Password == nil || *req.Password == "" {
		req.Password = new(string)
		*req.Password = "123456"
	}
	for idx := 1; idx <= req.Number; idx++ {
		UID := fmt.Sprintf("%s%02d", req.Prefix, idx)
		err := models.CreateUser(ctx, &dao.User{
			UID:    UID,
			Uname:  UID,
			Pass:   *req.Password,
			School: req.School,
		})

		if err != nil {
			logger.Errorf("call CreateUser failed,UID=%+v,err=%s", UID, err.Error())
			continue
		} else {
			resp.CreateNumber += 1
			usersItem := response.UsersItem{
				UID:      UID,
				Uname:    UID,
				Password: *req.Password,
				School:   req.School,
			}
			resp.Data = append(resp.Data, usersItem)
		}
	}
	resp.Response = response.CreateResponse(constanct.SuccessCode)
	return resp, nil
}

func AddUsers(ctx *gin.Context, req request.AddUsersReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	resp := response.AddUsersResp{}
	resp.CreateNumber = 0
	resp.Data = make([]response.UsersItem, 0)
	for _, item := range req {
		user := dao.User{
			UID:   item.UID,
			Pass:  item.Pass,
			Uname: item.UserName,
		}
		err := models.CreateUser(ctx, &user)
		if err != nil {
			logger.Errorf("call CreateUser failed,user=%+v,err=%s", user, err.Error())
			continue
		}
		resp.CreateNumber += 1
		usersItem := response.UsersItem{
			UID:      user.UID,
			Uname:    user.Uname,
			Password: user.Pass,
			School:   user.School,
		}
		resp.Data = append(resp.Data, usersItem)
	}
	return nil, nil
}

func GetUserStatusInfo(ctx *gin.Context, req request.UserStatusInfoReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	bigTime := time.Now() // 获取当前时间
	resp := response.UserStatusInfoResp{
		Response: response.CreateResponse(constanct.SuccessCode),
	}
	resp.Data = make([]response.UserStatusInfoItem, 0)
	switch req.Type {
	case constanct.Momth:
		{
			bigTime = bigTime.AddDate(0, -req.Time, 0)
		}
	case constanct.Year:
		{
			bigTime = bigTime.AddDate(-req.Time, 0, 0)
		}
	default:
		{
			bigTime = bigTime.AddDate(0, -6, 0)
		}
	}
	submit := dao.Submit{
		UID:    req.UID,
		Result: req.Result,
	}
	submits, err := models.GetUserStatusInfo(ctx, submit, bigTime.Unix())
	if err != nil {
		logger.Errorf("call GetUserStatusInfo failed, req=%+v,err=%s", utils.Sdump(req), err.Error())
		response.ResponseError(ctx, constanct.ParametersInvlidCode)
		return nil, err
	}
	for _, submit := range submits {
		temp := response.UserStatusInfoItem{
			PID:        submit.PID,
			Result:     submit.Result,
			SubmitTime: submit.SubmitTime,
		}
		resp.Data = append(resp.Data, temp)
	}

	return resp, nil
}

func CodeForceBind(ctx *gin.Context, req request.CodeForceBindReq) (interface{}, error) {
	req.CodeForcePass = strings.Trim(req.CodeForcePass, " ")
	req.CodeForceUser = strings.Trim(req.CodeForceUser, " ")
	if len(req.CodeForceUser) == 0 {
		return response.CreateResponseStr(constanct.GetResCode(constanct.User, constanct.Logic, constanct.UIDEmpty), "UID为空", response.ERROR), nil
	}
	if len(req.CodeForceUser) == 0 {
		return response.CreateResponseStr(constanct.GetResCode(constanct.User, constanct.Logic, constanct.PasswordEmpty), "PassWord为空", response.ERROR), nil
	}
	cj := originjudge.CodeForceJudge{
		Headers: originjudge.CfHeaders,
		JudgeUser: &originjudge.CFJudgeUser{
			OriginJudgeUser: originjudge.OriginJudgeUser{
				ID:       req.CodeForceUser,
				Password: req.CodeForcePass,
				Cookies:  make(map[string]string, 0),
			},
		},
	}
	err := cj.Login()
	if err != nil {
		return response.CreateResponseStr(constanct.GetResCode(constanct.User, constanct.Logic, constanct.PasswordError), "密码错误", response.ERROR), nil
	}
	/*应该有一步检查登录 暂时忽略*/
	user := dao.User{
		UID:           middlewares.GetUid(ctx),
		CodeForceUser: req.CodeForceUser,
	}

	err = mysqldao.UpdateUserByUID(ctx, &user)
	if err != nil {
		return nil, err
	}
	return response.CreateResponse(constanct.SuccessCode), nil
}
