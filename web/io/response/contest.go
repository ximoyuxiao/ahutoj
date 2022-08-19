package response

type ContestListItem struct {
	Cid        int64  `json:"cid"`
	Uid        string `json:"uid"`
	Title      string `json:"title"`
	Begin_time int64  `json:"begin_time"`
	End_time   int64  `json:"end_time"`
	Ctype      int    `json:"ctype"`
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
	Ctype       int          `json:"ctype"`
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
type ProblemItem struct {
	Pid        int   `json:"pid"`
	Time       int64 `json:"time"`
	Status     int64 `json:"status"`
	Submit_num int64 `json:"submitNum"`
}
type RankItem struct {
	Rank     int64         `json:"rank"`
	UserID   string        `json:"userID"`
	Uname    string        `json:"uname"`
	Solve    int64         `json:"solve"`
	Penalty  int64         `json:"penalty"`
	Mark     int64         `json:"mark"`
	Problems []ProblemItem `json:"probelms"`
}
type ConntestRankResp struct {
	Response
	Size int        `json:"size"`
	Data []RankItem `json:"Data"`
}
type RankItems []RankItem

func (r RankItems) Len() int {
	return len(r)
}

// 实现sort.Interface接口的比较元素方法
func (m RankItems) Less(i, j int) bool {
	if m[i].Solve == m[j].Solve {
		return m[i].Penalty < m[j].Penalty
	}
	return m[i].Solve > m[j].Solve
}

// 实现sort.Interface接口的交换元素方法
func (m RankItems) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
