package models

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	redisdao "ahutoj/web/dao/redisDao"
	"ahutoj/web/utils"
	"context"
	"time"
)

func AddContestToDb(ctx context.Context, contest dao.Contest) error {
	return mysqldao.InserContest(ctx, contest)
}
func GetCurrentCID(ctx context.Context, contest dao.Contest) (int64, error) {
	return mysqldao.SelectContestByUID(ctx, contest.UID)
}
func SaveContestDB(ctx context.Context, contest dao.Contest) error {
	logger := utils.GetLogInstance()
	err := mysqldao.UpdateContest(ctx, contest)
	if err != nil {
		logger.Errorf("call mysqldao.UpdateContest faile,contest=%v ,err=%s", utils.Sdump(contest), err.Error())
		return err
	}
	contestTmp, err := redisdao.GetContestFromDB(ctx, contest.CID)
	if err != nil {
		logger.Errorf("call redisdao.GetContestFromDB faile,contest=%v ,err=%s", utils.Sdump(contest), err.Error())
		return err
	}
	/*更新数据库内容*/
	if contestTmp != nil {
		return redisdao.SaveContestToRDB(ctx, contest)
	}
	return nil
}

func DeleteContestDB(ctx context.Context, CID int64) error {
	// logger := utils.GetLogInstance()
	return mysqldao.DeleteContest(ctx, CID)
}

func GetContestListFromDb(ctx context.Context, offset, pagesize int) ([]dao.Contest, error) {
	return mysqldao.SelectContests(ctx, offset, pagesize)
}

func GetContestFromDB(ctx context.Context, CID int64) (dao.Contest, error) {
	/*比赛，进行中一般不会修改*/
	logger := utils.GetLogInstance()
	contest, err := redisdao.GetContestFromDB(ctx, CID)
	if err != nil || contest == nil {
		contest, err = mysqldao.SelectContestByCID(ctx, CID)
		if err != nil {
			logger.Errorf("call SelectContestByCID failed,CID=%v, err=%v", CID, err.Error())
			return dao.Contest{}, err
		}
		now := time.Now().Unix()
		if contest.End_time >= now && contest.Begin_time <= now {
			err = redisdao.SaveContestToRDB(ctx, *contest)
		}
	}
	return *contest, err
}

func GetContestCountFromDB(ctx context.Context) (int64, error) {
	return mysqldao.SelectContestCount(ctx)
}
