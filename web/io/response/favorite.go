package response

type FavoriteAction struct {
	Response
	IsFavorite bool `json:"IsFavorite"`
	Count      int  `json:"Count"`
}
