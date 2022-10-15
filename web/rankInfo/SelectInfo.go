package rankinfo

import (
	"context"
	"time"
)

func UpdateUserInfoThread(ctx context.Context) {
	/*每个小时 去更新一次  用户的做题数据*/
	for {
		time.Sleep(time.Hour)
	}
}
