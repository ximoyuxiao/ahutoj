package request

type LoginReq struct {
	Uid  string `json:"uid"`
	Pass string `json:"pass"`
}
