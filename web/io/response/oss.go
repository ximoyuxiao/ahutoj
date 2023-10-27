package response

import (
	"github.com/minio/minio-go/v7"
)

type GetObjectsResp struct {
	ObjectInfo []minio.ObjectInfo `json:"ObjectInfo"`
}

//	type ObjectData struct {
//		Type     constanct.DataType `json:"Type"`
//		FileName string             `json:"Filename"`
//		Data     []FileData         `json:"Data"`
//		Body     []byte             `json:"Body"`
//	}

type GetBucketsResp struct {
	Buckets []minio.BucketInfo `json:"Buckets"`
}

type GetObjectResp struct {
	ObjectData string `json:"ObjectData"`
}

type FPutObjectResp struct {
	UpInfo minio.UploadInfo `json:"ObjectInfo"`
}

type CreateObjectResp struct {
	UpInfo minio.UploadInfo `json:"ObjectInfo"`
}

type GetObjectInfoResp struct {
	ObjectInfo minio.ObjectInfo `json:"ObjectInfo"`
}
