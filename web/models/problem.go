package models

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/utils"
	"context"
)

//判断题目是否存在
func IsProblemExistByPid(ctx context.Context, problem *dao.Problem) bool {
	count, err := mysqldao.SelectProblemCountByPid(ctx, *problem.Pid)
	if err != nil {
		return false
	}
	return count > 0
}

//创建题目
func CreateProblem(ctx context.Context, problem *dao.Problem) error {
	logger := utils.GetLogInstance()
	err := mysqldao.InsertProblemTable(ctx, *problem)
	if err != nil {
		logger.Errorf("call InsertProblemTable failed,problem= %+v, err=%s", utils.Sdump(problem), err.Error())
	}
	return err
}

//编辑题目
func EditProblem(ctx context.Context, problem *dao.Problem) error {
	logger := utils.GetLogInstance()
	err := mysqldao.EditProblemTable(ctx, *problem)
	if err != nil {
		logger.Errorf("call EditProblemTable failed,problem= %+v, err=%s", utils.Sdump(problem), err.Error())
	}
	return err
}
