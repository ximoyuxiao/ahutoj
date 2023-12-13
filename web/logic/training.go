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
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func AddTraining(req *request.ListAll, c *gin.Context) (interface{}, error) {
	list := dao.List{
		UID:         middlewares.GetUid(c),
		Title:       req.Title,
		StartTime:   time.Now().UnixMilli(),
		Description: req.Description,
		Problems:    req.Problems,
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
	PIDs := strings.Split(req.Problems, ",")
	//添加提单题目
	for _, PID := range PIDs {
		listproblem := dao.ListProblem{
			LID: list.LID,
			PID: PID,
		}
		err = models.CreateListProblem(c, &listproblem)
		if err != nil {
			//日志报错
			utils.GetLogInstance().Errorf("call CreateListProblem failed,err=%s", err.Error())
			return response.CreateResponse(constanct.TRAIN_ADD_FAILED), err
		}
	}

	return response.AddTrainResp{
		Response: response.CreateResponse(constanct.SuccessCode),
		LID:      list.LID,
	}, nil
}

func RegisterTraining(ctx *gin.Context, req *request.RegisterTrainingReq) (interface{}, error) {
	listUser := dao.ListUser{
		LID: req.LID,
		UID: req.UID,
	}
	if listUser.UID != middlewares.GetUid(ctx) {
		return response.CreateResponse(constanct.TRAIN_ADD_USER_USER_FAILED), nil
	}
	list, err := models.GetTraining(ctx, listUser.LID)
	if err != nil {
		return nil, err
	}
	if list.LID != listUser.LID {
		return response.CreateResponse(constanct.TRAIN_ADD_USER_LID_NOT_EXITIES), nil
	}
	err = models.SaveTraningUser(ctx, listUser)
	if err != nil {
		return nil, err
	}
	if list.FromLID != 0 {
		listUser.LID = list.FromLID
		list, _ := models.GetTraining(ctx, listUser.LID)
		if list.LID == listUser.LID {
			models.SaveTraningUser(ctx, listUser)
		}
	}
	return response.CreateResponse(constanct.SuccessCode), nil
}

func GetTrainUserInfo(ctx *gin.Context, req *request.ListUserReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	training, err := models.GetTraining(ctx, req.LID)
	if err != nil {
		logger.Errorf("call GetTrainingFromDB failed, LID=%d, err=%s", req.LID, err.Error())
		return response.CreateResponse(constanct.TRAIN_GET_FAILED), err
	}
	if training.LID != req.LID {
		logger.Errorf("Training not exist req=%+v", utils.Sdump(req))
		return response.CreateResponse(constanct.TRAIN_GET_LIDNotExistCode), nil
	}
	listuser := dao.ListUser{
		LID: req.LID,
		UID: middlewares.GetUid(ctx),
	}
	tinfos, err := models.FindTrainUserInfo(ctx, listuser)
	if err != nil {
		return nil, err
	}
	if len(tinfos) <= 0 {
		return response.CreateResponse(constanct.TRAIN_GET_USER_NOT_FOUND_CODE), nil
	}
	tinfo := tinfos[0]
	if tinfo.LID != listuser.LID || tinfo.UID != listuser.UID {
		return response.CreateResponse(constanct.TRAIN_GET_USER_NOT_FOUND_CODE), nil
	}
	PIDs := strings.Split(training.Problems, ",")
	SolvedPID := make([]string, 0)
	for _, PID := range PIDs {
		if PID == "" {
			continue
		}
		solved, err := models.GetSubmitListCount(ctx, dao.Submit{
			PID:    PID,
			Result: constanct.OJ_AC,
			UID:    listuser.UID,
		})
		if err != nil {
			return nil, err
		}
		if solved > 0 {
			SolvedPID = append(SolvedPID, PID)
			fmt.Println(SolvedPID)
			fmt.Println(PID)
			tinfo.Solved++
		}
	}
	return response.TrainingUserResp{
		Response:  response.CreateResponse(constanct.SuccessCode),
		Solved:    tinfo.Solved,
		Submited:  tinfo.Submited,
		SolvedPID: SolvedPID,
	}, nil
}

func EditTraining(req *request.EditListReq, c *gin.Context) (interface{}, error) {
	list := dao.List{
		LID:         req.LID,
		Title:       req.Title,
		Description: req.Description,
		Problems:    req.Problems,
	}
	//编辑题单
	err := models.EditList(c, &list)
	if err != nil {
		//日志报错
		utils.GetLogInstance().Errorf("call EditList failed,err=%s", err.Error())
		return response.CreateResponse(constanct.TRAIN_EDIT_FAILED), err
	}
	//编辑提单题目
	err2 := models.EditListProblem(c, req.LID, req.Problems)
	if err2 != nil {
		//日志报错
		utils.GetLogInstance().Errorf("call EditListProblem failed,err=%s", err2.Error())
		return response.CreateResponse(constanct.TRAIN_EDIT_FAILED), err2
	}
	return response.CreateResponse(constanct.SuccessCode), nil
}

func DeleteTraining(req *request.DelListReq, c *gin.Context) (interface{}, error) {
	for _, LID := range req.LID {
		//删除题单
		err := models.DeleteTraining(c, &dao.List{
			LID: LID,
		})
		if err != nil {
			//日志报错
			utils.GetLogInstance().Errorf("call DeleteList failed,err=%s", err.Error())
			return response.CreateResponse(constanct.TRAIN_DELETE_FAILED), err
		}
	}
	return response.CreateResponse(constanct.SuccessCode), nil
}

func GetTrainingList(ctx *gin.Context, req *request.TrainingListReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	var size = 20
	if req.Limit > 20 {
		size = req.Limit
	}
	var offset = 0
	if req.Page > 0 {
		offset = size * req.Page
	}
	TrainingList, err := models.GetTrainingList(ctx, offset, size)
	if err != nil {
		logger.Errorf("call GetTrainingListFromDb failed,err=%s", err.Error())
		return nil, err
	}
	fmt.Println(TrainingList)
	respData := make([]response.TrainingListItem, 0, len(TrainingList))
	fmt.Println(respData)
	for _, training := range TrainingList {
		respData = append(respData, response.TrainingListItem{
			LID:       training.LID,
			UID:       training.UID,
			Title:     training.Title,
			StartTime: training.StartTime,
		})
	}
	list := dao.List{}
	count, err := models.SelectListCountByList(ctx, list)
	if err != nil {
		logger.Errorf("call SelectListCountByList failed,err=%s", err.Error())
		return nil, err
	}
	return response.TrainingListResp{
		Response: response.CreateResponse(constanct.SuccessCode),
		Size:     count,
		Data:     respData,
	}, err
}

func GetTraining(ctx *gin.Context, req *request.TrainingReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	training, err := models.GetTraining(ctx, req.LID)
	if err != nil {
		logger.Errorf("call GetTrainingFromDB failed, LID=%d, err=%s", req.LID, err.Error())
		return response.CreateResponse(constanct.TRAIN_GET_FAILED), err
	}
	if training.LID != req.LID {
		logger.Errorf("Training not exist req=%+v", utils.Sdump(req))
		return response.CreateResponse(constanct.TRAIN_GET_LIDNotExistCode), nil
	}

	TrainPros, err := models.GetTrainingProblem(ctx, req.LID)
	if err != nil {
		logger.Errorf("call GetTraProblemFromDB failed, LID=%d, err=%s", req.LID, err.Error())
		return response.CreateResponse(constanct.TRAIN_GET_FAILED), err
	}
	PIDs := strings.Split(training.Problems, ",")
	PIDmap := make(map[string]int, 0)
	for i, v := range PIDs {
		PIDmap[v] = i
	}
	respData := make([]response.TrainingProblemItem, 0)
	for _, listproblem := range TrainPros {
		value, ok := PIDmap[listproblem.PID]
		if !ok {
			continue
		}
		temp := response.TrainingProblemItem{
			Sort:   value,
			PID:    listproblem.PID,
			Ptitle: listproblem.Title,
		}
		respData = append(respData, temp)
	}
	return response.GetTrainResp{
		Response:    response.CreateResponse(constanct.SuccessCode),
		LID:         training.LID,
		UID:         training.UID,
		Description: training.Description,
		Title:       training.Title,
		StartTime:   training.StartTime,
		ProblemData: respData,
		Problems:    training.Problems,
	}, nil
}
func initListRankItem(rank *response.TraininngRankItem, user dao.User, problemSize int) {
	rank.Uname = user.Uname
	rank.UID = user.UID
	rank.Uclass = user.Classes
	rank.Solved = 0
	rank.Problems = make(response.TrainingRankProblemItems, problemSize)
	for idx := range rank.Problems {
		problem := &rank.Problems[idx]
		problem.PID = ""
		problem.Time = 0
		problem.Status = constanct.OJ_JUDGE
	}
}
func GetRankTraining(ctx *gin.Context, req *request.GetTrainingRankReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	// 获取题单信息
	training, err := models.GetTraining(ctx, req.LID)
	if err != nil {
		logger.Errorf("call GetTraining Failed, LID=%d, err=%s", req.LID, err.Error())
		return response.CreateResponse(constanct.TRAIN_RANK_FAILED), err
	}
	problems := strings.Split(training.Problems, ",")
	problemIdxMap := make(map[string]int, 0)
	for idx, problem := range problems {
		problemIdxMap[problem] = idx
	}

	//获取题单用户
	listUsers, err := models.FindTrainUserInfo(ctx, dao.ListUser{LID: req.LID})
	if err != nil {
		logger.Errorf("call FindTrainUserInfo Failed, LID=%d, err=%s", req.LID, err.Error())
		return nil, err
	}
	// 获取信息
	submits, err := models.FindListSubmitInfo(ctx, listUsers)
	if err != nil {
		logger.Errorf("call FindListSubmitInfo Failed, LID=%d, err=%s", req.LID, err.Error())
		return response.CreateResponse(constanct.TRAIN_RANK_FAILED), err
	}
	userMap := make(map[string]int, 0)
	ranks := make(response.TraininngRankItems, 0)
	idx := 0
	for _, submit := range submits {
		fmt.Printf("SID:%v PID:%v\n", submit.SID, submit.PID)
		rid, ok := userMap[submit.UID]
		if !ok {
			rid = idx
			idx += 1
			userMap[submit.UID] = rid
			user := dao.User{UID: submit.UID}
			models.FindUserByUID(ctx, &user)
			ranks = append(ranks, response.TraininngRankItem{})
			initListRankItem(&ranks[rid], user, len(problems))
		}
		rank := &ranks[rid]
		problem := &rank.Problems[problemIdxMap[submit.PID]]
		problem.PID = submit.PID
		fmt.Println(utils.Sdump(problem))
		if problem.Status == constanct.OJ_AC {
			continue
		} else {
			problem.Status = submit.Result
			problem.Time = uint64(submit.SubmitTime - training.StartTime)
			problem.SubmitNumber++
			if submit.Result == constanct.OJ_AC {
				rank.Solved++
			}
		}
	}
	return response.TrainingRankResp{
		Response: response.CreateResponse(constanct.SuccessCode),
		Size:     int64(len(ranks)),
		Data:     ranks,
	}, nil
}

func CloneTrainUser(ctx *gin.Context, req *request.CloneTraniningReq) (interface{}, error) {
	uid := middlewares.GetUid(ctx)
	if uid != req.UID {
		return response.CreateResponse(constanct.ADMIN_ADD_UIDEmpty), nil
	}
	ok := models.IsListExistByLID(ctx, &dao.List{
		LID: req.LID,
	})
	if !ok {
		return response.CreateResponse(constanct.TRAIN_GET_LIDNotExistCode), nil
	}
	list, _ := models.GetTraining(ctx, req.LID)
	list.FromLID = req.LID
	list.LID = 0
	list.UID = req.UID
	list.StartTime = utils.GetNow()
	// 添加题单
	err := models.CreateList(ctx, list)
	if err != nil {
		//日志报错
		utils.GetLogInstance().Errorf("call CreateList failed,err=%s", err.Error())
		return response.CreateResponse(constanct.TRAIN_ADD_FAILED), err
	}
	list.LID, err = models.GetCurrentLID(ctx, *list)
	if err != nil {
		//日志报错
		utils.GetLogInstance().Errorf("call GetLID in CreateList failed,err=%s", err.Error())
		return response.CreateResponse(constanct.TRAIN_ADD_FAILED), err
	}
	problems, _ := models.GetTrainingProblem(ctx, req.LID)
	// 添加提单题目
	for _, PID := range problems {
		listproblem := dao.ListProblem{
			LID: list.LID,
			PID: PID.PID,
		}
		err = models.CreateListProblem(ctx, &listproblem)
		if err != nil {
			//日志报错
			utils.GetLogInstance().Errorf("call CreateListProblem failed,err=%s", err.Error())
			return response.CreateResponse(constanct.TRAIN_ADD_FAILED), err
		}
	}

	return response.CloneTraniningResp{
		Response: response.CreateResponse(constanct.SuccessCode),
		LID:      list.LID,
	}, nil
}
