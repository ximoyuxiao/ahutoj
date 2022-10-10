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
