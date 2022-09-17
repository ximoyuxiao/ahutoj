package request

type EditContestReq struct {
	CID         int64  `json:"CID"`
	UID         string `json:"UID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Begin_time  int64  `json:"BeginTime"`
	End_time    int64  `json:"EndTime"`
	Ctype       int    `json:"Type"`
	Ispublic    int    `json:"IsPublic"`
	Pass        string `json:"Pass"`
	Problems    string `json:"Problems"`
}

type AddContestReq struct {
	UID         string `json:"UID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Begin_time  int64  `json:"BeginTime"`
	End_time    int64  `json:"EndTime"`
	Ctype       int    `json:"Type"`
	Ispublic    int    `json:"IsPublic"`
	Pass        string `json:"Pass"`
	Problems    string `json:"Problems"`
}

type ContestListReq GetListReq

type DeleteContestReq struct {
	CID int64 `json:"CID"`
}

type GetContestReq struct {
	CID  int64   `param:"CID"`
	Pass *string `query:"Pass"`
}

type GetContestRankReq struct {
	ContestListReq
	GetContestReq
}
