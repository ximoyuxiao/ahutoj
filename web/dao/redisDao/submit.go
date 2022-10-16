package redisdao

import (
	"context"
	"strconv"
)

func SetLastSource(ctx context.Context, UID string, PID int64, Source string) bool {
	rdfd := GetRedis()
	if rdfd == -1 {
		return false
	}
	defer CloseRDB(rdfd)
	/*source-PID-UID*/
	key := "source-" + strconv.FormatInt(PID, 10) + "-" + UID
	err := SetKey(ctx, rdfd, key, Source)
	if err != nil && err.Error() == Nil {
		return false
	}
	return true
}

func GetLastSource(ctx context.Context, UID string, PID int64) string {
	rdfd := GetRedis()
	if rdfd == -1 {
		return ""
	}
	defer CloseRDB(rdfd)
	var ret string
	/*source-PID-UID*/
	key := "source-" + strconv.FormatInt(PID, 10) + "-" + UID
	err := GetKey(ctx, rdfd, key, &ret)
	if err != nil && err.Error() == Nil {
		return ""
	}
	return ret
}
