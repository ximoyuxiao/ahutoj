package logic

import (
	"ahutoj/web/dao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/middlewares"
	"ahutoj/web/models"
	"ahutoj/web/utils"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

func AddContest(ctx *gin.Context, req *request.AddContestReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	contest := dao.Contest{
		Cid:         middlewares.GenSnowflakeID(),
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
	respData := make([]response.ContestListItem, len(ContestList))
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
	counts, _ := models.GetContestCountFromDB(ctx)
	return response.ContestListResp{
		Response: response.CreateResponse(constanct.SuccessCode),
		Size:     counts,
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
func initRankItem(rank *response.RankItem, Uname, Userid string) {
	rank.Mark = 0
	rank.Penalty = 0
	rank.Rank = 1
	rank.Uname = Uname
	rank.UserID = Userid
}

//这个待后期优化
/*rank uid,uname,solve 罚时 A，B，C，D，E，F，G...*/
func GteRankContest(ctx *gin.Context, req *request.GetContestRankReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	contest, err := models.GetContestFromDB(ctx, req.Cid)
	if err != nil {
		logger.Errorf("call GetContestFromDB Failed, cid=%d, err=%s", req.Cid, err.Error())
		return nil, err
	}
	problems, err := models.GetConProblemFromDB(ctx, req.Cid) //获得竞赛的题目
	if err != nil {
		logger.Errorf("call GetConProblemFromDB Failed, cid=%d, err=%s", req.Cid, err.Error())
		return nil, err
	}
	problemMap := make(map[int]dao.ConPro, 0)
	problemIdxMap := make(map[int]int, 0)
	for idx, problem := range problems {
		temp := problem
		problemMap[problem.Pid] = temp
		problemIdxMap[problem.Pid] = idx
	}
	submits, err := models.GetSubmitByCidFromDB(ctx, int(req.Cid), req.Page, req.Limit) //获取使用这个竞赛的所有提交
	if err != nil {
		logger.Errorf("call GetContestFromDB Failed, cid=%d, err=%s", req.Cid, err.Error())
		return nil, err
	}
	userMap := make(map[string]int, 0)
	ranks := make(response.RankItems, 0)
	idx := 0
	for _, submit := range submits {
		rid, ok := userMap[submit.Uid]
		if !ok {
			rid = idx
			idx += 1
			userMap[submit.Uid] = rid
			user := dao.User{Uid: submit.Uid}
			models.FindUserByUid(ctx, &user)
			ranks = append(ranks, response.RankItem{})
			initRankItem(&ranks[rid], user.Uname, submit.Uid)
		}
		rank := &ranks[rid]
		if submit.Result == "Accept" {
			if rank.Problems[problemIdxMap[submit.Pid]].Status == 2 {
				continue
			} else {
				rank.Problems[problemIdxMap[submit.Pid]].Status = 2
				rank.Problems[problemIdxMap[submit.Pid]].Pid = submit.Pid
				rank.Problems[problemIdxMap[submit.Pid]].Time = submit.SubmitTime - contest.Begin_time
				rank.Penalty += submit.SubmitTime - contest.Begin_time
				rank.Solve += 1
			}
		} else {
			if rank.Problems[problemIdxMap[submit.Pid]].Status == 0 {
				rank.Problems[problemIdxMap[submit.Pid]].Status = 1
				rank.Problems[problemIdxMap[submit.Pid]].Pid = submit.Pid
				rank.Problems[problemIdxMap[submit.Pid]].Time = submit.SubmitTime - contest.Begin_time
			}
			if submit.Result != "Compile Error" {
				// 15分钟罚时
				rank.Penalty += 1000 * 15 * 60
			}
		}
		rank.Problems[problemIdxMap[submit.Pid]].Submit_num += 1
	}
	sort.Sort(ranks)
	for idx := range ranks {
		ranks[idx].Rank = int64(idx + 1)
	}
	return response.ConntestRankResp{
		Response: response.CreateResponse(constanct.SuccessCode),
		Size:     ranks.Len(),
		Data:     ranks,
	}, nil
}
