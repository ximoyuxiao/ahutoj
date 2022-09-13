package service

import (
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/response"
	"ahutoj/web/utils"
	"os"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

var passSuffix = []string{"in", "out", "zip"}

func checkSuccessFile(filename string) bool {
	logger := utils.GetLogInstance()
	filenames := strings.Split(filename, ".")
	logger.Debug("filenames:%+v", filenames)
	if len(filenames) == 1 {
		return false
	}
	suffix := filenames[len(filenames)-1]
	for _, data := range passSuffix {
		if data == suffix {
			return true
		}
	}
	return false
}
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func getPath(ctx *gin.Context) string {
	pid := ctx.Param("pid")
	//判断文件夹存在 这要求我们创建题目的时候 必须创建一个对应的文件夹
	return utils.GetConfInstance().DataPath + "/" + pid
}
func UpFile(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	pid := ctx.Param("pid")
	//判断文件夹存在 这要求我们创建题目的时候 必须创建一个对应的文件夹
	path := utils.GetConfInstance().DataPath + "/" + pid
	logger.Debug("path:%s", path)
	ok, _ := pathExists(path)
	if !ok {
		logger.Errorf("not exists path:%s", path)
		response.ResponseError(ctx, constanct.FILEUNSUPPORT)
		return
	}

	//检查文件正确
	file, err := ctx.FormFile("file")
	if err != nil {
		logger.Errorf("call FormFile filed, err=%s", err.Error())
		response.ResponseError(ctx, constanct.FILEUNSUPPORT)
		return
	}
	logger.Infof("upfile:%s", file.Filename)
	if !checkSuccessFile(file.Filename) {
		logger.Errorf("chekfile failed filename:%s", file.Filename)
		response.ResponseError(ctx, constanct.FILEUNSUPPORT)
		return
	}
	//SaveUploadedFile上传表单文件到指定的路径
	ctx.SaveUploadedFile(file, path+"/"+file.Filename)
	response.ResponseOK(ctx, response.CreateResponse(constanct.SuccessCode))
}

func RemoveFile(ctx *gin.Context) {
	path := getPath(ctx)
	filename := ctx.PostForm("file")
	if filename == "" {
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	filepath := path + "/" + filename
	err := os.Remove(filepath)
	if err != nil {
		response.ResponseError(ctx, constanct.FILEUNSUPPORT)
		return
	}
	response.ResponseOK(ctx, response.CreateResponse(constanct.SuccessCode))
}

func UnzipFile(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	path := getPath(ctx)
	ok, _ := pathExists(path)
	if !ok {
		response.ResponseError(ctx, constanct.NotLoginCode)
	}
	filename := ctx.PostForm("file")
	if filename == "" {
		response.ResponseError(ctx, constanct.InvalidParamCode)
	}
	filepath := path + "/" + filename
	logger.Infof("带解压文件:%s", filepath)
	cmd := exec.Command("unzip", filepath, "-d", path)
	cmd.Start()
	response.ResponseOK(ctx, response.CreateResponse(constanct.SuccessCode))
}

func UpProblemFile(ctx *gin.Context) {
	response.ResponseOK(ctx, response.CreateResponse(constanct.Notimplemented))
}
