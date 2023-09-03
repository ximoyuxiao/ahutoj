package response

import "ahutoj/web/io/constanct"

type GetObjectsResp struct {
	Response
	Data FileData `json:"Data"`
}
type FileData struct {
	Type     constanct.DataType `json:"Type"`
	FileName string             `json:"Filename"`
	Data     []FileData         `json:"Data"`
	Body     []byte             `json:"Body"`
}
type UpObjectResp struct {
	TargetBucket   string `json:"TargetPath"`
	TargetFileName string `json:"TargetFileName"`
	Data           []byte `json:"Data"`
}
