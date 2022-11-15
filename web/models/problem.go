package models

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	redisdao "ahutoj/web/dao/redisDao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/utils"
	"context"
	"fmt"
)

// 判断题目是否存在
func IsProblemExistByPID(ctx context.Context, problem *dao.Problem) bool {
	count, err := mysqldao.SelectProblemCountByPID(ctx, problem.PID)
	if err != nil {
		return false
	}
	return count > 0
}

// 创建题目
func CreateProblem(ctx context.Context, problem *dao.Problem) error {
	logger := utils.GetLogInstance()
	err := mysqldao.InsertProblemTable(ctx, *problem)
	if err != nil {
		logger.Errorf("call InsertProblemTable failed,problem= %+v, err=%s", utils.Sdump(problem), err.Error())
	}
	return err
}

// 编辑题目
func EditProblem(ctx context.Context, problem *dao.Problem) error {
	logger := utils.GetLogInstance()
	err := mysqldao.EditProblemTable(ctx, *problem)
	if err != nil {
		logger.Errorf("call EditProblemTable failed,problem= %+v, err=%s", utils.Sdump(problem), err.Error())
	}
	return err
}

func DeleteProblem(ctx context.Context, PID string) error {
	logger := utils.GetLogInstance()
	err := mysqldao.DeleteProblem(ctx, PID)
	if err != nil {
		logger.Errorf("call DeleteProblem failed,problem= %d, err=%s", PID, err.Error())
	}
	return err
}

// 前期的话 先用 mysql 后期针对活跃数据会引入redis
func GetProblemByPID(ctx context.Context, PID string) (dao.Problem, error) {
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
	for idx, PID := range PIDs {
		problem, err := GetProblemByPID(ctx, PID)
		if err != nil {
			logger.Errorf("call GetProblemByPID failed,err=%s", err.Error())
			return nil, err
		}
		problems[idx] = problem
	}
	return problems, nil
}
func ChekckProblemType(ctx context.Context, PType constanct.ProblemType) bool {
	if PType == "" {
		return true
	}
	if PType == constanct.LOCALTYPE {
		return true
	}
	if PType == constanct.ATCODERTYPE {
		return true
	}
	if PType == constanct.CODEFORCESTYPE {
		return true
	}
	return false
}
func GetNextProblemPID(ctx context.Context) (string, error) {
	logger := utils.GetLogInstance()
	PID, err := redisdao.GetLastANDPID(ctx)
	if err != nil || PID == 0 {
		PID, err = mysqldao.SelectProblemLastPID(ctx)
		logger.Debugf("PID:%v", PID)

		if err != nil {
			logger.Errorf("call SelectProblemLastPID failed,err:%v", err.Error())
			return "", err
		}
	}
	return fmt.Sprintf("%v", PID+1), nil
}
