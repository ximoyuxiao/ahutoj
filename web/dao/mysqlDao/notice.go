package mysqldao

import (
	"ahutoj/web/dao"
	"context"
)

func CreateNotice(ctx context.Context, notice dao.Notice) error {
	db := GetDB(ctx)
	return db.Create(&notice).Error
}

// 删除
func DeleteNoticeByNID(ctx context.Context, NID int) error {
	db := GetDB(ctx)
	return db.Delete(&dao.Notice{}, "NID = ?", NID).Error
}

// 更新
func UpdateNotice(ctx context.Context, Notice dao.Notice) error {
	db := GetDB(ctx)
	return db.Model(&dao.Notice{}).
		Where("NID = ?", Notice.NID).
		Updates(&Notice).
		Error
}

// 查询
func GetNotice(ctx context.Context, NID int) (*dao.Notice, error) {
	db := GetDB(ctx)
	notice := &dao.Notice{}
	err := db.Where("NID = ? and IsDelete=0", NID).Find(notice).Error
	if err != nil {
		return nil, err
	}
	return notice, nil
}

// 查询所有公告
func GetAllNotices(ctx context.Context) ([]dao.Notice, error) {
	db := GetDB(ctx)
	notices := make([]dao.Notice, 0)
	err := db.Order("NID desc").Where("IsDelete=0").Find(&notices).Error
	if err != nil {
		return nil, err
	}
	return notices, nil
}
