package models

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/utils"
	"context"
)

//判断题目是否存在
func IsProblemExistByPid(ctx context.Context, problem *dao.Problem) bool {
	count, err := mysqldao.SelectProblemCountByPid(ctx, problem.Pid)
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

func DeleteProblem(ctx context.Context, pid int64) error {
	logger := utils.GetLogInstance()
	err := mysqldao.DeleteProblem(ctx, pid)
	if err != nil {
		logger.Errorf("call DeleteProblem failed,problem= %d, err=%s", pid, err.Error())
	}
	return err
}

// 前期的话 先用 mysql 后期针对活跃数据会引入redis
func GetProblemByPID(ctx context.Context, pid int64) (dao.Problem, error) {
	logger := utils.GetLogInstance()
	problem := dao.Problem{}
	problem.Pid = int(pid)
	err := mysqldao.SelectProblemByPid(ctx, &problem)
	if err != nil {
		logger.Errorf("call SelectProblemByPid failed,pid=%d,err=%s", pid, err.Error())
		return problem, err
	}
	return problem, err
}

func GetProblemCount(ctx context.Context) (int64, error) {
	return mysqldao.SelectProblemCount(ctx)
}
