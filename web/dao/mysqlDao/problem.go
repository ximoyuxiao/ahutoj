package mysqldao

import (
	"ahutoj/web/dao"
	"context"
)

func SelectProblemByPID(ctx context.Context, problem *dao.Problem) error {
	db := GetDB(ctx)
	err := db.Table("Problem").Where("PID=?", problem.PID).Find(problem).Error
	return err
}

func SelectProblemCountByPID(ctx context.Context, PID int64) (count int64, err error) {
	db := GetDB(ctx)
	err = db.Table("Problem").Where("PID=?", PID).Count(&count).Error
	return count, err
}

func SelectProblemCount(ctx context.Context) (count int64, err error) {
	db := GetDB(ctx)
	err = db.Table("Problem").Count(&count).Error
	return count, err
}

func SelectProblemByLists(ctx context.Context, offset, size int) ([]dao.Problem, error) {
	db := GetDB(ctx)
	ret := make([]dao.Problem, 0, size)
	err := db.Table("Problem").Offset(offset).Limit(size).Find(&ret).Error
	return ret, err
}
func InsertProblemTable(ctx context.Context, problem dao.Problem) error {
	db := GetDB(ctx)
	err := db.Table("Problem").Create(&problem).Error
	return err
}
func EditProblemTable(ctx context.Context, problem dao.Problem) error {
	db := GetDB(ctx)
	err := db.Table("Problem").Where("PID=?", problem.PID).Updates(&problem).Error //这里不确定用法对不对
	return err
}

func DeleteProblem(ctx context.Context, PID int64) error {
	db := GetDB(ctx)
	err := db.Table("Problem").Where("PID=?", PID).Delete(PID).Error
	return err
}
