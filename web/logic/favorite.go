package logic

import (
	rediscache "ahutoj/web/cache/redis"
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/io/request"
	"fmt"
	"github.com/gin-gonic/gin"
	"errors"
	"gorm.io/gorm"
)

func MyFavorite(ctx *gin.Context, SID int, UID string) bool {
	if UID == "" {
		return false
	}
	db := mysqldao.GetDB(ctx)
	var count int64
	db.Model(dao.Favorite{}).Where("SID = ? and UID = ?", SID, UID).Count(&count)
	return count > 0
}
func Favorite(ctx *gin.Context, req *request.FavoriteReq) error {
	db := mysqldao.GetDB(ctx)
	var (
		favorite dao.Favorite
	)
	if FavoriteGetByRedis(ctx, req.SID) != 0 {
		FavoriteAddToRedis(ctx, req.SID)
	}
	result := db.Where("SID = ? and UID = ?", req.SID, req.UID).First(&favorite)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// 查询成功，没有找到数据
		err := db.Create(&dao.Favorite{SID: req.SID, UID: req.UID})
		if err != nil {
			return err.Error
		}
	} else {
		return result.Error
	}
	return nil
}

func UnFavorite(ctx *gin.Context, req *request.FavoriteReq) error {
	db := mysqldao.GetDB(ctx)
	var (
		favorite dao.Favorite
	)
	if FavoriteGetByRedis(ctx, req.SID) != 0 {
		FavoriteSubToRedis(ctx, req.SID)
	}
	result := db.Where("SID = ? and UID = ?", req.SID, req.UID).Find(&favorite)
	if result.Error == nil {
		// 删除
		err := db.Where(favorite).Delete(favorite).Error
		if err != nil {
			return err
		}
		return nil
	} else {
		return result.Error
	}

}

func GetFavoriteKey(SID int64) string {
	return fmt.Sprintf("Solution_favorite_%v", SID)
}

func FavoriteGetByRedis(ctx *gin.Context, SID int64) int64 {
	rdfd := rediscache.GetRedis()
	var count int64
	rediscache.GetKey(ctx, rdfd, GetFavoriteKey(SID), &count)
	return count
}

func FavoriteAddToRedis(ctx *gin.Context, SID int64) {
	count := FavoriteGetByRedis(ctx, SID)
	rdfd := rediscache.GetRedis()
	rediscache.SetKey(ctx, rdfd, GetFavoriteKey(SID), count+1)
}

func FavoriteSubToRedis(ctx *gin.Context, SID int64) {
	count := FavoriteGetByRedis(ctx, SID)
	rdfd := rediscache.GetRedis()
	rediscache.SetKey(ctx, rdfd, GetFavoriteKey(SID), count-1)
}

func GetFavoriteCountByDb(ctx *gin.Context, SID int64) int64 {
	db := mysqldao.GetDB(ctx)
	var count int64
	err := db.Model(dao.Favorite{}).Where("SID = ?", SID).Count(&count).Error
	if err != nil {
		fmt.Println(err.Error())
	}
	return count
}
func FavoriteToRedis(ctx *gin.Context, SID int64)int64{
	var count =GetFavoriteCountByDb(ctx,SID)
	rdfd := rediscache.GetRedis()
	rediscache.SetKey(ctx, rdfd, GetFavoriteKey(SID), count)
	return count
}
func GetFavoriteCount(ctx *gin.Context, SID int64) int64 {
	count := FavoriteGetByRedis(ctx, SID)
	if count == 0 {
		count = FavoriteToRedis(ctx, SID)
	}
	return count
}
