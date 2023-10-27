package logic

import (
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/middlewares"
	"ahutoj/web/utils"
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"io"
)

//	func BuildGetObjectsRespnse(path string, Data *response.FileData) error {
//		// 获取路径的文件信息
//		fileInfo, err := os.Stat(path)
//		if err != nil {
//			return err
//		}
//		// 如果路径是一个目录，则递归读取其子目录和文件
//		if fileInfo.IsDir() {
//			files, err := ioutil.ReadDir(path)
//			if err != nil {
//				return err
//			}
//			*Data = response.FileData{
//				Type:     constanct.DIR,
//				FileName: fileInfo.Name(),
//				Data:     make([]response.FileData, 0),
//			}
//			for _, file := range files {
//				filePath := path + "/" + file.Name()
//				fmt.Println(filePath)
//				data := response.FileData{}
//				err := BuildGetObjectsRespnse(filePath, &data)
//				if err != nil {
//					return err
//				}
//				Data.Data = append(Data.Data, data)
//			}
//			return nil
//		}
//		// 如果路径是一个文件，则读取其内容
//		fileContent, err := ioutil.ReadFile(path)
//		if err != nil {
//			return err
//		}
//		*Data = response.FileData{
//			Type:     constanct.FILE,
//			FileName: fileInfo.Name(),
//			Data:     nil,
//			Body:     fileContent,
//		}
//		return nil
//	}
//
// --------------------对象------------------
func GetObjects(ctx *gin.Context, req *request.GetObjectsReq) (interface{}, error) {
	oss := middlewares.GetOss()
	logger := utils.GetLogInstance()
	c, cancel := context.WithCancel(ctx.Request.Context())
	defer cancel() //这里不懂，研究一下
	objectCh := oss.ListObjects(c, req.BucketName, minio.ListObjectsOptions{})
	var (
		objects []minio.ObjectInfo
		err     error
	)
	for object := range objectCh {
		err = object.Err
		if err != nil {
			logger.Errorf("call ListObjects failed,param %v,err:%v", req.BucketName, object.Err)
		}
		objects = append(objects, object)
	}
	return &response.GetObjectsResp{ObjectInfo: objects}, err
}

func FGetObject(ctx *gin.Context, req *request.FGetObjectReq) (interface{}, error) {
	oss := middlewares.GetOss()
	logger := utils.GetLogInstance()
	err := oss.FGetObject(ctx.Request.Context(), req.BucketName, req.ObjectName, req.FilePath, minio.GetObjectOptions{})
	if err != nil {
		logger.Errorf("call FGetObject failed,param %v,err:%v", req.ObjectName, err)
	}
	return response.CreateResponse(constanct.SuccessCode), err
}

//func CreateObject(ctx *gin.Context, req *request.CreateObjectReq) (interface{}, error) {
//	oss := middlewares.GetOss()
//	logger := utils.GetLogInstance()
//	//todo
//	return &response.CreateObjectResp{}, err
//}

func FPutObject(ctx *gin.Context, req *request.FPutObjectReq) (interface{}, error) {
	oss := middlewares.GetOss()
	logger := utils.GetLogInstance()
	uploadInfo, err := oss.FPutObject(ctx.Request.Context(), req.BucketName, req.ObjectName, req.FilePath, minio.PutObjectOptions{})
	if err != nil {
		logger.Errorf("call FPutObject failed,param %v,err:%v", req.ObjectName, err)
	}
	return &response.FPutObjectResp{UpInfo: uploadInfo}, err
}

func GetObject(ctx *gin.Context, req *request.GetObjectReq) (interface{}, error) {
	oss := middlewares.GetOss()
	logger := utils.GetLogInstance()
	object, err := oss.GetObject(ctx.Request.Context(), req.BucketName, req.ObjectName, minio.GetObjectOptions{})
	if err != nil {
		logger.Errorf("call GetObject failed,param bucketname %v objectname %v,err:%v", req.BucketName, req.ObjectName, err)
	}
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, object)
	if err != nil {
		logger.Errorf("call BufCopy failed")
	}
	b64 := base64.StdEncoding.EncodeToString(buf.Bytes())
	return &response.GetObjectResp{ObjectData: b64}, err
}

func DeleteObject(ctx *gin.Context, req *request.DeleteObjectReq) (interface{}, error) {
	oss := middlewares.GetOss()
	logger := utils.GetLogInstance()
	err := oss.RemoveObject(ctx.Request.Context(), req.BucketName, req.ObjectName, minio.RemoveObjectOptions{})
	if err != nil {
		logger.Errorf("call RemoveObject failed,param bucketname %v objectname %v,err:%v", req.BucketName, req.ObjectName, err)
	}

	return response.CreateResponse(constanct.SuccessCode), err
}

//func CreateObject(ctx *gin.Context, req *request.CreateObjectReq) (interface{}, error) {
//	oss := middlewares.GetOss()
//	logger := utils.GetLogInstance()
//
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	uploadInfo, err := minioClient.PutObject(context.Background(), "mybucket", "myobject", file, fileStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
//	uploadInfo, err := oss.PutObject(ctx.Request.Context(), req.BucketName, req.ObjectName, minio.GetObjectOptions{})
//	if err != nil {
//		logger.Errorf("call GetObject failed,param bucketname %v objectname %v,err:%v", req.BucketName, req.ObjectName, err)
//	}
//
//}

func GetObjectInfo(ctx *gin.Context, req *request.GetObjectInfoReq) (interface{}, error) {
	oss := middlewares.GetOss()
	logger := utils.GetLogInstance()
	objectInfo, err := oss.StatObject(ctx.Request.Context(), req.BucketName, req.ObjectName, minio.StatObjectOptions{})
	if err != nil {
		logger.Errorf("call StatObject failed,param bucketname %v objectname %v,err:%v", req.BucketName, req.ObjectName, err)
	}

	return &response.GetObjectInfoResp{ObjectInfo: objectInfo}, err
}

//func UnzipObject(ctx *gin.Context, req *request.UnzipReq) (interface{}, error) {
//	cmd := exec.Command("unzip", req.ObjectPath, "-d", req.TargetFilePath)
//	cmd.Start()
//	response.ResponseOK(ctx, response.CreateResponse(constanct.SuccessCode))
//	return response.CreateResponse(constanct.SuccessCode), nil
//}

// -------------------桶----------------------
func GetBucket(ctx *gin.Context) (interface{}, error) {
	logger := utils.GetLogInstance()
	buckets, err := middlewares.GetOss().ListBuckets(ctx.Request.Context())
	if err != nil {
		logger.Errorf("call ListBuckets failed")
	}
	return &response.GetBucketsResp{Buckets: buckets}, err
}

func CreateBucket(ctx *gin.Context, req *request.CreateBucketreq) (interface{}, error) {
	oss := middlewares.GetOss()
	logger := utils.GetLogInstance()
	err := oss.MakeBucket(ctx.Request.Context(), req.BucketName, minio.MakeBucketOptions{Region: "cn-north-1"})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := oss.BucketExists(ctx, req.BucketName)
		if errBucketExists == nil && exists {
			logger.Errorf("We already own bucket %s\n", req.BucketName)
		} else {
			logger.Errorf(fmt.Sprintf("%v", err))
		}
	} else {
		logger.Info("Successfully created %s\n", req.BucketName)
	}
	return response.CreateResponse(constanct.SuccessCode), err
}
func RemoveBucket(ctx *gin.Context, req *request.RemoveBucketreq) (interface{}, error) {
	oss := middlewares.GetOss()
	logger := utils.GetLogInstance()
	err := oss.RemoveBucket(ctx.Request.Context(), req.BucketName)
	if err != nil {
		logger.Errorf("call RemoveBucket failed,param %v ", req.BucketName)
	}
	return response.CreateResponse(constanct.SuccessCode), err
}
