package response

type UserResp struct {
	Response
	Token string `json:"token"`
}
