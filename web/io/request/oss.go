package request

import "ahutoj/web/io/constanct"

type GetObjectReq struct {
	Type     constanct.DataType `json:"Type"`
	FileName string             `json:"FileName"`
}
type GetObjectsReq struct {
	FilePath string `json:"FileName"`
}
type UpObjectReq struct {
	TargetBucket   string `json:"TargetPath"`
	TargetFileName string `json:"TargetFileName"`
	Data           []byte `json:"Data"`
}

type ModifyObjectReq struct {
}

type DeleteObjectReq struct {
}

type InfoObjectReq struct {
}
