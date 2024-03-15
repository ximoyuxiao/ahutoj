package logic

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/utils"
	"github.com/gin-gonic/gin"
	"sort"
)

func GetCommentList(ctx *gin.Context, req *request.CommentListReq) response.CommentListResp {
	db := mysqldao.GetDB(ctx)
	logger := utils.GetLogInstance()
	var (
		comments    []dao.Comment
		refcomments response.CommentListResp
	)
	offset, limit := utils.GetPageInfo(req.Page, req.Limit)
	if err := db.Where("SID = ? and isDelete = 0", req.SID).Offset(offset).Limit(limit).Find(&comments).Error; err != nil {
		logger.Errorf("call SelectCommentBySID failed, err = %s", err.Error())
	}
	// 使用sort.Slice函数来排序comments切片
	sort.Slice(comments, func(i, j int) bool {
		return comments[i].UpdateTime > comments[j].UpdateTime
	})
	var count int64
	if err := db.Model(dao.Comment{}).Where("SID = ? and isDelete = 0", req.SID).Count(&count).Error; err != nil {
		logger.Errorf("call CountCommentBySID failed, err = %s", err.Error())
	}
	refcomments.Count = count
	for idx := range comments {
		item := comments[idx]
		refcomments.Data = append(refcomments.Data, response.SubComment{
			Cid:        &item.CID,
			FCID:       &item.FCID,
			Text:       &item.Text,
			Uid:        &item.UID,
			UpdateTime: item.UpdateTime,
		})
	}
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
