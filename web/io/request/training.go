package request

type DelListReq struct {
	LID []int64 `json:"LID"`
}

type ListProblem struct {
	LID int64  `json:"LID"`
	PID string `json:"PID"`
}
type ListUserReq struct {
	LID int64 `query:"LID"`
}
type ListAll struct {
	Problems    string `json:"Problems"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}
type EditListReq struct {
	LID         int64  `json:"LID"`
	Description string `json:"Description"`
	Title       string `json:"Title"`
	Problems    string `json:"Problems"`
}
type RegisterTrainingReq struct {
	LID int64  `json:"LID"`
	UID string `json:"UID"`
}

type TrainingListReq GetListReq
type TrainingReq struct {
	LID int64 `query:"LID"`
}

type GetTrainingRankReq struct {
	TrainingListReq
	TrainingReq
}
