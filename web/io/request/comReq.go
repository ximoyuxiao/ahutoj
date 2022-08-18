package request

type LoginReq struct {
	Uid  string `json:"uid"`
	Pass string `json:"pass"`
}

type GetListReq struct {
	Page  int `query:"page"`
	Limit int `query:"limit"`
}
