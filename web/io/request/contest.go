package request

type EditContestReq struct {
	Cid         int64  `json:"cid"`
	Uid         string `json:"uid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Begin_time  int64  `json:"begin_time"`
	End_time    int64  `json:"end_time"`
	Ctype       int    `json:"ctype"`
	Ispublic    int    `json:"ispublic"`
	Pass        string `json:"pass"`
	Pids        string `json:"pids"`
}

type AddContestReq struct {
	Uid         string `json:"uid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Begin_time  int64  `json:"begin_time"`
	End_time    int64  `json:"end_time"`
	Ctype       int    `json:"ctype"`
	Ispublic    int    `json:"ispublic"`
	Pass        string `json:"pass"`
	Pids        string `json:"pids"`
}

type ContestListReq GetListReq

type DeleteContestReq struct {
	Cid int64 `json:"cid"`
}

type GetContestReq struct {
	Cid  int64   `param:"cid"`
	Pass *string `query:"pass"`
}

type GetContestRankReq struct {
	ContestListReq
	GetContestReq
}
