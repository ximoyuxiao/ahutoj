package mysqldao

import (
	"ahutoj/web/dao"
	"context"
)

func SelectTraningByLID(ctx context.Context, LID int64) (dao.List, error) {
	db := GetDB(ctx)
	ret := dao.List{}
	err := db.Table("List").Where("LID=?", LID).Find(&ret).Error
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

func DeleteTraning(ctx context.Context, LID int64) error {
	db := GetDB(ctx)
	ret := dao.List{
		LID: LID,
	}
	return db.Table("List").Delete(&ret).Error
}

func UpdateTraning(ctx context.Context, traning dao.List) error {
	db := GetDB(ctx)
	return db.Table("List").Where("LID=?", traning.LID).Updates(traning).Error
}
func SaveTraning(ctx context.Context, traning dao.List) error {
	db := GetDB(ctx)
	return db.Table("List").Where("LID=?", traning.LID).Save(traning).Error
}
func SelectListCountByLID(ctx context.Context, LID int64) (count int64, err error) {
	db := GetDB(ctx)
	err = db.Table("List").Where("LID=?", LID).Count(&count).Error
	return count, err
}
func InsertListProblem(ctx context.Context, training dao.ListProblem) error {
	db := GetDB(ctx)
	return db.Table("ListProblem").Create(&training).Error
}
func UpdateListProblem(ctx context.Context, training dao.ListProblem) error {
	db := GetDB(ctx)
	return db.Table("ListProblem").Where("LID=?", training.LID).Updates(training).Error
}
func GetTrainingList(ctx context.Context, offset, limit int) ([]dao.List, error) {
	db := GetDB(ctx)
	tp := dao.List{}
	ret := make([]dao.List, 0)
	err := db.Table(tp.Title).Offset(offset).Limit(limit).Find(&ret).Error
	return ret, err
}
