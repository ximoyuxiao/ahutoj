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
