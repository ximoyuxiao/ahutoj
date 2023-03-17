package rediscache

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
	// 函数 结束的时候   调用 defer后面的函数
	defer CloseRDB(rdfd)
	ret := new(dao.Contest) /*指针*/
	/*hash 表 通过 key -  value 去 获得*/
	key := "contest-" + strconv.FormatInt(CID, 10)
	err := GetKey(ctx, rdfd, key, ret)
	/*这句话 就是判断 没有这个key的情况下*/
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
