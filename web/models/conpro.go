package models

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/utils"
	"context"
	"strings"
)

func AddConProblemToDb(ctx context.Context, conPro dao.ConPro) error {
	return mysqldao.InsertConProblem(ctx, conPro)
}

func GetConProblemFromDB(ctx context.Context, contest dao.Contest) ([]dao.ConPro, error) {
	ret := make([]dao.ConPro, 0)
	conPros, err := mysqldao.SelectConProblemByCID(ctx, contest.CID)
	if err != nil {
		return ret, err
	}
	PIDs := strings.Split(contest.Problems, ",")
	PIDMap := make(map[string]int, 0)
	for i, PID := range PIDs {
		PIDMap[PID] = i
	}
	for _, conPro := range conPros {
		_, ok := PIDMap[conPro.PID]
		if !ok {
			continue
		}
		ret = append(ret, conPro)
	}
	return ret, nil
}
func CheckHasConProInContest(ctx context.Context, PID string, CID int64) bool {
	return mysqldao.SelectCountConProInContestByProblem(ctx, CID, PID) > 0
}

func AddConproblems(ctx context.Context, PIDs string, CID int64) error {
	logger := utils.GetLogInstance()
	PIDstrs := strings.Split(PIDs, ",")
	problems, err := GetProblems(ctx, PIDstrs)
	if err != nil {
		logger.Errorf("call GetProblems failed,err=%s", err.Error())
		return err
	}
	for _, problem := range problems {
		if CheckHasConProInContest(ctx, problem.PID, CID) {
			continue
		}
		conPro := dao.ConPro{
			CID:      CID,
			PID:      problem.PID,
			Ptitle:   problem.Title,
			Submited: 0,
			Solved:   0,
		}
		err := AddConProblemToDb(ctx, conPro)
		if err != nil {
			logger.Errorf("call AddProblemToDb failed, err=%s", err.Error())
			return err
		}
	}
	return nil
}
