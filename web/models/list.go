package models

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/utils"

	"github.com/gin-gonic/gin"
)

func IsListExistByLid(ctx *gin.Context, list *dao.List) bool {
	count, err := mysqldao.SelectListCountByLid(ctx, list.Lid)
	if err != nil {
		return false
	}
	return count > 0
}

func CreateList(ctx *gin.Context, list *dao.List) error {
	logger := utils.GetLogInstance()
	if IsListExistByLid(ctx, list) {
		return nil
	}
	err := mysqldao.InsertTraning(ctx, *list)
	if err != nil {
		logger.Errorf("call InsertListTable failed,list= %+v, err=%s", utils.Sdump(list), err.Error())
	}
	return err
}

func CreateListProblem(ctx *gin.Context, listproblem *dao.ListProblem) error {
	logger := utils.GetLogInstance()
	err := mysqldao.InsertListProblem(ctx, *listproblem)
	if err != nil {
		logger.Errorf("call InsertListProblemTable failed,listproblem= %+v,err=%s", utils.Sdump(listproblem), err.Error())
	}
	return err

}
