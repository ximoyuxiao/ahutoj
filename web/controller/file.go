package controller

import (
	"ahutoj/web/dao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/response"
	"ahutoj/web/logic"
	"ahutoj/web/models"
	"ahutoj/web/utils"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

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
		response.ResponseError(ctx, constanct.FILE_UP_UNSUPPORTCode)
		return
	}
	logger.Infof("upfile:%s", file.Filename)
	if !checkSuccessFile(file.Filename) {
		logger.Errorf("chekfile failed filename:%s", file.Filename)
		response.ResponseError(ctx, constanct.FILE_UP_UNSUPPORTCode)
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
		response.ResponseError(ctx, constanct.FILE_REMOVE_FAILEDCode)
		return
	}
	response.ResponseOK(ctx, response.CreateResponse(constanct.SuccessCode))
}

func UnzipFile(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	path := getPath(ctx)
	ok, _ := pathExists(path)
	if !ok {
		response.ResponseError(ctx, constanct.FILE_UNZIP_NotExistCode)
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

func CheckAndCreatDir(ctx *gin.Context, filepath string) error {
	logger := utils.GetLogInstance()
	ok, err := pathExists(filepath)
	if err != nil {
		logger.Errorf("call pathExists failed,filepath:%s, err=%v", filepath, err.Error())
	}
	if !ok {
		err = os.Mkdir(filepath, 0777)
		if err != nil {
			logger.Errorf("call Mkdir failed,filepath:%s, err=%v", filepath, err.Error())
			return err
		}
	}
	return nil
}

func GetFileType(filename string) string {
	strs := strings.Split(filename, ".")
	return strs[len(strs)-1]
}

func GetFileList(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	filepath := getPath(ctx)
	if filepath == "" {
		logger.Errorf("has no pid Invailed")
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	PID := ctx.Param("pid")
	ok := models.IsProblemExistByPID(ctx, &dao.Problem{PID: PID})
	if !ok {
		logger.Errorf("the problem not exist pid=%s", PID)
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	err := CheckAndCreatDir(ctx, filepath)
	if err != nil {
		logger.Errorf("call CheckAndCreatDir failed,filepath:%v,err=%v", filepath, err.Error())
		response.ResponseError(ctx, constanct.FILE_UP_FAILEDCode)
		return
	}
	files, err := os.ReadDir(filepath)
	if err != nil {
		logger.Errorf("call ReadDir faile,filepath=%s err=%s", filepath, err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp := response.GetFileListResp{}
	resp.Data = make([]response.FileItem, 0)
	for _, file := range files {
		fileinfo, _ := file.Info()
		resp.Data = append(resp.Data, response.FileItem{
			Filename: file.Name(),
			FileSize: fileinfo.Size(),
			FileType: GetFileType(file.Name()),
		})
	}
	resp.Response = response.CreateResponse(constanct.SuccessCode)
	resp.Count = len(resp.Data)
	response.ResponseOK(ctx, resp)
}

func checkImageFile(filename string) bool {
	logger := utils.GetLogInstance()
	var passImageSuffix = []string{
		"png", "jpg", "jpeg",
	}
	filenames := strings.Split(filename, ".")
	logger.Debugf("filenames:%+v", filenames)
	if len(filenames) == 1 {
		return false
	}
	suffix := filenames[len(filenames)-1]
	for _, data := range passImageSuffix {
		if data == suffix {
			return true
		}
	}
	return false
}

func getFileSuffix(filename string) string {
	strs := strings.Split(filename, ".")
	return strs[len(strs)-1]
}

func buildFileName(file *multipart.FileHeader) string {
	now := time.Now().UnixNano()
	strs := strings.Split(file.Filename, ".")
	src, _ := file.Open()
	defer src.Close()
	var bytes = make([]byte, file.Size)
	src.Read(bytes)
	md5str, _ := utils.MD5EnCodeStr(string(bytes))
	suffix := strs[len(strs)-1]
	imageName := fmt.Sprintf("%v%v.%v", md5str, now, suffix)
	return imageName
}

func UpImagefile(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	file, err := ctx.FormFile("file")
	if err != nil {
		logger.Errorf("call FormFile filed, err=%s", err.Error())
		response.ResponseError(ctx, constanct.FILE_UPIMAGE_FAILED)
		return
	}
	logger.Infof("upfile:%s", file.Filename)
	if !checkImageFile(file.Filename) {
		logger.Errorf("chekfile failed filename:%s", file.Filename)
		response.ResponseError(ctx, constanct.FILE_UPIMAGE_FAILED)
		return
	}

	imagePath := utils.GetConfInstance().ImagePath
	//SaveUploadedFile上传表单文件到指定的路径
	err = CheckAndCreatDir(ctx, imagePath)
	if err != nil {
		logger.Errorf("call CheckAndCreatDir failed imagePath:%s", imagePath)
		response.ResponseError(ctx, constanct.FILE_UPIMAGE_FAILED)
		return
	}
	name := buildFileName(file)
	ctx.SaveUploadedFile(file, imagePath+name)
	response.ResponseOK(ctx, struct {
		response.Response
		ImageURL string `json:"ImageURL"`
	}{
		Response: response.CreateResponse(constanct.SuccessCode),
		ImageURL: "../../resourse/image/" + name,
	},
	)
}

// 下载题目数据到JSON
func DownloadProblemFromJson(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := ctx.Query("PIDs")
	if req == "" {
		logger.Errorf("call Param failed")
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}

	resp, err := logic.DownloadProblemFromJson(ctx, req)
	if err != nil {
		logger.Errorf("call DownloadProblemFromJson failed, req=%+v, err=%v", utils.Sdump(req), err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	/*返回先下载文件*/
	if str, ok := resp.(string); ok {
		fileName := fmt.Sprintf("%v.json", time.Now().Unix())
		ctx.Header("Content-Type", "application/octet-stream")
		ctx.Header("Content-Disposition", "attachment; filename="+fileName) // 用来指定下载下来的文件名
		ctx.Header("Content-Transfer-Encoding", "binary")
		ctx.String(http.StatusOK, str)
	} else {
		response.ResponseOK(ctx, resp)
	}
}

func UpProblemFile(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	file, err := ctx.FormFile("file")
	if err != nil {
		logger.Errorf("call FormFile filed, err=%s", err.Error())
		response.ResponseError(ctx, constanct.FILE_UP_UNSUPPORTCode)
		return
	}
	resp, err := logic.UpProblemFile(ctx, file)
	if err != nil {
		logger.Errorf("call UpProblemFile failed, req=%+v, err=%v", utils.Sdump(file.Filename), err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}
