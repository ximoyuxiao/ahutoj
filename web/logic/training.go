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
	"time"

	"github.com/gin-gonic/gin"
)

func AddTraining(req *request.ListAll, c *gin.Context) (interface{}, error) {
	list := dao.List{
		UID:       middlewares.GetUid(c),
		Title:     req.Title,
		StartTime: req.StartTime,
	}
	listproblem := dao.ListProblem{
		LID: req.LID,
		PID: req.PID,
	}
	//添加题单
	err := models.CreateList(c, &list)
	if err != nil {
		//日志报错
		utils.GetLogInstance().Errorf("call CreateList failed,err=%s", err.Error())
		return response.CreateResponse(constanct.TRAIN_ADD_FAILED), err
	}
	list.LID, err = models.GetCurrentLID(c, list)
	if err != nil {
		//日志报错
		utils.GetLogInstance().Errorf("call GetLID in CreateList failed,err=%s", err.Error())
		return response.CreateResponse(constanct.TRAIN_ADD_FAILED), err
	}

	//添加提单题目
	err = models.CreateListProblem(c, &listproblem)
	if err != nil {
		//日志报错
		utils.GetLogInstance().Errorf("call CreateListProblem failed,err=%s", err.Error())
		return response.CreateResponse(constanct.TRAIN_ADD_FAILED), err
	}
	return response.CreateResponse(constanct.SuccessCode), nil
}

func EditTraining(req *request.ListAll, c *gin.Context) (interface{}, error) {
	list := dao.List{
		LID:       req.LID,
		UID:       req.UID,
		Title:     req.Title,
		StartTime: req.StartTime,
	}
	listproblem := dao.ListProblem{
		LID: req.LID,
		PID: req.PID,
	}
	//编辑题单
	err := models.EditList(c, &list)
	if err != nil {
		//日志报错
		utils.GetLogInstance().Errorf("call EditList failed,err=%s", err.Error())
		return response.CreateResponse(constanct.TRAIN_EDIT_FAILED), err
	}
	//编辑提单题目
	err2 := models.EditListProblem(c, &listproblem)
	if err2 != nil {
		//日志报错
		utils.GetLogInstance().Errorf("call EditListProblem failed,err=%s", err2.Error())
		return response.CreateResponse(constanct.TRAIN_EDIT_FAILED), err2
	}
	return response.CreateResponse(constanct.SuccessCode), nil
}

func DeleteTraining(req *request.List, c *gin.Context) (interface{}, error) {
	list := dao.List{
		LID:       req.LID,
		UID:       req.UID,
		Title:     req.Title,
		StartTime: req.StartTime,
	}
	//删除题单
	err := models.DeleteTraining(c, &list)
	if err != nil {
		//日志报错
		utils.GetLogInstance().Errorf("call DeleteList failed,err=%s", err.Error())
		return response.CreateResponse(constanct.TRAIN_DELETE_FAILED), err
	}
	return response.CreateResponse(constanct.SuccessCode), nil
}

func GetTrainingList(ctx *gin.Context, req *request.TrainingListReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	var size int = 20
	if req.Limit > 20 {
		size = req.Limit
	}
	var offset int = 0
	if req.Page > 0 {
		offset = size * req.Page
	}
	TrainingList, err := models.GetTrainingList(ctx, offset, size)
	if err != nil {
		logger.Errorf("call GetTrainingListFromDb failed,err=%s", err.Error())
		return nil, err
	}
	respData := make([]response.TrainingListItem, 0, len(TrainingList))
	for i, training := range TrainingList {
		respData[i] = response.TrainingListItem{
			LID:       training.LID,
			UID:       training.UID,
			Title:     training.Title,
			StartTime: training.StartTime,
		}
	}
	return response.TrainingListResp{
		Response: response.CreateResponse(constanct.SuccessCode),
		Size:     int64(len(TrainingList)),
		Data:     respData,
	}, err
}
func GetTraining(ctx *gin.Context, req *request.TrainingReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	training, err := models.GetTraining(ctx, req.LID)
	if err != nil {
		logger.Errorf("call GetTrainingFromDB failed, CID=%s, err=%s", req.LID, err.Error())
		return response.CreateResponse(constanct.TRAIN_GET_FAILED), err
	}
	if training.LID != req.LID {
		logger.Errorf("Training not exist req=%+v", utils.Sdump(req))
		return response.CreateResponse(constanct.TRAIN_GET_LIDNotExistCode), err
	}

	TrainPros, err := models.GetTrainingProblem(ctx, req.LID)
	if err != nil {
		logger.Errorf("call GetTraProblemFromDB failed, LID=%s, err=%s", req.LID, err.Error())
		return response.CreateResponse(constanct.TRAIN_GET_FAILED), err
	}
	respData := make([]response.TrainingProblemItem, 0)
	for _, listproblem := range TrainPros {
		temp := response.TrainingProblemItem{
			PID:    listproblem.PID,
			Ptitle: listproblem.Title,
		}
		respData = append(respData, temp)
	}
	return response.GetTrainResp{
		Response:    response.CreateResponse(constanct.SuccessCode),
		LID:         training.LID,
		UID:         training.UID,
		Title:       training.Title,
		StartTime:   training.StartTime,
		ProblemData: respData,
	}, nil
}
func GetRankTraining(ctx *gin.Context, req *request.GetTrainingRankReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	training, err := models.GetTraining(ctx, req.LID)
	if err != nil {
		logger.Errorf("call GetTraining Failed, LID=%d, err=%s", req.LID, err.Error())
		return response.CreateResponse(constanct.TRAIN_RANK_FAILED), err
	}
	problems, err := models.GetTrainingProblem(ctx, req.LID) //获得竞赛的题目
	if err != nil {
		logger.Errorf("call GetConProblemFromDB Failed, CID=%d, err=%s", req.LID, err.Error())
		return response.CreateResponse(constanct.TRAIN_RANK_FAILED), err
	}

	problemIdxMap := make(map[string]int, 0)
	for idx, problem := range problems {
		problemIdxMap[problem.PID] = idx
	}
	currentTime := time.Now().UnixMilli()
	fb := int64(utils.GetConfInstance().Terminal*(float64(currentTime)-float64(training.StartTime)) + float64(training.StartTime))
	if currentTime > int64(utils.GetConfInstance().OpenTime*float64(time.Hour)) {
		fb = 0
	}
	//封榜时间
	submits, err := models.GetSubmitByLIDFromDB(ctx, req.LID, fb) //获取使用这个竞赛的所有提交
	sort.Slice(submits, func(i, j int) bool {
		return submits[i].SubmitTime < submits[j].SubmitTime
	})
	if err != nil {
		logger.Errorf("call GetTraining Failed, LID=%d, err=%s", req.LID, err.Error())
		return response.CreateResponse(constanct.TRAIN_RANK_FAILED), err
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
			initRankItem(&ranks[rid], user, len(problems))
		}
		rank := &ranks[rid]
		problem := &rank.Problems[problemIdxMap[submit.PID]]
		problem.PID = submit.PID
		if problem.Status == constanct.OJ_AC {
			continue
		} else {
			problem.Status = submit.Result
			problem.Time = submit.SubmitTime - training.StartTime
			rank.AllSubmit++
			problem.SubmitNumber++
			if submit.Result == constanct.OJ_AC {
				rank.ACNumber++
			}
			if submit.Result == constanct.OJ_CE {
				rank.CENumber++
			}
		}
	}
	return response.ConntestRankResp{
		Response: response.CreateResponse(constanct.SuccessCode),
		Size:     ranks.Len(),
		Data:     ranks,
	}, nil
}
