package response

type FileItem struct {
	Filename string `json:"Filename"`
	FileSize int64  `json:"FileSize"`
	FileType string `josn:"FileType"`
}

type GetFileListResp struct {
	Response
	Count int        `json:"Count"`
	Data  []FileItem `json:"Data"`
}
