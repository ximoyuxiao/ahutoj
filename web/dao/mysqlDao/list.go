package mysqldao

import (
	"ahutoj/web/dao"
	"context"
)

func SelectTraningByLid(ctx context.Context, lid int64) (dao.List, error) {
	db := GetDB(ctx)
	ret := dao.List{}
	err := db.Table("List").Where("lid=?", lid).Find(&ret).Error
	return ret, err
}

func SelectTraningList(ctx context.Context, offset, limit int64) ([]dao.List, error) {
	db := GetDB(ctx)
	ret := make([]dao.List, 0)
	err := db.Table("List").Offset(int(offset)).Limit(int(limit)).Find(&ret).Error
	return ret, err
}

func InsertTraning(ctx context.Context, traning dao.List) error {
	db := GetDB(ctx)
	return db.Table("List").Create(&traning).Error
}

func DeleteTraning(ctx context.Context, lid int64) error {
	db := GetDB(ctx)
	ret := dao.List{
		Lid: lid,
	}
	return db.Table("List").Delete(&ret).Error
}

func UpdateTraning(ctx context.Context, traning dao.List) error {
	db := GetDB(ctx)
	return db.Table("List").Where("lid=?", traning.Lid).Updates(traning).Error
}
func SaveTraning(ctx context.Context, traning dao.List) error {
	db := GetDB(ctx)
	return db.Table("List").Where("lid=?", traning.Lid).Save(traning).Error
}
func SelectListCountByLid(ctx context.Context, lid int64) (count int64, err error) {
	db := GetDB(ctx)
	err = db.Table("List").Where("lid=?", lid).Count(&count).Error
	return count, err
}
func InsertListProblem(ctx context.Context, training dao.ListProblem) error {
	db := GetDB(ctx)
	return db.Table("ListProblem").Create(&training).Error
}
