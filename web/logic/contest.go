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

	"github.com/gin-gonic/gin"
)

func AddContest(ctx *gin.Context, req *request.AddContestReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	contest := dao.Contest{
		UID:         middlewares.GetUid(ctx),
		Title:       req.Title,
		Description: req.Description,
		Begin_time:  req.Begin_time,
		End_time:    req.End_time,
		Ctype:       req.Ctype,
		Ispublic:    req.Ispublic,
		Problems:    req.Problems,
		Pass:        req.Pass,
	}
	err := models.AddContestToDb(ctx, contest)
	if err != nil {
		logger.Errorf("call AddContestToDb failed, err=%s", err.Error())
		return nil, err
	}

	contest.CID, err = models.GetCurrentCID(ctx, contest)
	if err != nil {
		logger.Errorf("call GetCurrentCID failed, err=%s", err.Error())
		return nil, err
	}
	err = models.AddConproblems(ctx, req.Problems, contest.CID)
	if err != nil {
		logger.Errorf("call AddConproblems failed, err=%s", err.Error())
		return nil, err
	}
	return response.CreateResponse(constanct.SuccessCode), nil
}

func EditContest(ctx *gin.Context, req *request.EditContestReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	contest := dao.Contest{
		CID:         req.CID,
		UID:         req.UID,
		Title:       req.Title,
		Description: req.Description,
		Begin_time:  req.Begin_time,
		End_time:    req.End_time,
		Ctype:       req.Ctype,
		Ispublic:    req.Ispublic,
		Problems:    req.Problems,
		Pass:        req.Pass,
	}
	err := models.SaveContestDB(ctx, contest)
	if err != nil {
		logger.Errorf("call SaveContestDB failed, err=%s", err.Error())
		return nil, err
	}
	err = models.AddConproblems(ctx, req.Problems, contest.CID)
	if err != nil {
		logger.Errorf("call AddConproblems failed, err=%s", err.Error())
		return nil, err
	}
	return response.CreateResponse(constanct.SuccessCode), nil
}

func DeleteContest(ctx *gin.Context, req *request.DeleteContestReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	err := models.DeleteContestDB(ctx, req.CID)
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
			CID:       contest.CID,
			UID:       contest.UID,
			Title:     contest.Title,
			BeginTime: contest.Begin_time,
			EndTime:   contest.End_time,
			Type:      contest.Ctype,
			Ispublic:  contest.Ispublic,
		}
	}
	counts, _ := models.GetContestCountFromDB(ctx)
	sort.Slice(respData, func(i, j int) bool {
		if respData[i].EndTime == respData[j].EndTime {
			if respData[i].BeginTime == respData[j].BeginTime {
				return respData[i].CID < respData[j].CID
			}
			return respData[i].BeginTime < respData[j].BeginTime
		}
		return respData[i].EndTime > respData[j].EndTime
	})
	return response.ContestListResp{
		Response: response.CreateResponse(constanct.SuccessCode),
		Size:     counts,
		Data:     respData,
	}, nil
}

func GetContest(ctx *gin.Context, req *request.GetContestReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	contest, err := models.GetContestFromDB(ctx, req.CID)
	if err != nil {
		logger.Errorf("call GetContestFromDB failed, CID=%s, err=%s", req.CID, err.Error())
		return nil, err
	}
	if contest.CID != req.CID {
		logger.Errorf("contest not exites req=%+v", utils.Sdump(req))
		return response.CreateResponse(constanct.CIDNotExistCode), nil
	}

	uid := middlewares.GetUid(ctx)
	isAdmin := false
	if uid != "" {
		isAdmin = models.CheckUserPermission(ctx, uid, models.Contest_creator)
	}

	//不是管理员的情况下 竞赛私有  并且 （没有密码，或者密码错误）
	if !isAdmin && contest.Ispublic != 1 && ((req.Pass == nil) || (req.Pass != nil && *req.Pass != contest.Pass)) {
		logger.Errorf("contest pass word error req=%+v", utils.Sdump(req))
		return response.CreateResponse(constanct.CIDPassWordErrorCode), nil
	}
	conPros, err := models.GetConProblemFromDB(ctx, req.CID)
	if err != nil {
		logger.Errorf("call GetConProblemFromDB failed, CID=%s, err=%s", req.CID, err.Error())
		return nil, err
	}
	respData := make([]response.ConProItem, 0)
	for _, problem := range conPros {
		temp := response.ConProItem{
			PID:        problem.PID,
			Ptitle:     problem.Ptitle,
			Submit_num: problem.Submit_num,
			Ac_num:     problem.Ac_num,
		}
		respData = append(respData, temp)
	}
	return response.GetContestResp{
		Response:    response.CreateResponse(constanct.SuccessCode),
		CID:         contest.CID,
		UID:         contest.UID,
		Title:       contest.Title,
		Description: contest.Description,
		Begin_time:  contest.Begin_time,
		End_time:    contest.End_time,
		Ctype:       contest.Ctype,
		Ispublic:    contest.Ispublic,
		Size:        int64(len(conPros)),
		Problems:    contest.Problems,
		ProblemData: respData,
	}, nil
}

func initRankItem(rank *response.RankItem, Uname, Userid string, problemSize int) {
	rank.Uname = Uname
	rank.UserID = Userid
	rank.AllSubmit = 0
	rank.Problems = make([]response.ProblemItem, problemSize)
	for idx := range rank.Problems {
		problem := &rank.Problems[idx]
		problem.PID = 0
		problem.Time = 0
		problem.Status = constanct.OJ_JUDGE
	}
}

//这个待后期优化
/*rank UID,uname,solve 罚时 A，B，C，D，E，F，G...*/
func GteRankContest(ctx *gin.Context, req *request.GetContestRankReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	contest, err := models.GetContestFromDB(ctx, req.CID)
	if err != nil {
		logger.Errorf("call GetContestFromDB Failed, CID=%d, err=%s", req.CID, err.Error())
		return nil, err
	}
	problems, err := models.GetConProblemFromDB(ctx, req.CID) //获得竞赛的题目
	if err != nil {
		logger.Errorf("call GetConProblemFromDB Failed, CID=%d, err=%s", req.CID, err.Error())
		return nil, err
	}

	problemIdxMap := make(map[int]int, 0)
	for idx, problem := range problems {
		problemIdxMap[problem.PID] = idx
	}

	submits, err := models.GetSubmitByCIDFromDB(ctx, int(req.CID), req.Page, req.Limit) //获取使用这个竞赛的所有提交
	sort.Slice(submits, func(i, j int) bool {
		return submits[i].SubmitTime < submits[j].SubmitTime
	})
	if err != nil {
		logger.Errorf("call GetContestFromDB Failed, CID=%d, err=%s", req.CID, err.Error())
		return nil, err
	}
	userMap := make(map[string]int, 0)
	ranks := make(response.RankItems, 0)
	idx := 0
	for _, submit := range submits {
		rid, ok := userMap[submit.UID]
		if !ok {
			rid = idx
			idx += 1
			userMap[submit.UID] = rid
			user := dao.User{UID: submit.UID}
			models.FindUserByUID(ctx, &user)
			ranks = append(ranks, response.RankItem{})
			initRankItem(&ranks[rid], user.Uname, submit.UID, len(problems))
		}
		rank := &ranks[rid]
		problem := &rank.Problems[problemIdxMap[submit.PID]]
		problem.PID = submit.PID
		if problem.Status == constanct.OJ_AC {
			continue
		} else {
			problem.Status = submit.Result
			problem.Time = submit.SubmitTime - contest.Begin_time
			rank.AllSubmit++
			if submit.Result == constanct.OJ_AC {
				rank.ACNumber++
			}
		}
	}
	return response.ConntestRankResp{
		Response: response.CreateResponse(constanct.SuccessCode),
		Size:     ranks.Len(),
		Data:     ranks,
	}, nil
}
