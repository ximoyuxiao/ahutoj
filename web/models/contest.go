package models

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"context"
)

func AddContestToDb(ctx context.Context, contest dao.Contest) error {
	return mysqldao.InserContest(ctx, contest)
}
func GetCurrentCID(ctx context.Context, contest dao.Contest) (int64, error) {
	return mysqldao.SelectContestByUID(ctx, contest.UID)
}
func SaveContestDB(ctx context.Context, contest dao.Contest) error {
	return mysqldao.UpdateContest(ctx, contest)
}

func DeleteContestDB(ctx context.Context, CID int64) error {
	return mysqldao.DeleteContest(ctx, CID)
}

func GetContestListFromDb(ctx context.Context, offset, pagesize int) ([]dao.Contest, error) {
	return mysqldao.SelectContests(ctx, offset, pagesize)
}

func GetContestFromDB(ctx context.Context, CID int64) (dao.Contest, error) {
	return mysqldao.SelectContestByCID(ctx, CID)
}

func GetContestCountFromDB(ctx context.Context) (int64, error) {
	return mysqldao.SelectContestCount(ctx)
}
