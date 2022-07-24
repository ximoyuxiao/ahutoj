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

func SelectProblemByLists(ctx context.Context, offset, size int64) ([]dao.Problem, error) {
	db := GetDB(ctx)
	ret := make([]dao.Problem, 0, size)
	err := db.Table("Problem").Offset(int(offset)).Limit(int(size)).Find(&ret).Error
	return ret, err
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

func DeleteProblem(ctx context.Context, pid int64) error {
	db := GetDB(ctx)
	err := db.Table("Problem").Where("Pid=?", pid).Delete(pid).Error
	return err
}
