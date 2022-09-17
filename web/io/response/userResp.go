package response

import (
	"ahutoj/web/dao"
	"ahutoj/web/io/constanct"
)

type LoginResp struct {
	Response
	Token string `json:"Token"`
	Uname string `json:"UserName"`
	Permission
}

type RegisterResp LoginResp

type UserResp struct {
	Response
	Uid     string `json:"UID"`
	Uname   string `json:"UserName"`
	School  string `json:"School"`
	Classes string `json:"Classes"`
	Major   string `json:"Major"`
	Adept   string `json:"Adept"`
	Vjid    string `json:"Vjid"`
	Email   string `json:"Email"`
}

func CreateUserResp(user *dao.User) UserResp {
	return UserResp{
		Response: CreateResponse(constanct.SuccessCode),
		Uid:      user.UID,
		Uname:    user.Uname,
		School:   user.School,
		Classes:  user.Classes,
		Major:    user.Major,
		Adept:    user.Adept,
		Vjid:     user.Vjid,
		Email:    user.Email,
	}
}
