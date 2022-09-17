package request

type LoginReq struct {
	UID  string `json:"UID"`
	Pass string `json:"Pass"`
}

type GetListReq struct {
	Page  int `query:"Page"`
	Limit int `query:"Limit"`
}
