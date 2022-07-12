package mysqldao

import (
	"ahutoj/web/dao"
	"context"
)

func SelectProblemByPid(ctx context.Context, problem *dao.Problem) error {
	db := GetDB(ctx)
	err := db.Table("Problem").Where("Pid=?", problem.Pid).Find(problem).Error
	return err
}

func SelectProblemCountByPid(ctx context.Context, pid int) (count int64, err error) {
	db := GetDB(ctx)
	err = db.Table("Problem").Where("Pid=?", pid).Count(&count).Error
	return count, err
}
func InsertProblemTable(ctx context.Context, problem dao.Problem) error {
	db := GetDB(ctx)
	err := db.Table("Problem").Create(&problem).Error
	return err
}
func EditProblemTable(ctx context.Context, problem dao.Problem) error {
	db := GetDB(ctx)
	err := db.Table("Problem").Where("Pid=?", problem.Pid).Updates(&problem).Error //这里不确定用法对不对
	return err
}
