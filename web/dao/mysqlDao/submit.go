package mysqldao

import (
	"ahutoj/web/dao"
	"ahutoj/web/io/constanct"
	"context"
)

func SelectSubmitList(ctx context.Context, submit dao.Submit, offset, limit int) (ans []dao.Submit, err error) {
	db := GetDB(ctx)
	err = db.Table(submit.TableName()).Where(&submit).Order("SID desc").Limit(limit).Offset(offset).Find(&ans).Error
	return ans, err
}
func SelectSubmitByCID(ctx context.Context, CID int64, CheckTime int64) (ans []dao.Submit, err error) {
	db := GetDB(ctx)
	sql := db.Table(dao.Submit{}.TableName()).Where("CID =?", CID)
	if CheckTime != 0 {
		sql = sql.Where("SubmitTime <= ?", CheckTime)
	}
	err = sql.Find(&ans).Error
	return ans, err
}
func SelectSubmitBySID(ctx context.Context, SID int64) (ans dao.Submit, err error) {
	db := GetDB(ctx)
	err = db.Table(dao.Submit{}.TableName()).Where("SID=?", SID).Find(&ans).Error
	return ans, err
}

func InsertSubmit(ctx context.Context, submit dao.Submit) (err error) {
	db := GetDB(ctx)
	err = db.Table(submit.TableName()).Create(&submit).Error
	return err
}

func RejudgeSubmit(ctx context.Context, submit dao.Submit) (err error) {
	db := GetDB(ctx)
	temp := dao.Submit{
		Result: constanct.OJ_REJUDGE,
	}
	err = db.Table(temp.TableName()).Where(&submit).Updates(&temp).Error
	return err
}
func SelectCountSubmit(ctx context.Context, submit dao.Submit) (int64, error) {
	db := GetDB(ctx)
	var ans int64
	err = db.Table(submit.TableName()).Where(&submit).Count(&ans).Error
	return ans, err
}
func SelectSubmitByUID(ctx context.Context, submit dao.Submit, lastTime int64) ([]dao.Submit, error) {
	ans := make([]dao.Submit, 0)
	db := GetDB(ctx)
	err := db.Table(submit.TableName()).Where(submit).Where("SubmitTime > ?", lastTime).Find(&ans).Error
	return ans, err
}
