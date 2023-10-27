package request

import "io"

type GetObjectReq struct {
	BucketName string `json:"BucketName"`
	ObjectName string `json:"ObjectName"`
}

type GetObjectsReq struct {
	BucketName string `json:"BucketName"`
}

type GetObjectInfoReq struct {
	BucketName string `json:"BucketName"`
	ObjectName string `json:"ObjectName"`
}

type FGetObjectReq struct {
	BucketName string `json:"BucketName"`
	ObjectName string `json:"ObjectName"`
	FilePath   string `json:"FilePath"`
}

type FPutObjectReq struct {
	BucketName string `json:"BucketName"`
	ObjectName string `json:"ObjectName"`
	FilePath   string `json:"FilePath"`
}

type CreateObjectReq struct {
	BucketName string    `json:"BucketName"`
	ObjectName string    `json:"ObjectName"`
	Reader     io.Reader `json:"Reader"`
	ObjectSize int64     `json:"ObjectSize,omitempty"`
}

type DeleteObjectReq struct {
	BucketName string `json:"BucketName"`
	ObjectName string `json:"ObjectName"`
}

type CreateBucketreq struct {
	BucketName string `json:"BucketName"`
}

type RemoveBucketreq struct {
	BucketName string `json:"BucketName"`
}

type UnzipReq struct {
	ObjectPath     string `json:"ObjectPath"`
	TargetFilePath string `json:"TargetPath"`
}
