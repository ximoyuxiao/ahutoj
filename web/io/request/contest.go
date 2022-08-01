package request

type EditContestReso struct {
	Cid         int64  `json:"cid"`
	Uid         string `json:"uid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Begin_time  int64  `json:"begin_time"`
	End_time    int64  `json:"end_time"`
	Ctype       string `json:"ctype"`
	Ispublic    string `json:"ispublic"`
	Pass        string `json:"pass"`
}

type AddContestResp struct {
	Uid         string `json:"uid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Begin_time  int64  `json:"begin_time"`
	End_time    int64  `json:"end_time"`
	Ctype       string `json:"ctype"`
	Ispublic    string `json:"ispublic"`
	Pass        string `json:"pass"`
}

type ContestListReq GetListReq

type DeleteContestReq struct {
	Cid []int64 `json:"Cids"`
}
