package response

type SolutionPublishResp struct {
	Response
	SID int64 `json:"SID"`
}
type SoultionResp struct {
	Response
	Count        int `json:"count"`
	SolutionList SolutionResponseElement
}
type SoultionsResp struct {
	Response
	Count        int `json:"count"`
	SolutionList []SolutionResponseElement
}

type SolutionResponseElement struct {
	Data  []SubComment `json:"data,omitempty"`
	Sid   *int64       `json:"SID,omitempty"`
	Text  *string      `json:"text,omitempty"`
	Title *string      `json:"Title,omitempty"`
	Uid   *string      `json:"UID,omitempty"`
}
type SolutionEditResp struct {
	Title string `gorm:"column:Title"`
	Text  string `gorm:"column:Text"` //内容
}

// 题解下的评论
type SubComment struct {
	Cid  *int64  `json:"CID,omitempty"`
	Cuid *string `json:"CUID,omitempty"`
	Uid  *string `json:"UID,omitempty"`
	//SubComment *string `json:"SubComment,omitempty"`
	Text *string `json:"Text,omitempty"`
}
