package response

type LoginResp struct {
	Response
	Token string `json:"token"`
	Uname string `json:"name"`
	Permission
}
