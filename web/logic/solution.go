package logic

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/middlewares"
	"ahutoj/web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
)

func AddSoulution(ctx *gin.Context, req *request.SolutionReq) (*response.SolutionPublishResp, error) {
	// 新建数据库事务
	db := mysqldao.GetDB(ctx)
	tx := db.Begin()
	newComment := dao.Solution{
		UID:        req.Uid,
		PID:        req.Pid,
		Title:      req.Title,
		Text:       req.Text,
		CreateTime: utils.GetNow(),
		UpdateTime: utils.GetNow(),
	}
	var resp response.SolutionPublishResp
	//新增一条记录
	err := tx.Create(&newComment).Error
	if err != nil {
		tx.Rollback()
		return &resp, err
	}
	//提交事务
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return &resp, err
	}
	resp.SID = newComment.SID
	return &resp, nil
}
func EditSolution(ctx *gin.Context, req *request.SolutionReq) error {
	db := mysqldao.GetDB(ctx)
	tx := db.Begin()
	var newSoution dao.Solution
	err := tx.First(&newSoution, req.Sid).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//这里加入原文文本
	ctx.JSON(http.StatusOK, response.SolutionEditResp{Title: newSoution.Title, Text: newSoution.Text})
	newSoution.Title = req.Title
	newSoution.UpdateTime = utils.GetNow()
	newSoution.Text = req.Text
	err = tx.Save(&newSoution).Error
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

func DeleteSolution(ctx *gin.Context, req *request.SolutionReq) error {
	//开启事务
	db := mysqldao.GetDB(ctx)
	tx := db.Begin()
	var solution dao.Solution
	err := tx.First(&solution, req.Sid).Error
	if err != nil {
		tx.Rollback()
		response.ResponseError(ctx, constanct.SOLUTION_DELETE_FAILED)
		return err
	}
	err = tx.Delete(&solution).Error
	if err != nil {
		tx.Rollback()
		response.ResponseError(ctx, constanct.SOLUTION_DELETE_FAILED)
		return err
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		response.ResponseError(ctx, constanct.SOLUTION_DELETE_FAILED)
		return err
	}
	return nil
}

func GetSolutiontList(ctx *gin.Context, req *request.GetSolutionListReq) (*response.SoultionsResp, error) {
	db := mysqldao.GetDB(ctx)
	var solutions []dao.Solution
	var refsolutions response.SoultionsResp
	refsolutions.SolutionList = make([]response.SolutionResponseElement, 0)
	offset, limit := utils.GetPageInfo(req.Page, req.Limit)
	if err := db.Where("PID = ? and isDelete = ?", req.PID, 0).Offset(offset).Limit(limit).Find(&solutions).Error; err != nil {
		return &refsolutions, nil
	}
	var count int64
	if err := db.Model(dao.Solution{}).Where("PID = ? and isDelete = ?", req.PID, 0).Count(&count).Error; err != nil {
		return &refsolutions, err
	}
	UID := middlewares.GetUid(ctx)
	// 先按 FavoriteCount 降序排序，如果相同再按 UpdateTime 降序序排序

	for idx := range solutions {
		item := solutions[idx]
		count := GetFavoriteCount(ctx, item.SID)
		refsolutions.SolutionList = append(refsolutions.SolutionList, response.SolutionResponseElement{
			Data:          GetSubCommentList(ctx, item.SID),
			Sid:           &item.SID,
			Text:          &item.Text,
			Title:         &item.Title,
			Uid:           &item.UID,
			CreateTime:    item.CreateTime,
			UpdateTime:    item.UpdateTime,
			FavoriteCount: &count,
			IsFavorite:    MyFavorite(ctx, int(item.SID), UID),
		})
	}
	sort.Slice(refsolutions.SolutionList, func(i, j int) bool {
		if refsolutions.SolutionList[i].FavoriteCount == refsolutions.SolutionList[j].FavoriteCount {
			return solutions[i].UpdateTime > solutions[j].UpdateTime
		}
		return *refsolutions.SolutionList[i].FavoriteCount > *refsolutions.SolutionList[j].FavoriteCount
	})
	refsolutions.Count = int(count)
	//没错误，返回
	return &refsolutions, nil
}
