package logic

import (
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/response"
	"ahutoj/web/utils"
	"os"

	"github.com/gin-gonic/gin"
)

func BuildFilePathResp(ctx *gin.Context, filepath string) (interface{}, error) {
	logger := utils.GetLogInstance()
	files, err := os.ReadDir(filepath)
	if err != nil {
		logger.Errorf("call ReadDir faile,filepath=%s err=%s", filepath, err.Error())
		return response.CreateResponse(constanct.InvalidParamCode), nil
	}
	resp := response.GetFileListResp{}
	resp.Data = make([]response.FileItem, 0)
	for _, file := range files {
		fileinfo, _ := file.Info()
		resp.Data = append(resp.Data, response.FileItem{
			Filename: file.Name(),
			FileSize: fileinfo.Size(),
			FileType: utils.GetFileType(file.Name()),
		})
	}
	response.CreateResponse(constanct.SuccessCode)
	return resp, nil
}
