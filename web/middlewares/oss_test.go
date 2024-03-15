package middlewares

import (
	"ahutoj/web/utils"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"log"
	"testing"
)

func TestOss(t *testing.T) {
	utils.ConfigInit("../../config/config.yaml")
	osscfg := utils.GetConfInstance().OssConfig
	oss, err := NewOss(osscfg.Host, osscfg.Port, osscfg.AccessKeyID, osscfg.SecretAccessKey, osscfg.UseSSL)
	if err != nil {
		log.Fatalf("failed to create oss: %v", err)
	}
	BucketName := "test"
	ctx := context.Background()
	err = oss.MakeBucket(ctx, BucketName, minio.MakeBucketOptions{Region: "cn-north-1"})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := oss.BucketExists(ctx, BucketName)
		if errBucketExists == nil && exists {
			fmt.Printf("We already own bucket %s\n", BucketName)
		} else {
			fmt.Printf("Failed to create bucket %s: %v\n", BucketName, err)
			return
		}
	} else {
		fmt.Printf("Successfully created bucket %s\n", BucketName)
	}
	buckets, err := oss.ListBuckets(ctx)
	if err != nil {
		fmt.Printf("Failed to list buckets: %v\n", err)
		return
	}
	for _, bucket := range buckets {
		objects := oss.ListObjects(ctx, bucket.Name, minio.ListObjectsOptions{})
		for object := range objects {
			fmt.Println(object.Key)
		}
	}

}
