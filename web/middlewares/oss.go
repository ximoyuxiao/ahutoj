package middlewares

import (
	"ahutoj/web/utils"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var oss *minio.Client

func NewOss(Host string, Port uint16, AccessKeyID string, SecretAccessKey string, UseSSL bool) (*minio.Client, error) {
	if oss != nil {
		return oss, nil
	}
	oss,err:=utils.Re(func()(*minio.Client,error){
		endpoint := fmt.Sprintf("%v:%v", Host, Port)
		return minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(AccessKeyID, SecretAccessKey, ""),
		Secure: UseSSL,
	})},3,5,"minio.New")
	return oss, err

}
func GetOss() *minio.Client {
	return oss
}
