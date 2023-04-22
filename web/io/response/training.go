package response

import "ahutoj/web/io/constanct"

type TrainingListItem struct {
	LID       int64  `json:"LID"`
	UID       string `json:"UID"`
	Title     string `json:"Title"`
	StartTime int64  `json:"StartTime"`
}
type TrainingListResp struct {
	Response
	Size int64              `json:"size"`
	Data []TrainingListItem `json:"data"`
}
type GetTrainResp struct {
	Response
	LID         int64                 `json:"LID"`
	UID         string                `json:"UID"`
	Description string                `json:"Description"`
	Title       string                `json:"Title"`
	StartTime   int64                 `json:"StartTime"`
	Problems    string                `json:"Problems"`
	ProblemData []TrainingProblemItem `json:"Data"`
}
type TrainingProblemItem struct {
	Sort   int    `json:"Sort"`
	PID    string `json:"PID"`
	Ptitle string `json:"Ptitle"`
}
type AddTrainResp struct {
	Response
	LID int64 `json:"LID"`
}

type TrainingUserResp struct {
	Response
	Solved    int      `json:"Solved"`
	Submited  int      `json:"Submited"`
	SolvedPID []string `json:"SolvedPID"`
}
type TrainingRankProblemItem struct {
	PID          string             `json:"PID"`
	Time         uint64             `json:"Time"`
	SubmitNumber int64              `json:"SubmitNumber"`
	Status       constanct.OJResult `json:"Status"`
}
type TrainingRankProblemItems []TrainingRankProblemItem
type TraininngRankItem struct {
	UID      string                   `json:"UID"`
	Uname    string                   `json:"Uname"`
	Uclass   string                   `json:"Uclass"`
	Solved   int64                    `json:"Solved"`
	Problems TrainingRankProblemItems `json:"Problems"`
}

type TraininngRankItems []TraininngRankItem
type TrainingRankResp struct {
	Response
	Size int64              `json:"Size"`
	Data TraininngRankItems `json:"Data"`
}
