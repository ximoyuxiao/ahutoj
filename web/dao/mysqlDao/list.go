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

func SelectListByUID(ctx context.Context, UID string) (int64, error) {
	db := GetDB(ctx)
	ret := dao.List{}
	err := db.Where("UID=?", UID).Last(&ret).Error
	return ret.LID, err
}
func SelectTraningListCount(ctx context.Context) (int64, error) {
	db := GetDB(ctx)
	var ret int64
	err := db.Count(&ret).Error
	return ret, err
}
func GetTraining(ctx context.Context, LID int64) (*dao.List, error) {
	db := GetDB(ctx)
	ret := new(dao.List)
	err := db.Table(ret.TableName()).Where("LID=?", LID).Find(ret).Error
	return ret, err
}

func SelectTrainProblemByLID(ctx context.Context, LID int64) ([]dao.ListProblem, error) {
	db := GetDB(ctx)
	ret := make([]dao.ListProblem, 0)
	err := db.Select(&ret, "SELECT * FROM ListProblem WHERE LID=?", LID).Error
	return ret, err
}
