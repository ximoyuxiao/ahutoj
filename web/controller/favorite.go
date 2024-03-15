package controller

import (
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/logic"
	"ahutoj/web/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func FavoriteAction(ctx *gin.Context) {
	//从上下文中获取uid username
	logger := utils.GetLogInstance()
	req := new(request.FavoriteReq)
	err := ctx.ShouldBindWith(req, binding.JSON)
	if err != nil {
		logger.Errorf("User '%s' favoriteAction err,because %s", req.UID, err.Error())
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	actionType := req.ActionType
	resp := response.FavoriteActionResp{}
	// action_type 1-点赞，2-取消点赞
	//点赞相关操作
	if actionType == constanct.ADDCODE {
		err = logic.Favorite(ctx, req)
		if err != nil {
			logger.Errorf("User '%s' favoriteAction err,because %s", req.UID, err.Error())
			favoriteActionRespErr(ctx, err.Error())
			return
		}
		//点赞成功 响应
		resp.Response = response.CreateResponse(constanct.SuccessCode)
		resp.Count = logic.GetFavoriteCount(ctx, req.SID)
		resp.IsFavorite = true
		response.ResponseOK(ctx, resp)
		return
	}

	//取消点赞
	if actionType == constanct.DELETECODE {
		err = logic.UnFavorite(ctx, req)
		if err != nil {
			logger.Errorf("User '%s' UnfavoriteAction err,because %s.", req.UID, err.Error())
			//todo 更加详细的错误
			favoriteActionRespErr(ctx, err.Error())
			return
		}
		resp.Response = response.CreateResponse(constanct.SuccessCode)
		resp.IsFavorite = false
		resp.Count = logic.GetFavoriteCount(ctx, req.SID)
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
