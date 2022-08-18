package mysqldao

import (
	"ahutoj/web/dao"
	"context"
)

func InsertConProblem(ctx context.Context, conPro dao.ConPro) error {
	db := GetDB(ctx)
	err := db.Table(conPro.TableName()).Create(&conPro).Error
	return err
}

func SelectConProblemByCid(ctx context.Context, cid int64) ([]dao.ConPro, error) {
	db := GetDB(ctx)
	ret := make([]dao.ConPro, 0)
	err := db.Table(dao.ConPro{}.TableName()).Create(&ret).Error
	return ret, err
}
