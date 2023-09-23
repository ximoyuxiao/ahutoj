package rediscache

import (
	"ahutoj/web/dao"
	"context"
	"errors"
)

var ProblemMAXValue string

func GetProblemFromDB(ctx context.Context, PID string) (*dao.Problem, error) {
	rdfd := GetRedis()
	if rdfd == -1 {
		return nil, errors.New("insufficient Redis connection resources")
	}
	defer CloseRDB(rdfd)
	ret := new(dao.Problem)
	key := "problem-" + PID
	err := GetKey(ctx, rdfd, key, ret)
	if err != nil && err.Error() == Nil {
		return nil, nil
	}
	return ret, err
}
func DelProblem(ctx context.Context, PID string) (*dao.Problem, error) {
	rdfd := GetRedis()
	if rdfd == -1 {
		return nil, errors.New("insufficient Redis connection resources")
	}
	defer CloseRDB(rdfd)
	ret := new(dao.Problem)
	key := PID
	err := DelKey(ctx, rdfd, key)
	if err != nil {
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
	key := "problem-" + problem.PID
	err := SetKey(ctx, rdfd, key, problem)
	return err
}

func GetLastANDPID(ctx context.Context) (int64, error) {
	rdfd := GetRedis()
	if rdfd == -1 {
		return 0, errors.New("insufficient Redis connection resources")
	}
	var PID int64
	defer CloseRDB(rdfd)
	err := GetKey(ctx, rdfd, "NewPID", &PID)
	return PID, err
}

func UpdateNextPID(ctx context.Context, PID string) error {
	rdfd := GetRedis()
	if rdfd == -1 {
		return errors.New("insufficient Redis connection resources")
	}
	defer CloseRDB(rdfd)
	err := SetKey(ctx, rdfd, "NewPID", PID)
	return err
}
