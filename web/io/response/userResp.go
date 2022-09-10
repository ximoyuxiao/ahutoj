package response

import (
	"ahutoj/web/dao"
	"ahutoj/web/io/constanct"
)

type LoginResp struct {
	Response
	Token string `json:"token"`
	Uname string `json:"name"`
	Permission
}

type UserResp struct {
	Response
	Uid     string `json:"uid"`
	Uname   string `json:"uname"`
	School  string `json:"school"`
	Classes string `json:"classes"`
	Major   string `json:"major"`
	Adept   string `json:"adept"`
	Vjid    string `json:"vjid"`
	Email   string `json:"email"`
}

func CreateUserResp(user *dao.User) UserResp {
	return UserResp{
		Response: CreateResponse(constanct.SuccessCode),
		Uid:      user.Uid,
		Uname:    user.Uname,
		School:   user.School,
		Classes:  user.Classes,
		Major:    user.Major,
		Adept:    user.Adept,
		Vjid:     user.Vjid,
		Email:    user.Email,
	}
}
