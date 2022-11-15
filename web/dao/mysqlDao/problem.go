package mysqldao

import (
	"ahutoj/web/dao"
	"ahutoj/web/io/constanct"
	"context"
)

func SelectProblemByPID(ctx context.Context, problem *dao.Problem) error {
	db := GetDB(ctx)
	err := db.Table("Problem").Where("PID=?", problem.PID).Find(problem).Error
	return err
}

func SelectProblemCountByPID(ctx context.Context, PID string) (count int64, err error) {
	db := GetDB(ctx)
	err = db.Table("Problem").Where("PID=?", PID).Count(&count).Error
	return count, err
}

func SelectProblemCount(ctx context.Context, problem dao.Problem) (count int64, err error) {
	db := GetDB(ctx)
	err = db.Table("Problem").Where(problem).Count(&count).Error
	return count, err
}

func SelectListProblem(ctx context.Context, offset, size int, problem dao.Problem) ([]dao.Problem, error) {
	db := GetDB(ctx)
	ret := make([]dao.Problem, 0, size)
	err := db.Table("Problem").Where(problem).Offset(offset).Limit(size).Find(&ret).Error
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

func DeleteProblem(ctx context.Context, PID string) error {
	db := GetDB(ctx)
	err := db.Table("Problem").Where("PID=?", PID).Delete(PID).Error
	return err
}

func SelectProblemLastPID(ctx context.Context) (int64, error) {
	/*
		SELECT MAX(CONVERT(SUBSTR( PID, 2 ) ,UNSIGNED)) as MAX_LOCAL_PID

		FROM
			Problem
		WHERE
			PTYPE = 'LOCAL'
	*/
	var ans int64
	db := GetDB(ctx)
	err := db.Table("Problem").Where("PType=?", constanct.LOCALTYPE).Select("MAX(CONVERT(SUBSTR( PID, 2 ) ,UNSIGNED))").Find(&ans).Error
	return ans, err
}
