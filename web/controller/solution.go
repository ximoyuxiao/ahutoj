package controller

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func GetUserInfo(ctx *gin.Context, id string) (*dao.User, error) {
	// 根据id 检测用户是否存在数据库中
	db := mysqldao.GetDB(ctx)
	var userInter *dao.User
	err := db.Where("UID = ?", id).First(&userInter).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userInter, errors.New("the user is not exist")
		}
		return userInter, err
	}
	return userInter, err
}

func AddSoulution(ctx *gin.Context, req *request.SolutionReq) (*response.SolutionPublish, error) {
	// 新建数据库事务
	db := mysqldao.GetDB(ctx)
	tx := db.Begin()
	newComment := dao.Solution{
		UID:   req.Uid,
		PID:   req.Pid,
		Title: req.Title,
		Text:  req.Text,
	}
	var resp response.SolutionPublish
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
	return &resp, nil
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

func GetSolutiontList(ctx *gin.Context, req *request.SolutionListReq) (response.SolutionList, error) {
	db := mysqldao.GetDB(ctx)
	var solutions []dao.Solution
	var refsolutions response.SolutionList
	if err := db.Where("PID = ?", req.PID).Find(&solutions).Error; err != nil {
		return refsolutions, err
	}
	for _, solution := range solutions {
		refsolutions.SolutionList = append(refsolutions.SolutionList, solution)
	}
	//没错误，返回
	return refsolutions, nil
}
func SolutionOperator(ctx *gin.Context) {
	logger := utils.GetLogInstance()
	req := new(request.SolutionReq)
	err := ctx.ShouldBindWith(req, binding.JSON)
	user, err := GetUserInfo(ctx, req.Uid)
	if err != nil {
		// 请求参数有误 直接返回响应
		logger.Errorf("call ShouldBindWith failed, err = %s", err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	//fmt.Printf("req:%+v\n", req)
	if req.ActionType == 1 {
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

	} else if req.ActionType == 2 {
		// 检查id不为空
		if req.Sid == "" {
			logger.Errorf("user '%s' delete solution failed, because solutionIDStr is null.", user.Uname)
			response.ResponseError(ctx, constanct.InvalidParamCode)
			return
		}
		// 执行删除题解操作
		err = DeleteSolution(ctx, req)
		if err != nil {
			logger.Errorf("user '%s' delete solution failed.beceuse %v", user.Uname, err)
			response.ResponseError(ctx, constanct.InvalidParamCode)
			return
		}
		//成功
		response.ResponseOK(ctx, response.CreateResponse(constanct.SuccessCode))
		return
	} else {
		logger.Errorf("Unknown request parameters")
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
}
