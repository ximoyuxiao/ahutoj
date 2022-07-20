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

var passSuffix = []string{".io", ".out", ".zip"}

func checkSuccessFile(filename string) bool {
	filenames := strings.Split(filename, ".")
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
	pid := ctx.Param("id")
	//判断文件夹存在 这要求我们创建题目的时候 必须创建一个对应的文件夹
	return utils.GetConfInstance().DataPath + "/" + pid
}
func UpFile(ctx *gin.Context) {
	pid := ctx.Param("id")
	//判断文件夹存在 这要求我们创建题目的时候 必须创建一个对应的文件夹
	path := utils.GetConfInstance().DataPath + "/" + pid
	ok, _ := pathExists(path)
	if !ok {
		response.ResponseError(ctx, constanct.NotLoginCode)
	}

	//检查文件正确
	file, _ := ctx.FormFile("file")
	if !checkSuccessFile(file.Filename) {
		response.ResponseError(ctx, constanct.NotLoginCode)
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
	}
	filepath := path + "/" + filename
	err := os.Remove(filepath)
	if err != nil {
		response.ResponseError(ctx, constanct.MySQLErrorCode)
	}
	response.ResponseOK(ctx, response.CreateResponse(constanct.SuccessCode))
}

func UnzipFile(ctx *gin.Context) {
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
	cmd := exec.Command("unzip", filepath, "-d ./")
	cmd.Start()
	response.ResponseOK(ctx, response.CreateResponse(constanct.SuccessCode))
}
