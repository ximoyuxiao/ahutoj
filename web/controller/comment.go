package controller

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/middlewares"
	"ahutoj/web/utils"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func GetCommentList(ctx *gin.Context, sid int64) response.CommentListResp {
	db := mysqldao.GetDB(ctx)
	var (
		comments    []dao.Comment
		refcomments response.CommentListResp
	)
	if err := db.Where("SID = ? and isDelete = 0", sid).Find(&comments).Error; err != nil {
		return response.CommentListResp{
			Response: response.CreateResponse(constanct.ADMIN_ADD_FAILED),
		}
	}
	// 使用sort.Slice函数来排序comments切片
	sort.Slice(comments, func(i, j int) bool {
		return comments[i].UpdateTime > comments[j].UpdateTime
	})
	refcomments.Count = len(comments)
	refcomments.Response = response.CreateResponse(constanct.SuccessCode)
	for _, item := range comments {
		refcomments.Data = append(refcomments.Data, response.SubComment{
			Cid:        &item.CID,
			FCID:       &item.FCID,
			Text:       &item.Text,
			Uid:        &item.UID,
			UpdateTime: item.UpdateTime,
		})
	}
	//没错误，返回
	return refcomments
}

func GetSubCommentList(ctx *gin.Context, sid int64) []response.SubComment {
	db := mysqldao.GetDB(ctx)
	var (
		comments    []dao.Comment
		refcomments []response.SubComment
	)
	if err := db.Where("SID = ? and isDelete = 0", sid).Find(&comments).Error; err != nil {
		return nil
	}
	// 使用sort.Slice函数来排序comments切片
	sort.Slice(comments, func(i, j int) bool {
		return comments[i].UpdateTime > comments[j].UpdateTime
	})
	for idx := range comments {
		item := comments[idx]
		refcomments = append(refcomments, response.SubComment{
			Cid:        &item.CID,
			FCID:       &item.FCID,
			Text:       &item.Text,
			Uid:        &item.UID,
			UpdateTime: item.UpdateTime,
		})
	}
	//没错误，返回
	return refcomments
}

func DeleteComment(ctx *gin.Context, req *request.CommentReq) error {
	//开启事务
	db := mysqldao.GetDB(ctx)
	tx := db.Begin()
	var comment dao.Comment
	err := tx.First(&comment, req.CID).Error
	if err != nil {
		tx.Rollback()
		response.ResponseError(ctx, constanct.SOLUTION_DELETE_FAILED)
		return err
	}
	err = tx.Delete(&comment).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func AddComment(ctx *gin.Context, req *request.CommentReq) error {
	// 新建数据库事务
	db := mysqldao.GetDB(ctx)
	tx := db.Begin()
	newComment := dao.Comment{
		UID:        req.Uid,
		SID:        req.Sid,
		FCID:       req.FCID,
		Text:       req.Text,
		CreateTime: utils.GetNow(),
		UpdateTime: utils.GetNow(),
	}
	//新增一条记录
	err := tx.Create(&newComment).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//提交事务
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
func CommentOperator(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.CommentReq)
	err := ctx.ShouldBindWith(req, binding.JSON)
	if err != nil {
		// 请求参数有误 直接返回响应
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	UID := req.Uid
	if UID != middlewares.GetUid(ctx) {
		logger.Errorf("Failed to get user information, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
	}
	//fmt.Printf("req:%+v\n", req)
	if req.ActionType == constanct.ADDCODE {
		//判断内容是否为空
		if req.Text == "" {
			logger.Errorf("add comment failed, because text is null")
			response.ResponseError(ctx, constanct.InvalidParamCode)
			return
		}
		err := AddComment(ctx, req)
		if err != nil {
			logger.Errorf("add comment failed")
			response.ResponseError(ctx, constanct.InvalidParamCode)
			return
		}
		//响应
		response.ResponseOK(ctx, response.CreateResponse(constanct.SuccessCode))

	} else if req.ActionType == constanct.DELETECODE {
		// 检查id不为空
		if req.CID == 0 {
			logger.Errorf("user '%s' delete solution failed, because solutionIDStr is null.", req)
			response.ResponseError(ctx, constanct.InvalidParamCode)
			return
		}
		// 执行删除题解操作
		err = DeleteComment(ctx, req)
		if err != nil {
			logger.Errorf("user '%s' delete comment failed.beceuse %v", req, err)
			response.ResponseError(ctx, constanct.InvalidParamCode)
			return
		}
		//成功
		response.ResponseOK(ctx, response.CreateResponse(constanct.SuccessCode))
	} else {
		logger.Errorf("Unknown request parameters")
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
}
