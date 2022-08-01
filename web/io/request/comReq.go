package request

type LoginReq struct {
	Uid  string `json:"uid"`
	Pass string `json:"pass"`
}

type GetListReq struct {
	Page  int64 `query:"page"`
	Limit int64 `query:"limit"`
}
