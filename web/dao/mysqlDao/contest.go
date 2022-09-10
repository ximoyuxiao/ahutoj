package mysqldao

import (
	"ahutoj/web/dao"
	"context"
)

func SelectContestByCid(ctx context.Context, cid int64) (dao.Contest, error) {
	db := GetDB(ctx)
	ret := dao.Contest{}
	err := db.Table(ret.TableName()).Where("cid=?", cid).Find(&ret).Error
	return ret, err
}

func SelectContests(ctx context.Context, offset, limit int) ([]dao.Contest, error) {
	db := GetDB(ctx)
	tp := dao.Contest{}
	ret := make([]dao.Contest, 0)
	err := db.Table(tp.TableName()).Offset(offset).Limit(limit).Find(&ret).Error
	return ret, err
}

func DeleteContest(ctx context.Context, cid int64) error {
	db := GetDB(ctx)
	tp := dao.Contest{}
	err := db.Table(tp.TableName()).Where("cid=?", cid).Delete(&cid).Error
	return err
}

func InserContest(ctx context.Context, contest dao.Contest) error {
	db := GetDB(ctx)
	err := db.Table(contest.TableName()).Create(&contest).Error
	return err
}

func UpdateContest(ctx context.Context, contest dao.Contest) error {
	db := GetDB(ctx)
	err := db.Table(contest.TableName()).Updates(&contest).Error
	return err
}

func SelectContestCount(ctx context.Context) (int64, error) {
	db := GetDB(ctx)
	contest := dao.Contest{}
	count := int64(0)
	err := db.Table(contest.TableName()).Count(&count).Error
	return count, err
}
