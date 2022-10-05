package redisdao

import (
	"ahutoj/web/dao"
	"context"
	"errors"
	"strconv"
)

func GetContestFromDB(ctx context.Context, CID int64) (*dao.Contest, error) {
	rdfd := GetRedis()
	if rdfd == -1 {
		return nil, errors.New("insufficient Redis connection resources")
	}
	defer CloseRDB(rdfd)
	ret := new(dao.Contest)
	key := "contest-" + strconv.FormatInt(CID, 10)
	err := GetKey(ctx, rdfd, key, ret)
	if err != nil && err.Error() == Nil {
		return nil, nil
	}

	return ret, err
}

func SaveContestToRDB(ctx context.Context, contest dao.Contest) error {
	rdfd := GetRedis()
	if rdfd == -1 {
		return errors.New("insufficient Redis connection resources")
	}
	defer CloseRDB(rdfd)
	err := SetKey(ctx, rdfd, "contest-"+strconv.FormatInt(contest.CID, 10), contest)
	return err
}
