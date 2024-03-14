package logic

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/middlewares"
	"ahutoj/web/models"
	"ahutoj/web/utils"
	"time"
	"github.com/gin-gonic/gin"
)

func AddSubmit(ctx *gin.Context, req *request.AddSubmitReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	// 提交代码
	submit := dao.Submit{
		PID:           req.PID,
		CID:           req.CID,
		UID:           req.UID,
		Source:        req.Source,
		Lang:          req.Lang,
		Result:        constanct.OJ_PENDING,
		SubmitTime:    time.Now().UnixMilli(),
		IsOriginJudge: false,
		OriginPID:     "",
		OJPlatform:    -1,
	}
	if models.EqualLastSource(ctx, req.UID, req.PID, submit.Source) {
		return response.CreateResponseStr(constanct.SUBMIT_ADD_DUPLICATECODE, "禁止频繁重复提交代码", response.WARNING), nil
	}
	problem, err := models.GetProblemByPID(ctx, req.PID)
	if err != nil {
		logger.Errorf("call GetProblemByPID failed,pid=%v, err=%s", req.PID, err.Error())
		return response.CreateResponse(constanct.SUBMIT_ADD_FAILEDCode), err
	}
	if problem.Origin != -1 {
		submit.OJPlatform = problem.Origin
		submit.IsOriginJudge = true
		submit.OriginPID = problem.OriginPID
	}

	if req.CID != 0 && req.CID > 0 {
		contest, err := models.GetContestFromDB(ctx, req.CID)
		if err != nil {
			logger.Errorf("call GetContestFromDB failed,pid=%v, err=%s", req.PID, err.Error())
			return response.CreateResponse(constanct.SUBMIT_ADD_FAILEDCode), err
		}
		// 比赛未开始 不能提交代码
		if contest.Begin_time > time.Now().UnixMilli() {
			return response.CreateResponse(constanct.SUBMIT_ADD_CONTESTNOTSTART_CODE), nil
		}
	}
	/*提交代码*/
	err = models.CreateSubmit(ctx, submit)
	if err != nil {
		logger.Errorf("call CreateSubmit failed, submit=%v, err=%s", submit, err.Error())
		return response.CreateResponse(constanct.SUBMIT_ADD_FAILEDCode), err
	}
	submit, err = models.FindLastSIDByUID(ctx, submit.UID)
	if err != nil {
		logger.Errorf("call FindLastSIDByUID failed, UID=%v, err=%s", submit.UID, err.Error())
		return response.CreateResponse(constanct.SUBMIT_ADD_FAILEDCode), err
	}
	// 更新用户提交代码
	err = mysqldao.IncUserSubmited(ctx, submit.UID)
	if err != nil {
		logger.Errorf("call IncUserSubmited failed, submit=%v, err=%s", submit, err.Error())
		return response.CreateResponse(constanct.SUBMIT_ADD_FAILEDCode), err
	}
	err = mysqldao.IncProblemSubmited(ctx, submit.PID)
	if err != nil {
		logger.Errorf("call IncProblemSubmited failed, submit=%v, err=%s", submit, err.Error())
		return response.CreateResponse(constanct.SUBMIT_ADD_FAILEDCode), err
	}
	// 更新题目提交代码
	if req.CID > 0 {
		err = mysqldao.IncConProSubmit(ctx, req.CID, submit.PID)
		if err != nil {
			logger.Errorf("call IncConProSubmit failed, submit=%v, err=%s", submit, err.Error())
			return response.CreateResponse(constanct.SUBMIT_ADD_FAILEDCode), err
		}
	}
	rabitmq := middlewares.GetRabbitMq()
	produce := middlewares.NewProducer(rabitmq)
	if submit.IsOriginJudge {
		err := produce.SendMessage(constanct.ORIGINJUDGE, submit)
		if err != nil {
			logger.Errorf("call SendMessage(%s) failed, submit=%v, err=%s", constanct.ORIGINJUDGE, submit, err.Error())
			return response.CreateResponse(constanct.SUBMIT_ADD_FAILEDCode), err
		}
	} else {
		err := produce.SendMessage(constanct.INNERJUDGE, submit)
		if err != nil {
			logger.Errorf("call SendMessage(%s) failed, submit=%v, err=%s", constanct.INNERJUDGE, submit, err.Error())
			return response.CreateResponse(constanct.SUBMIT_ADD_FAILEDCode), err
		}
	}
	return response.AddSubmitResp{
		Response: response.CreateResponse(constanct.SuccessCode),
		SID:      submit.SID,
	}, nil
}

func RejudgeSubmit(ctx *gin.Context, req *request.RejudgeSubmitReq) (interface{}, error) {
	submit := dao.Submit{}
	if req.SID != nil {
		submit.SID = *req.SID
	}
	if req.CID != nil {
		submit.CID = *req.CID
	}
	if req.PID != nil {
		submit.PID = *req.PID
	}
	if req.UID != nil {
		submit.UID = *req.UID
	}
	err := models.RejudgeSubmit(ctx, submit)
	if err != nil {
		return response.CreateResponse(constanct.SUBMIT_REJUDG_FAILEDCode), err
	}
	submits, err := models.GetSubmitList(ctx, submit, 0, 100000)
	if err != nil {
		return nil, err
	}
	for _, submit := range submits {
		err := models.CommitRabitMQ(ctx, submit)
		if err != nil {
			return response.CreateResponse(constanct.SUBMIT_ADD_FAILEDCode), nil
		}
	}
	return response.CreateResponse(constanct.SuccessCode), err
}

func GetSubmits(ctx *gin.Context, req *request.SubmitListReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	resp := response.SubmitListResp{}
	submit := dao.Submit{}
	if req.CID != nil {
		submit.CID = *req.CID
	}
	if req.PID != nil {
		submit.PID = *req.PID
	}
	if req.UID != nil {
		submit.UID = *req.UID
	}
	if req.Lang != nil {
		submit.Lang = *req.Lang
	}
	if req.Result != nil {
		submit.Result = *req.Result
	}
	offset, limit := utils.GetPageInfo(req.Page, req.Limit)
	submits, err := models.GetSubmitList(ctx, submit, offset, limit)
	if err != nil {
		logger.Errorf("call SelectSubmitList failed,req=%+v,err=%s", utils.Sdump(req), err.Error())
		return response.CreateResponse(constanct.SUBMIT_LIST_FAILEDCode), err
	}
	resp.Count, err = models.GetSubmitListCount(ctx, submit)
	if err != nil {
		logger.Errorf("call GetSubmitListCount failed,req=%+v,err=%s", utils.Sdump(req), err.Error())
		return response.CreateResponse(constanct.SUBMIT_LIST_FAILEDCode), err
	}
	resp.Response = response.CreateResponse(constanct.SuccessCode)
	resp.Data = make([]response.SubmitLIstItem, len(submits))
	for i, temp := range submits {
		resp.Data[i] = response.SubmitLIstItem{
			SID:        temp.SID,
			UID:        temp.UID,
			PID:        temp.PID,
			Lang:       temp.Lang,
			Result:     temp.Result,
			UseTime:    temp.Usetime,
			UseMemory:  temp.UseMemory,
			SubmitTime: temp.SubmitTime,
		}
	}
	return resp, nil
}

func GetSubmit(ctx *gin.Context, req *request.GetSubmitReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	submit, err := mysqldao.SelectSubmitBySID(ctx, req.SID)
	if !middlewares.CheckUserHasPermission(ctx, constanct.SourceBorwser) &&
		submit.UID != middlewares.GetUid(ctx) {
		return response.CreateResponse(constanct.AUTH_Token_URLVerifyCode), err
	}
	if err != nil {
		logger.Errorf("Call SelectSubmitBySID failed, SID=%v, err=%s", req.SID, err.Error())
		return response.CreateResponse(constanct.SUBMIT_GET_FAILEDCode), err
	}
	var ceInfo *string = nil
	if submit.Result == constanct.OJ_CE {
		ceInfo = new(string)
		*ceInfo = models.FindSubmitCeInfo(ctx, req.SID) //
	}
	return response.GetSubmitResp{
		Response:     response.CreateResponse(constanct.SuccessCode),
		SID:          submit.SID,
		UID:          submit.UID,
		PID:          submit.PID,
		Source:       submit.Source,
		Lang:         submit.Lang,
		Result:       submit.Result, //
		PassSample:   submit.PassSample,
		SampleNumber: submit.SampleNmuber,
		UseTime:      submit.Usetime,
		UseMemory:    submit.UseMemory,
		SubmitTime:   submit.SubmitTime,
		CeInfo:       ceInfo,
	}, nil
}
