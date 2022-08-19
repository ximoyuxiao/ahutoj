package mysqldao

import (
	"ahutoj/web/dao"
	"context"
)

func SelectSubmitByCid(ctx context.Context, cid int64) (ans []dao.Submit, err error) {
	db := GetDB(ctx)
	err = db.Where("cid=?", cid).Find(&ans).Error
	return ans, err
}
