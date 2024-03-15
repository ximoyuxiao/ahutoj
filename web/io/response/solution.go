package response

type SolutionPublishResp struct {
	Response
	SID int64 `json:"SID"`
}
type SoultionResp struct {
	Response
	Count        int `json:"Count"`
	SolutionList SolutionResponseElement
}
type SoultionsResp struct {
	Response
	Count        int `json:"Count"`
	SolutionList []SolutionResponseElement
}

type SolutionResponseElement struct {
	Data          []SubComment `json:"Data,omitempty"`
	Sid           *int64       `json:"SID,omitempty"`
	Text          *string      `json:"Text,omitempty"`
	Title         *string      `json:"Title,omitempty"`
	Uid           *string      `json:"UID,omitempty"`
	FavoriteCount *int64       `json:"FavoriteCount,omitempty"`
	IsFavorite    bool         `json:"IsFavorite"`
	UpdateTime    int64        `json:"UpdateTime"`
	CreateTime    int64        `json:"CreateTime"`
}

type SolutionEditResp struct {
	Title string `json:"Title"`
	Text  string `json:"Text"` //内容
}

// 题解下的评论
type SubComment struct {
	Cid  *int64  `json:"CID,omitempty"`
	FCID *int64  `json:"FCID,omitempty"`
	Uid  *string `json:"UID,omitempty"`
	//SubComment *string `json:"SubComment,omitempty"`
	Text       *string `json:"Text,omitempty"`
	UpdateTime int64   `json:"UpdateTime"`
}

type CommentListResp struct {
	Response
	Count int          `json:"Count"`
	Data  []SubComment `json:"Data"`
}
