package redisdao

import (
	"ahutoj/web/dao"
	"context"
	"errors"
	"strconv"
)

func GetProblemFromDB(ctx context.Context, pid int64) (*dao.Problem, error) {
	rdfd := GetRedis()
	if rdfd == -1 {
		return nil, errors.New("insufficient Redis connection resources")
	}
	defer CloseRDB(rdfd)
	ret := new(dao.Problem)
	key := "problem-" + strconv.FormatInt(pid, 10)
	err := GetKey(ctx, rdfd, key, ret)
	if err != nil && err.Error() == Nil {
		return nil, nil
	}
	return ret, err
}

func SaveProblemToRDB(ctx context.Context, problem dao.Problem) error {
	rdfd := GetRedis()
	if rdfd == -1 {
		return errors.New("insufficient Redis connection resources")
	}
	defer CloseRDB(rdfd)
	key := "problem-" + strconv.FormatInt(int64(problem.PID), 10)
	err := SetKey(ctx, rdfd, key, problem)
	return err
}
