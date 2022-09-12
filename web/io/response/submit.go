package response

type GetSubmitResp struct {
	Response
	Sid        int    `json:"sid"`
	Pid        int    `json:"pid"`
	Source     string `json:"Source"`
	Lang       int    `json:"lang"`
	Result     string `json:"result"`
	UseTime    int    `json:"useTime"`
	UseMemory  int    `json:"useMemory"`
	SubmitTime int64  `json:"submitTime"`
}
type SubmitLIstItem struct {
	Sid        int    `json:"sid"`
	Pid        int    `json:"pid"`
	Lang       int    `json:"lang"`
	Result     string `json:"result"`
	UseTime    int    `json:"useTime"`
	UseMemory  int    `json:"useMemory"`
	SubmitTime int64  `json:"submitTime"`
}
type SubmitListResp struct {
	Response
	Count    int64            `json:"count"`
	LastTime int64            `json:"lastTime"`
	Data     []SubmitLIstItem `json:"Data"`
}
