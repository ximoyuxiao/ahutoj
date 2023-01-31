package response

import "ahutoj/web/io/constanct"

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
	PID        string `json:"PID"`
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
	Problems    string       `json:"Problems"`
	ProblemData []ConProItem `json:"Data"`
}
type ContestListResp struct {
	Response
	Size int64             `json:"Size"`
	Data []ContestListItem `json:"Data"`
}
type ProblemItemWithACM struct {
	PID          string             `json:"PID"`          // 题目ID 其实我觉得这个可以不写的
	Time         int64              `json:"Time"`         // 最后一次提交时间
	SubmitNumber int64              `json:"SubmitNumber"` // 题目总的提交次数
	Status       constanct.OJResult `json:"Status"`       // 最终状态
}
type ProblemItemWithOI struct {
	PID        string `json:"PID"`      // 题目ID 其实我觉得这个可以不写的
	Time       int64  `json:"Time"`     // 提交使用时间
	Submited   bool   `json:"Submited"` //是否提交过
	PassSample int64  `json:"Score"`    // 最终状态
}
type RankItemWithAcm struct {
	UserID           string               `json:"UserID"`
	Uname            string               `json:"Uname"`
	Uclass           string               `json:"Uclass"`
	AllSubmit        int64                `json:"AllSubmit"`
	ACNumber         int64                `json:"ACNumber"`
	CENumber         int64                `json:"CENumber"`
	JudgeErrorNumber int64                `json:"JudgeErrorNumber"`
	Problems         []ProblemItemWithACM `json:"Problems"`
}

type RankItemWithOI struct {
	UserID   string              `json:"UserID"`
	Uname    string              `json:"Uname"`
	Uclass   string              `json:"Uclass"`
	ACNumber int64               `json:"ACNumber"`
	Problems []ProblemItemWithOI `json:"Problems"`
}

type ConntestRankRespWithAcm struct {
	Response
	Size int               `json:"Size"`
	Data []RankItemWithAcm `json:"Data"`
}

type ConntestRankRespWithOI struct {
	Response
	Size int              `json:"Size"`
	Data []RankItemWithOI `json:"Data"`
}
type RankItemsWithAcm []RankItemWithAcm
type RankItemsWithOI []RankItemWithOI

func (r RankItemsWithAcm) Len() int {
	return len(r)
}
func (r RankItemsWithOI) Len() int {
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
