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

func GetConProblemFromDB(ctx context.Context, CID int64) ([]dao.ConPro, error) {
	return mysqldao.SelectConProblemByCID(ctx, CID)
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
			CID:        CID,
			PID:        problem.PID,
			Ptitle:     problem.Title,
			Submit_num: 0,
			Ac_num:     0,
		}
		err := AddConProblemToDb(ctx, conPro)
		if err != nil {
			logger.Errorf("call AddProblemToDb failed, err=%s", err.Error())
			return err
		}
	}
	return nil
}
