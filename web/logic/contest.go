package logic

import (
	"ahutoj/web/dao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/models"
	"ahutoj/web/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AddContest(ctx *gin.Context, req *request.AddContestReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	contest := dao.Contest{
		Cid:         utils.GenSnowflakeID(),
		Uid:         req.Uid,
		Title:       req.Title,
		Description: req.Description,
		Begin_time:  req.Begin_time,
		End_time:    req.End_time,
		Ctype:       req.Ctype,
		Ispublic:    req.Ispublic,
		Pass:        req.Pass,
	}
	err := models.AddContestToDb(ctx, contest)
	if err != nil {
		logger.Errorf("call AddContestToDb failed, err=%s", err.Error())
		return nil, err
	}
	pids := strings.Split(req.Pids, ",")
	problems, err := models.GetProblems(ctx, pids)
	if err != nil {
		logger.Errorf("call GetProblems failed,err=%s", err.Error())
		return nil, err
	}
	for _, problem := range problems {
		conPro := dao.ConPro{
			Cid:        contest.Cid,
			Pid:        problem.Pid,
			Ptitle:     problem.Title,
			Submit_num: 0,
			Ac_num:     0,
		}
		err := models.AddConProblemToDb(ctx, conPro)
		if err != nil {
			logger.Errorf("call AddProblemToDb failed, err=%s", err.Error())
			return nil, err
		}
	}
	return response.CreateResponse(constanct.SuccessCode), nil
}

func EditContest(ctx *gin.Context, req *request.EditContestReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	contest := dao.Contest{
		Cid:         req.Cid,
		Uid:         req.Uid,
		Title:       req.Title,
		Description: req.Description,
		Begin_time:  req.Begin_time,
		End_time:    req.End_time,
		Ctype:       req.Ctype,
		Ispublic:    req.Ispublic,
		Pass:        req.Pass,
	}
	err := models.SaveContestDB(ctx, contest)

	if err != nil {
		logger.Errorf("call SaveContestDB failed, err=%s", err.Error())
		return nil, err
	}
	return response.CreateResponse(constanct.SuccessCode), nil
}

func DeleteContest(ctx *gin.Context, req *request.DeleteContestReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	err := models.DeleteContestDB(ctx, req.Cid)
	if err != nil {
		logger.Errorf("call DeleteContestDB failed, err=%s", err.Error())
		return nil, err
	}
	return response.CreateResponse(constanct.SuccessCode), nil
}

func GetListContest(ctx *gin.Context, req *request.ContestListReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	var size int = 20
	if req.Limit > 20 {
		size = req.Limit
	}
	var offset int = 0
	if req.Page > 0 {
		offset = size * req.Page
	}
	ContestList, err := models.GetContestListFromDb(ctx, offset, size)
	if err != nil {
		logger.Errorf("call GetContestListFromDb failed,err=%s", err.Error())
		return nil, err
	}
	respData := make([]response.ContestListItem, 0, len(ContestList))
	for i, contest := range ContestList {
		respData[i] = response.ContestListItem{
			Cid:        contest.Cid,
			Uid:        contest.Uid,
			Title:      contest.Title,
			Begin_time: contest.Begin_time,
			End_time:   contest.End_time,
			Ctype:      contest.Ctype,
			Ispublic:   contest.Ispublic,
		}
	}
	return response.ContestListResp{
		Response: response.CreateResponse(constanct.SuccessCode),
		Size:     int64(len(ContestList)),
		Data:     respData,
	}, nil
}

func GetContest(ctx *gin.Context, req *request.GetContestReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	contest, err := models.GetContestFromDB(ctx, req.Cid)
	if err != nil {
		logger.Errorf("call GetContestFromDB failed, cid=%s, err=%s", req.Cid, err.Error())
		return nil, err
	}
	conPros, err := models.GetConProblemFromDB(ctx, req.Cid)
	if err != nil {
		logger.Errorf("call GetConProblemFromDB failed, cid=%s, err=%s", req.Cid, err.Error())
		return nil, err
	}
	respData := make([]response.ConProItem, 0)
	for _, problem := range conPros {
		temp := response.ConProItem{
			Pid:        problem.Pid,
			Ptitle:     problem.Ptitle,
			Submit_num: problem.Submit_num,
			Ac_num:     problem.Ac_num,
		}
		respData = append(respData, temp)
	}
	return response.GetContestResp{
		Response:    response.CreateResponse(constanct.SuccessCode),
		Cid:         contest.Cid,
		Uid:         contest.Uid,
		Title:       contest.Title,
		Description: contest.Description,
		Begin_time:  contest.Begin_time,
		End_time:    contest.End_time,
		Ctype:       contest.Ctype,
		Ispublic:    contest.Ispublic,
		Pass:        contest.Pass,
		Size:        int64(len(conPros)),
		ProblemData: respData,
	}, nil
}

//这个待后期优化
func GteRankContest(ctx *gin.Context, req *request.GetContestReq) (interface{}, error) {
	return nil, nil
}
