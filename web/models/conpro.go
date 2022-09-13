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

func GetConProblemFromDB(ctx context.Context, cid int64) ([]dao.ConPro, error) {
	return mysqldao.SelectConProblemByCid(ctx, cid)
}
func CheckHasConProInContest(ctx context.Context, pid, cid int64) bool {
	return mysqldao.SelectCountConProInContestByProblem(ctx, pid, cid) > 0
}

func AddConproblems(ctx context.Context, pids string, cid int64) error {
	logger := utils.GetLogInstance()
	pidstrs := strings.Split(pids, ",")
	problems, err := GetProblems(ctx, pidstrs)
	if err != nil {
		logger.Errorf("call GetProblems failed,err=%s", err.Error())
		return err
	}
	for _, problem := range problems {
		if CheckHasConProInContest(ctx, int64(problem.Pid), cid) {
			continue
		}
		conPro := dao.ConPro{
			Cid:        cid,
			Pid:        problem.Pid,
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
