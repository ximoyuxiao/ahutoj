package response

type ContestListItem struct {
	Cid        int64  `json:"cid"`
	Uid        string `json:"uid"`
	Title      string `json:"title"`
	Begin_time int64  `json:"begin_time"`
	End_time   int64  `json:"end_time"`
	Ctype      string `json:"ctype"`
	Ispublic   string `json:"ispublic"`
}
type ConProItem struct {
	Pid        int    `json:"pid"`
	Ptitle     string `json:"title"`
	Submit_num int    `json:"submit_num"`
	Ac_num     int    `json:"ac_num"`
}
type GetContestResp struct {
	Response
	Cid         int64        `json:"cid"`
	Uid         string       `json:"uid"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Begin_time  int64        `json:"begin_time"`
	End_time    int64        `json:"end_time"`
	Ctype       string       `json:"ctype"`
	Ispublic    string       `json:"ispublic"`
	Pass        string       `json:"pass"`
	Size        int64        `json:"size"`
	ProblemData []ConProItem `json:"Data"`
}
type ContestListResp struct {
	Response
	Size int64             `json:"size"`
	Data []ContestListItem `json:"data"`
}
