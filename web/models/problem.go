package models

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/utils"
	"context"
	"strconv"
)

//判断题目是否存在
func IsProblemExistByPID(ctx context.Context, problem *dao.Problem) bool {
	count, err := mysqldao.SelectProblemCountByPID(ctx, problem.PID)
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

func DeleteProblem(ctx context.Context, PID int64) error {
	logger := utils.GetLogInstance()
	err := mysqldao.DeleteProblem(ctx, PID)
	if err != nil {
		logger.Errorf("call DeleteProblem failed,problem= %d, err=%s", PID, err.Error())
	}
	return err
}

// 前期的话 先用 mysql 后期针对活跃数据会引入redis
func GetProblemByPID(ctx context.Context, PID int64) (dao.Problem, error) {
	logger := utils.GetLogInstance()
	problem := dao.Problem{}
	problem.PID = PID
	err := mysqldao.SelectProblemByPID(ctx, &problem)
	if err != nil {
		logger.Errorf("call SelectProblemByPID failed,PID=%d,err=%s", PID, err.Error())
		return problem, err
	}
	return problem, err
}

func GetProblemCount(ctx context.Context, problem dao.Problem) (int64, error) {
	return mysqldao.SelectProblemCount(ctx, problem)
}

func GetProblems(ctx context.Context, PIDs []string) ([]dao.Problem, error) {
	problems := make([]dao.Problem, len(PIDs))
	logger := utils.GetLogInstance()
	for idx, PIDstr := range PIDs {
		PID, err := strconv.ParseInt(PIDstr, 10, 64)
		if err != nil {
			logger.Errorf("call ParseInt failed,err=%s", err.Error())
			return nil, err
		}
		problem, err := GetProblemByPID(ctx, PID)
		if err != nil {
			logger.Errorf("call GetProblemByPID failed,err=%s", err.Error())
			return nil, err
		}
		problems[idx] = problem
	}
	return problems, nil
}
