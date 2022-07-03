package request

type LoginReq struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
