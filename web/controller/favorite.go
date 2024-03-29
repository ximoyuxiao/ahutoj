package controller

import (
	rediscache "ahutoj/web/cache/redis"
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/utils"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
	result := db.Where("SID = ? and UID = ?", req.SID, req.UID).First(&favorite)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// 查询成功，但没有找到数据
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
	result := db.Where("SID = ? and UID = ?", req.SID, req.UID).Find(&favorite)
	fmt.Println(favorite)
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

func FavoriteAction(ctx *gin.Context) {
	//从上下文中获取uid username
	logger := utils.GetLogInstance()
	req := new(request.FavoriteReq)
	err := ctx.ShouldBindWith(req, binding.JSON)
	if err != nil {
		logger.Errorf("User '%s' favoriteAction err,because %s", req.UID, err.Error())
		favoriteActionRespErr(ctx, "params video_id  invalid")
		return
	}
	actionType := req.ActionType
	resp := response.FavoriteAction{}
	fmt.Println(actionType)
	// action_type 1-点赞，2-取消点赞
	//点赞相关操作
	if actionType == constanct.ADDCODE {
		err = Favorite(ctx, req)
		if err != nil {
			logger.Errorf("User '%s' favoriteAction err,because %s", req.UID, err.Error())
			favoriteActionRespErr(ctx, err.Error())
			return
		}
		//点赞成功 响应
		logger.Errorf("User %s Favorite Solution %d success. ", req.UID, req.SID)
		resp.Response = response.CreateResponse(constanct.SuccessCode)
		resp.Count = int(GetFaviroateCount(ctx, int(req.SID)))
		resp.IsFavorite = true
		response.ResponseOK(ctx, resp)
		return
	}

	//取消点赞
	if actionType == constanct.DELETECODE {
		err = UnFavorite(ctx, req)
		if err != nil {
			logger.Errorf("User '%s' UnfavoriteAction err,because %s.", req.UID, err.Error())
			//todo 更加详细的错误
			favoriteActionRespErr(ctx, err.Error())
			return
		}
		//取消点赞成功 响应
		logger.Errorf("User %s UnFavorite %d success.", req.UID, req.SID)
		resp.Response = response.CreateResponse(constanct.SuccessCode)
		resp.IsFavorite = false
		resp.Count = int(GetFaviroateCount(ctx, int(req.SID)))
		response.ResponseOK(ctx, resp)
		return
	}
	// 未知的actionType
	logger.Errorf("Illegal actionType. ")
	favoriteActionRespErr(ctx, "Illegal actionType")
	return
}
func favoriteActionRespErr(c *gin.Context, err string) {
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 1,
		StatusMsg:  err,
	})
	return
}

func FavoriteToRedis(ctx *gin.Context, SID int) {
	db := mysqldao.GetDB(ctx)
	var count int64
	err := db.Model(dao.Favorite{}).Where("SID=?", SID).Count(&count).Error
	if err != nil {

	}
	rdfd := rediscache.GetRedis()
	rediscache.SetKey(ctx, rdfd, fmt.Sprintf("Solution_favorite_%v", SID), count)
}
func GetFavoriteKey(SID int) string {
	return fmt.Sprintf("Solution_favorite_%v", SID)
}

func FavoriteGetByRedis(ctx *gin.Context, SID int) int {
	rdfd := rediscache.GetRedis()
	var count int
	rediscache.GetKey(ctx, rdfd, GetFavoriteKey(SID), &count)
	return count
}

func FavoriteAddToRedis(ctx *gin.Context, SID int) {
	count := FavoriteGetByRedis(ctx, SID)
	rdfd := rediscache.GetRedis()
	rediscache.SetKey(ctx, rdfd, GetFavoriteKey(SID), count+1)
}

func FavoriteSubToRedis(ctx *gin.Context, SID int) {
	count := FavoriteGetByRedis(ctx, SID)
	rdfd := rediscache.GetRedis()
	rediscache.SetKey(ctx, rdfd, GetFavoriteKey(SID), count-1)
}

func GetFaviroateCount(ctx *gin.Context, SID int) int64 {
	db := mysqldao.GetDB(ctx)
	var count int64
	err := db.Model(dao.Favorite{}).Where("SID = ?", SID).Count(&count).Error
	if err != nil {
		fmt.Println(err.Error())
	}
	return count
}
