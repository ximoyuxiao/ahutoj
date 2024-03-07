package middlewares

import (
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var oss *minio.Client

func NewOss(Host string, Port int, AccessKeyID string, SecretAccessKey string, UseSSL bool) (*minio.Client, error) {
	if oss != nil {
		return oss, nil
	}
	endpoint := fmt.Sprintf("%v:%v", Host, Port)
	var err error
	oss, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(AccessKeyID, SecretAccessKey, ""),
		Secure: UseSSL,
	})
	if err != nil {
		return nil, err
	}
	return oss, err

}
func GetOss() *minio.Client {
	return oss
}
