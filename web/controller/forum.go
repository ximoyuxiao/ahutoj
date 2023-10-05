package controller

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func GetSolution(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.GetSolutionReq)
	SIDstr := ctx.Param("id")
	SID, err := strconv.Atoi(SIDstr)
	req.SID = int64(SID)
	if err != nil {
		logger.Errorf("call Atoi failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	db := mysqldao.GetDB(ctx)
	var solution dao.Solution
	err = db.Where(req.SID).Find(&solution).Error
	resp := response.SoultionResp{
		Response: response.CreateResponse(constanct.SuccessCode),
		SolutionList: response.SolutionResponseElement{
			Data:  GetSubCommentList(ctx, int64(req.SID)),
			Sid:   &solution.SID,
			Text:  &solution.Text,
			Title: &solution.Title,
			Uid:   &solution.UID,
		},
	}
	if err != nil {
		logger.Errorf("call AddPermission failed,req=%+v, err=%s", *req, err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}

func GetSoulutions(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.GetSolutionListReq)
	if err := ctx.ShouldBindWith(req, binding.Query); err != nil {
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp, err := GetSolutiontList(ctx, req)
	if err != nil {
		logger.Errorf("call AddPermission failed,req=%+v, err=%s", *req, err.Error())
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	response.ResponseOK(ctx, resp)
}

func GetFaviroate(ctx *gin.Context) {
	response.ResponseError(ctx, constanct.NotimplementedCode)
}
func DoFaviroate(ctx *gin.Context) {
	response.ResponseError(ctx, constanct.NotimplementedCode)
}

//func AddComment(ctx *gin.Context) {
//	response.ResponseError(ctx, constanct.NotimplementedCode)
//}

func EditComment(ctx *gin.Context) {
	response.ResponseError(ctx, constanct.NotimplementedCode)
}

func GetComment(ctx *gin.Context) {
	response.ResponseError(ctx, constanct.NotimplementedCode)
}

func GetComments(ctx *gin.Context) {
	// GetCommentList
	logger := utils.GetLogInstance()
	req := new(request.CommentListReq)
	if err := ctx.ShouldBindWith(req, binding.Query); err != nil {
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	resp := GetCommentList(ctx, req.SID)
	// if err != nil {
	// 	logger.Errorf("call AddPermission failed,req=%+v, err=%s", *req, err.Error())
	// 	response.ResponseError(ctx, constanct.ServerErrorCode)
	// 	return
	// }
	response.ResponseOK(ctx, resp)
}

//func DeleteComment(ctx *gin.Context) {
//	response.ResponseError(ctx, constanct.NotimplementedCode)
//}
