package controller

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/middlewares"
	"ahutoj/web/utils"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
	resp.Response = response.CreateResponse(constanct.SuccessCode)
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
	//todo: 在这里加入原文文本
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
		return err
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func GetSolutiontList(ctx *gin.Context, req *request.GetSolutionListReq) (*response.SoultionsResp, error) {
	db := mysqldao.GetDB(ctx)
	var solutions []dao.Solution
	var refsolutions response.SoultionsResp
	offset, limit := utils.GetPageInfo(req.Page, req.Limit)
	if err := db.Where("PID = ? and isDelete = ?", req.PID, 0).Offset(offset).Limit(limit).Find(&solutions).Error; err != nil {
		return &refsolutions, err
	}
	var count int64
	if err := db.Model(solutions[0]).Where("PID = ? and isDelete = ?", req.PID, 0).Count(&count).Error; err != nil {
		return &refsolutions, err
	}
	refsolutions.Response = response.CreateResponse(constanct.SuccessCode)
	// 先按 FavoriteCount 降序排序，如果相同再按 UpdateTime 降序序排序
	sort.Slice(solutions, func(i, j int) bool {
		if solutions[i].FavoriteCount == solutions[j].FavoriteCount {
			return solutions[i].UpdateTime > solutions[j].UpdateTime
		}
		return solutions[i].FavoriteCount > solutions[j].FavoriteCount
	})

	for idx := range solutions {
		item := solutions[idx]
		refsolutions.SolutionList = append(refsolutions.SolutionList, response.SolutionResponseElement{
			Data: GetSubCommentList(ctx, item.SID),
			//todo:将一个题解的所有评论加入
			Sid:           &item.SID,
			Text:          &item.Text,
			Title:         &item.Title,
			Uid:           &item.UID,
			CreateTime:    item.CreateTime,
			UpdateTime:    item.UpdateTime,
			FavoriteCount: &item.FavoriteCount,
		})
	}
	refsolutions.Count = int(count)
	//没错误，返回
	return &refsolutions, nil
}

func SolutionOperator(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.SolutionReq)
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
			logger.Errorf("add solution failed, because text is null")
			response.ResponseError(ctx, constanct.InvalidParamCode)
			return
		}
		resp, err := AddSoulution(ctx, req)
		if err != nil {
			logger.Errorf("add solution failed")
			response.ResponseError(ctx, constanct.InvalidParamCode)
			return
		}
		//响应
		response.ResponseOK(ctx, resp)

	} else if req.ActionType == constanct.EDITCODE {
		err := EditSolution(ctx, req)
		if err != nil {
			logger.Errorf("call EditSolution failed, err = %s", err.Error())
			response.ResponseError(ctx, constanct.ServerErrorCode)
			return
		}
		response.ResponseOK(ctx, response.CreateResponse(constanct.SuccessCode))
	} else if req.ActionType == constanct.DELETECODE {
		// 检查id不为空
		if req.Sid == 0 {
			logger.Errorf("user '%s' delete solution failed, because solutionIDStr is null.", req)
			response.ResponseError(ctx, constanct.InvalidParamCode)
			return
		}
		// 执行删除题解操作
		err = DeleteSolution(ctx, req)
		if err != nil {
			logger.Errorf("user '%s' delete solution failed.beceuse %v", req, err)
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
