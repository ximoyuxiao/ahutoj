package models

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"context"
)

func CreateNotice(ctx context.Context, notice dao.Notice) error {
	return mysqldao.CreateNotice(ctx, notice)
}

// 删除
func DeleteNoticeByNID(ctx context.Context, NID int) error {
	return mysqldao.DeleteNoticeByNID(ctx, NID)
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
