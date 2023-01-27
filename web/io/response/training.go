package response

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
	Title       string                `json:"Title"`
	StartTime   int64                 `json:"StartTime"`
	Problems    string                `json:"Problems"`
	ProblemData []TrainingProblemItem `json:"Data"`
}
type TrainingProblemItem struct {
	PID    string `json:"PID"`
	Ptitle string `json:"Ptitle"`
}
