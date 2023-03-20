package mysqldao

import (
	"ahutoj/web/dao"
	"context"

	"gorm.io/gorm"
)

func InsertConProblem(ctx context.Context, conPro dao.ConPro) error {
	db := GetDB(ctx)
	err := db.Table(conPro.TableName()).Create(&conPro).Error
	return err
}

func SelectConProblemByCID(ctx context.Context, CID int64) ([]dao.ConPro, error) {
	db := GetDB(ctx)
	ret := make([]dao.ConPro, 0)
	err := db.Table(dao.ConPro{}.TableName()).Where("CID=?", CID).Find(&ret).Error
	return ret, err
}

func SelectCountConProInContestByProblem(ctx context.Context, CID int64, PID string) int64 {
	db := GetDB(ctx)
	var ret int64
	db.Table(dao.ConPro{}.TableName()).Where("CID=? and PID=?", CID, PID).Count(&ret)
	return ret
}

func IncConProSubmit(ctx context.Context, CID int64, PID string) error {
	db := GetDB(ctx)
	conpor := dao.ConPro{}
	return db.Table(conpor.TableName()).Where("CID=? and PID=?", CID, PID).UpdateColumn("Submited", gorm.Expr("Submited+1")).Error
}

func IncConProSolved(ctx context.Context, CID int64, PID string) error {
	db := GetDB(ctx)
	conpor := dao.ConPro{}
	return db.Table(conpor.TableName()).Where("CID=? and PID=?", CID, PID).UpdateColumn("Solved", gorm.Expr("Solved+1")).Error
}
