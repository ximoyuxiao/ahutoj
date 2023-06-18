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
	Perm           uint8  `json:"Perm"`
}

type ModifyObjectReq struct {
}

type DeleteObjectReq struct {
}

type InfoObjectReq struct {
}

type GetBucketReq struct {
	BucketPath string `json:"Bucketpath"`
}
type CreateBucketreq struct {
	RootBucket string `json:"Root"`
	BucketName string `json:"BucketName"`
}
type UnzipReq struct {
	ObjectPath     string `json:"ObjectPath"`
	TargetFilePath string `json:"TargetPath"`
}
