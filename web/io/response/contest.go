package response

type ContestListItem struct {
	CID       int64  `json:"CID"`
	UID       string `json:"UID"`
	Title     string `json:"Title"`
	BeginTime int64  `json:"BeginTime"`
	EndTime   int64  `json:"EndTime"`
	Type      int    `json:"Type"`
	Ispublic  int    `json:"IsPublic"`
}
type ConProItem struct {
	PID        int    `json:"PID"`
	Ptitle     string `json:"Title"`
	Submit_num int    `json:"SubmitNum"`
	Ac_num     int    `json:"ACNum"`
}
type GetContestResp struct {
	Response
	CID         int64        `json:"CID"`
	UID         string       `json:"UID"`
	Title       string       `json:"Title"`
	Description string       `json:"Description"`
	Begin_time  int64        `json:"BeginTime"`
	End_time    int64        `json:"EndTime"`
	Ctype       int          `json:"Type"`
	Ispublic    int          `json:"IsPublic"`
	Size        int64        `json:"Size"`
	ProblemData []ConProItem `json:"Data"`
}
type ContestListResp struct {
	Response
	Size int64             `json:"Size"`
	Data []ContestListItem `json:"Data"`
}
type ProblemItem struct {
	PID    int   `json:"PID"`
	Time   int64 `json:"Time"`
	Status int64 `json:"Status"`
}
type RankItem struct {
	UserID   string        `json:"UserID"`
	Uname    string        `json:"Uname"`
	Problems []ProblemItem `json:"Probelms"`
}
type ConntestRankResp struct {
	Response
	Size int        `json:"Size"`
	Data []RankItem `json:"Data"`
}
type RankItems []RankItem

func (r RankItems) Len() int {
	return len(r)
}

// 实现sort.Interface接口的比较元素方法
// func (m RankItems) Less(i, j int) bool {
// 	if m[i].Solve == m[j].Solve {
// 		return m[i].Penalty < m[j].Penalty
// 	}
// 	return m[i].Solve > m[j].Solve
// }

// 实现sort.Interface接口的交换元素方法
// func (m RankItems) Swap(i, j int) {
// 	m[i], m[j] = m[j], m[i]
// }
