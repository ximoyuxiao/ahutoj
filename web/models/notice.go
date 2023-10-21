package models

import (
	rediscache "ahutoj/web/cache/redis"
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/utils"
	"context"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/gin-gonic/gin"
)

func CreateNotice(ctx context.Context, notice dao.Notice) error {
	return mysqldao.CreateNotice(ctx, notice)
}

// 删除
func DeleteNoticeByNID(ctx context.Context, NID int) error {
	return mysqldao.DeleteNoticeByNID(ctx, NID)
}
func NoticeEqualLastSource(ctx *gin.Context, title string, content string) bool {
	SourceMD5, err := utils.MD5EnCodeStr(title + content)
	if err != nil {
		logger.Errorf("call MD5EnCodeStr failed. Source:%s", title+content)
		return false
	}
	rdfd := rediscache.GetRedis()
	if rdfd == -1 {
		return false
	}
	defer rediscache.CloseRDB(rdfd)
	var ret string
	err = rediscache.GetKey(ctx, rdfd, SourceMD5, &ret)
	if err != nil && err.Error() == rediscache.Nil {
		rediscache.SetKey2(ctx, rdfd, SourceMD5, "1")
		return false
	}
	return true
}

// 更新
func UpdateNotice(ctx context.Context, notice dao.Notice) error {
	return mysqldao.UpdateNotice(ctx, notice)
}

// 查询
func GetNotice(ctx context.Context, NID int) (*dao.Notice, error) {
	return mysqldao.GetNotice(ctx, NID)
}

// 查询所有公告
func GetAllNotices(ctx context.Context) ([]dao.Notice, error) {
	return mysqldao.GetAllNotices(ctx)
}
