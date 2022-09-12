package logic

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/models"
	"ahutoj/web/utils"

	"github.com/gin-gonic/gin"
)

func AddSubmit(ctx *gin.Context, req *request.AddSubmitReq) (interface{}, error) {
	submit := dao.Submit{
		Pid:        req.Pid,
		Cid:        req.Cid,
		Uid:        req.Uid,
		Source:     req.Source,
		Lang:       req.Lang,
		SubmitTime: req.SubmitTime,
	}
	err := models.CreateSubmit(ctx, submit)
	if err != nil {
		return nil, err
	}
	return response.CreateResponse(constanct.SuccessCode), nil
}

func RejudgeSubmit(ctx *gin.Context, req *request.RejudgeSubmitReq) (interface{}, error) {
	submit := dao.Submit{}
	if req.Sid != nil {
		submit.Sid = *req.Sid
	}
	if req.Cid != nil {
		submit.Cid = *req.Cid
	}
	if req.Pid != nil {
		submit.Pid = *req.Pid
	}
	if req.Uid != nil {
		submit.Uid = *req.Uid
	}
	err := models.RejudgeSubmit(ctx, submit)
	if err != nil {
		return nil, err
	}
	return response.CreateResponse(constanct.SuccessCode), nil
}

func GetSubmits(ctx *gin.Context, req *request.SubmitListReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	resp := response.SubmitListResp{}
	submit := dao.Submit{}
	if req.Cid != nil {
		submit.Cid = *req.Cid
	}
	if req.Pid != nil {
		submit.Pid = *req.Pid
	}
	if req.Uid != nil {
		submit.Uid = *req.Uid
	}
	offset, limit := utils.GetPageInfo(req.Page, req.Limit)
	submits, err := models.GetSubmitList(ctx, submit, offset, limit)
	if err != nil {
		logger.Errorf("call SelectSubmitList failed,req=%+v,err=%s", utils.Sdump(req), err.Error())
		return nil, err
	}
	resp.Count, err = models.GetSubmitListCount(ctx, submit)
	if err != nil {
		logger.Errorf("call GetSubmitListCount failed,req=%+v,err=%s", utils.Sdump(req), err.Error())
		return nil, err
	}
	resp.Response = response.CreateResponse(constanct.SuccessCode)
	resp.Data = make([]response.SubmitLIstItem, len(submits))
	for i, temp := range submits {
		resp.Data[i] = response.SubmitLIstItem{
			Sid:        temp.Sid,
			Pid:        temp.Pid,
			Lang:       temp.Lang,
			Result:     temp.Result,
			UseTime:    temp.Usetime,
			UseMemory:  temp.Memory,
			SubmitTime: temp.SubmitTime,
		}
	}
	return response.CreateResponse(constanct.Notimplemented), nil
}

func GetSubmit(ctx *gin.Context, req *request.GetSubmitReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	submit, err := mysqldao.SelectSubmitBySid(ctx, req.Sid)
	if err != nil {
		logger.Errorf("Call SelectSubmitBySid failed, sid=%v, err=%s", req.Sid, err.Error())
		return response.CreateResponse(constanct.MySQLErrorCode), err
	}
	return response.GetSubmitResp{
		Response:   response.CreateResponse(constanct.Notimplemented),
		Sid:        submit.Sid,
		Pid:        submit.Pid,
		Source:     submit.Source,
		Lang:       submit.Lang,
		Result:     submit.Result,
		UseTime:    submit.Usetime,
		UseMemory:  submit.Memory,
		SubmitTime: submit.SubmitTime,
	}, nil
}
