package mysqldao

import (
	"ahutoj/web/dao"
	"context"
)

func SelectCeinfoBySID(ctx context.Context, SID int64) (dao.CeInfo, error) {
	db := GetDB(ctx)
	ret := dao.CeInfo{}
	err := db.Table(ret.TableName()).Where("SID=?", SID).Find(&ret).Error
	return ret, err
}

func InsertCeInfo(ctx context.Context, ceInfo dao.CeInfo) error {
	db := GetDB(ctx)
	err := db.Table(ceInfo.TableName()).Create(&ceInfo).Error
	return err
}
