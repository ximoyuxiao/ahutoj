package redisdao

import (
	"ahutoj/web/dao"
	"context"
	"strconv"
)

func GetContestFromDB(ctx context.Context, CID int64) (*dao.Contest, error) {
	rdfd := GetRedis()
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
	defer CloseRDB(rdfd)
	err := SetKey(ctx, rdfd, "contest-"+strconv.FormatInt(contest.CID, 10), contest)
	return err
}
