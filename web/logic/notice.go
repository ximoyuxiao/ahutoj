package logic

import (
	"ahutoj/web/dao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/middlewares"
	"ahutoj/web/models"
	"ahutoj/web/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetNotice(ctx *gin.Context, req *request.GetNoticeReq) (interface{}, error) {
	notice, err := models.GetNotice(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	fmt.Println(notice)
	if notice.NID != req.ID {
		return response.CreateResponseStr(constanct.AUTH_LOGIN_UIDNotExistCode, "这个公告被删除或者不存在", response.ERROR), nil
	}
	return response.GetNoticeResp{
		Response: response.CreateResponse(constanct.SuccessCode),
		GetNOticeItem: response.GetNOticeItem{
			UID:         notice.UID,
			ID:          notice.NID,
			Title:       notice.Title,
			Content:     notice.Content,
			CreatedTime: notice.CreateTime,
			UpdatedTime: notice.UpdateTime,
		},
	}, nil
}

func GetNoticeList(ctx *gin.Context) (interface{}, error) {
	notices, err := models.GetAllNotices(ctx)
	if err != nil {
		return nil, err
	}
	data := make([]response.GetNOticeItem, 0)
	for _, notice := range notices {
		data = append(data, response.GetNOticeItem{
			UID:         notice.UID,
			ID:          notice.NID,
			Title:       notice.Title,
			Content:     notice.Content,
			CreatedTime: notice.CreateTime,
			UpdatedTime: notice.UpdateTime,
		})
	}
	return response.GetListNoticeResp{
		Response: response.CreateResponse(constanct.SuccessCode),
		Count:    len(notices),
		Data:     data,
	}, nil
}

func CreateNotice(ctx *gin.Context, req *request.CreateNoticeReq) (interface{}, error) {
	notice := dao.Notice{
		UID:        middlewares.GetUid(ctx),
		Title:      req.Title,
		Content:    req.Content,
		CreateTime: utils.GetNow(),
		UpdateTime: 0,
		IsDelete:   false,
	}
	err := models.CreateNotice(ctx, notice)
	if err != nil {
		return nil, err
	}
	return response.CreateResponse(constanct.SuccessCode), nil
}

func UpdateNotice(ctx *gin.Context, req *request.UpdateNoticeReq) (interface{}, error) {
	notice := dao.Notice{
		NID:        req.ID,
		Title:      req.Title,
		Content:    req.Content,
		UpdateTime: utils.GetNow(),
	}
	err := models.UpdateNotice(ctx, notice)
	if err != nil {
		return nil, err
	}
	return response.CreateResponse(constanct.SuccessCode), nil
}

func DeleteNotice(ctx *gin.Context, req *request.DeleteNoticeReq) (interface{}, error) {
	notice := dao.Notice{
		NID:      req.ID,
		IsDelete: true,
	}
	err := models.UpdateNotice(ctx, notice)
	if err != nil {
		return nil, err
	}
	return response.CreateResponse(constanct.SuccessCode), nil
}
