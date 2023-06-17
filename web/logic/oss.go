package logic

import (
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/utils"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
)

func BuildGetObjectsRespnse(path string, Data *response.FileData) error {
	// 获取路径的文件信息
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}
	// 如果路径是一个目录，则递归读取其子目录和文件
	if fileInfo.IsDir() {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			return err
		}
		*Data = response.FileData{
			Type:     constanct.DIR,
			FileName: fileInfo.Name(),
			Data:     make([]response.FileData, 0),
		}
		fmt.Println(files)
		for _, file := range files {
			filePath := path + "/" + file.Name()
			fmt.Println(filePath)
			data := response.FileData{}
			err := BuildGetObjectsRespnse(filePath, &data)
			if err != nil {
				return err
			}
			Data.Data = append(Data.Data, data)
		}
		return nil
	}
	// 如果路径是一个文件，则读取其内容
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	*Data = response.FileData{
		Type:     constanct.FILE,
		FileName: fileInfo.Name(),
		Data:     nil,
		Body:     fileContent,
	}
	return nil
}
func GetObjects(ctx *gin.Context, req *request.GetObjectsReq) (interface{}, error) {
	root := utils.GetConfInstance().OssConfig.BasePath
	logger := utils.GetLogInstance()
	ok, err := utils.CheckPathExists(root + req.FilePath)
	if err != nil {
		logger.Errorf("call CheckPathExists failed,param %v,err:%v ", req.FilePath, err.Error())
		return nil, err
	}
	if !ok {
		logger.Errorf("not exist path %v ", req.FilePath)
		return response.CreateResponseStr(constanct.ServerErrorCode, fmt.Sprintf("not exist path %v", req.FilePath), "error"), nil
	}
	// 获取目录下所有文件
	resp := response.GetObjectsResp{
		Response: response.CreateResponse(constanct.SuccessCode),
	}
	err = BuildGetObjectsRespnse(root+req.FilePath, &resp.Data)
	return resp, err
}

func CreateObject(ctx *gin.Context, req *request.UpObjectReq) (interface{}, error) {
	root := utils.GetConfInstance().OssConfig.BasePath
	logger := utils.GetLogInstance()
	ok, err := utils.CheckPathExists(root + req.TargetBucket)
	if err != nil {
		logger.Errorf("call CheckPathExists failed,param %v,err:%v ", req.TargetBucket, err.Error())
		return nil, err
	}
	if !ok {
		logger.Errorf("not exist path %v ", req.TargetBucket)
		return nil, err
	}
	err = os.WriteFile(req.TargetBucket+req.TargetFileName, req.Data, 0666)
	if err != nil {
		logger.Errorf("call WriteFile failed,param %v,err:%v ", utils.Sdump(req), err.Error())
		return nil, err
	}
	return response.CreateResponse(constanct.SuccessCode), nil
}
