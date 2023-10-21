package rediscache

import (
	"context"
)

func SetLastSource(ctx context.Context, UID string, PID string, SourceMD5 string) bool {
	rdfd := GetRedis()
	if rdfd == -1 {
		return false
	}
	defer CloseRDB(rdfd)
	/*source-PID-UID*/
	key := "source-" + PID + "-" + UID
	err := SetKey(ctx, rdfd, key, SourceMD5)
	if err != nil && err.Error() == Nil {
		return false
	}
	return true
}

func GetLastSource(ctx context.Context, UID string, PID string) string {
	rdfd := GetRedis()
	if rdfd == -1 {
		return ""
	}
	defer CloseRDB(rdfd)
	var ret string
	/*source-PID-UID*/
	key := "source-" + PID + "-" + UID
	err := GetKey(ctx, rdfd, key, &ret)
	if err != nil && err.Error() == Nil {
		return ""
	}
	return ret
}
