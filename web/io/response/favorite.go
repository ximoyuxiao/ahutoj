package response

type FavoriteActionResp struct {
	Response
	IsFavorite bool  `json:"IsFavorite"`
	Count      int64 `json:"Count"`
}
