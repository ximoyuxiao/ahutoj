package service

import (
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/response"
	"ahutoj/web/utils"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

var passSuffix = []string{"in", "out", "zip"}

func checkSuccessFile(filename string) bool {
	logger := utils.GetLogInstance()
	filenames := strings.Split(filename, ".")
	logger.Debugf("filenames:%+v", filenames)
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
	if pid == "" {
		return ""
	}
	//判断文件夹存在 这要求我们创建题目的时候 必须创建一个对应的文件夹
	return utils.GetConfInstance().DataPath + "/" + pid
}
func UpFile(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	pid := ctx.Param("pid")
	//判断文件夹存在 这要求我们创建题目的时候 必须创建一个对应的文件夹
	path := utils.GetConfInstance().DataPath + "/" + pid
	logger.Debugf("path:%s", path)
	ok, _ := pathExists(path)
	if !ok {
		logger.Errorf("not exists path:%s", path)
		os.Mkdir(path, os.ModeDir)
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
	logger := utils.GetLogInstance()
	filepath := path + "/" + filename
	logger.Debugf("待删除文件:%s", filepath)
	err := os.Remove(filepath)
	if err != nil {
		logger.Errorf("call Remove failed, err=%s", err.Error())
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
func GetFileType(filename string) string {
	return ""
}
func GetFileList(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	filepath := getPath(ctx)
	if filepath == "" {
		logger.Errorf("has no pid Invailed")
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	files, err := ioutil.ReadDir(filepath)
	if err != nil {
		logger.Errorf("call ReadDir faile,filepath=%s err=%s", filepath, err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp := response.GetFileListResp{}
	resp.Data = make([]response.FileItem, 0)
	for _, file := range files {
		resp.Data = append(resp.Data, response.FileItem{
			Filename: file.Name(),
			FileSize: file.Size(),
			FileType: GetFileType(file.Name()),
		})
	}
	resp.Response = response.CreateResponse(constanct.SuccessCode)
	resp.Count = len(resp.Data)
	response.ResponseOK(ctx, resp)
}
