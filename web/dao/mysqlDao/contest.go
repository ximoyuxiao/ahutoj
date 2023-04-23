package mysqldao

import (
	"ahutoj/web/dao"
	"context"
)

func SelectContestByCID(ctx context.Context, CID int64) (*dao.Contest, error) {
	db := GetDB(ctx)
	ret := new(dao.Contest)
	err := db.Table(ret.TableName()).Where("CID=?", CID).Find(ret).Error
	return ret, err
}

func SelectContests(ctx context.Context, offset, limit int) ([]dao.Contest, error) {
	db := GetDB(ctx)
	tp := dao.Contest{}
	ret := make([]dao.Contest, 0)
	err := db.Table(tp.TableName()).Order("EndTime desc").Offset(offset).Limit(limit).Find(&ret).Error
	return ret, err
}

func DeleteContest(ctx context.Context, CID int64) error {
	db := GetDB(ctx)
	tp := dao.Contest{}
	err := db.Table(tp.TableName()).Where("CID=?", CID).Delete(&CID).Error
	return err
}

func InserContest(ctx context.Context, contest dao.Contest) error {
	db := GetDB(ctx)
	err := db.Table(contest.TableName()).Create(&contest).Error
	return err
}

func SelectContestByUID(ctx context.Context, UID string) (int64, error) {
	db := GetDB(ctx)
	ret := dao.Contest{}
	err := db.Where("UID=?", UID).Last(&ret).Error
	return ret.CID, err
}

func UpdateContest(ctx context.Context, contest dao.Contest) error {
	db := GetDB(ctx)
	err := db.Table(contest.TableName()).Where("CID=?", contest.CID).Updates(&contest).Error
	return err
}

func SelectContestCount(ctx context.Context) (int64, error) {
	db := GetDB(ctx)
	contest := dao.Contest{}
	count := int64(0)
	err := db.Table(contest.TableName()).Count(&count).Error
	return count, err
}

func SelectContestRecently(ctx context.Context, Recently int64) (contests []dao.Contest, err error) {
	db := GetDB(ctx)
	contest := dao.Contest{}
	err = db.Table(contest.TableName()).Where("EndTime>?", Recently).Find(&contests).Error
	return contests, err
}
