package logic

import (
	rediscache "ahutoj/web/cache/redis"
	"ahutoj/web/dao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/mapping"
	"ahutoj/web/middlewares"
	"ahutoj/web/models"
	"ahutoj/web/utils"
	"context"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

type ParseProblemHandlerFunc func(ctx context.Context, fileText string) (mapping.JsonProblems, error)

var ParseProblemFuncMap = map[string]ParseProblemHandlerFunc{
	"XML":  models.ParseXmlToproblem,
	"JSON": models.ParseJsonToProblem,
}

var AddProblemLock sync.Mutex

func AddProblem(req *request.Problem, c *gin.Context) (interface{}, error) {
	logger := utils.GetLogInstance()
	var nextPID = ""
	var err error
	AddProblemLock.Lock()
	defer AddProblemLock.Unlock()
	if req.PType == "" {
		req.PType = constanct.LOCALTYPE
	}
	problem := mapping.ProblemReqToDao(*req)
	if !models.ChekckProblemType(c, req.PType) {
		return response.CreateResponse(constanct.Problem_ADD_PTYPEERR_CODE), nil
	}
	if req.PType == constanct.LOCALTYPE {
		nextPID, err = models.GetNextProblemPID(c)
		if err != nil {
			logger.Errorf("call GetNewProblemPID failed,err:%v", err.Error())
			return nil, err
		}
		problem.PID = "P" + nextPID

	}
	//题目不存在 添加题目
	err = models.CreateProblem(c, &problem)
	if err != nil {
		//日志报错
		logger.Errorf("call CreateProblem failed,err=%s", err.Error())
		return response.CreateResponse(constanct.PROBLEM_ADD_FAILED), err
	}
	if nextPID != "" {
		err = rediscache.UpdateNextPID(c, nextPID)
		if err != nil {
			logger.Errorf("call UpdateNextPID failed,err:%v", err.Error())
		}
	}
	// 创建题目成功之后，需要支持特判
	if req.SpjJudge != nil && *req.SpjJudge == "Y" {
		err = models.SaveSpjSource(problem, req.SpjSource)
		if err != nil {
			logger.Errorf("call AddSpjSource failed,PID=%v,req.SpjSource=%v", problem.PID, req.SpjSource)
			return nil, err
		}
	}
	//成功返回
	return response.AddProblemResp{
		Response: response.CreateResponse(constanct.SuccessCode),
		PID:      problem.PID,
	}, nil
}

func EditProblem(req *request.EditProblemReq, c *gin.Context) (interface{}, error) {
	problem := mapping.ProblemReqToDao(request.Problem(*req))
	if req.PID == nil || *req.PID == "" {
		return response.CreateResponse(constanct.PROBLEM_EDIT_PIDNoteExistCode), nil
	}
	err := models.EditProblem(c, &problem)
	if err != nil {
		//日志报错
		utils.GetLogInstance().Errorf("call EditProblem failed,err=%s", err.Error())
		return response.CreateResponse(constanct.PROBLEM_EDIT_FAILED), err
	}
	if req.SpjSource != nil && *req.SpjSource != "" {
		models.SaveSpjSource(problem, req.SpjSource)
	}
	return response.CreateResponse(constanct.SuccessCode), nil
}

func DeleteProblem(ctx *gin.Context, req *request.DeleteProblemReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	for _, PID := range req.PIDs {
		err := models.DeleteProblem(ctx, PID)
		if err != nil {
			logger.Errorf("call DeleteProblem failed,err=%s", err.Error())
			return response.CreateResponse(constanct.PROBLEM_DELETE_FAILED), err
		}
	}
	return response.CreateResponse(constanct.SuccessCode), nil
}

func GetProblemList(ctx *gin.Context, req *request.ProblemListReq) (interface{}, error) {

	var ret response.ProblemListResp
	offset, size := utils.GetPageInfo(req.Page, req.Limit)
	admin := middlewares.CheckUserHasPermission(ctx, constanct.ProblemAdmin)
	problem := dao.Problem{}
	if !admin {
		problem.Visible = 1
	}
	if req.PType == nil {
		problem.PType = constanct.LOCALTYPE
	} else {
		problem.PType = *req.PType
	}
	problems, err := models.GetProblemList(ctx, offset, size, problem)
	if err != nil {
		return response.CreateResponse(constanct.PROBLEM_LIST_FAILED), err
	}
	ret.Response = response.CreateResponse(constanct.SuccessCode)
	ret.Count, _ = models.GetProblemCount(ctx, problem)
	ret.Data = make([]response.ProblemItemResp, 0, len(problems))
	for _, problem := range problems {
		ret.Data = append(ret.Data, response.ProblemItemResp{
			PID:   problem.PID,
			Title: problem.Title,
			Label: problem.Label,
		})
	}
	return ret, nil
}

func GetProblemInfo(ctx *gin.Context, PID string) (interface{}, error) {
	logger := utils.GetLogInstance()
	if !models.IsProblemExistByPID(ctx, &dao.Problem{PID: PID}) {
		return response.CreateResponse(constanct.PROBLEM_GET_PIDNotExistCode), nil
	}
	problem, err := models.GetProblemByPID(ctx, PID)
	if err != nil {
		logger.Errorf("call GetProblemByPID failed,err=%s", err.Error())
		return nil, err
	}
	number, err := models.GetSolutionNumberByPID(ctx, PID)
	if err != nil {
		logger.Errorf("call GetSolutionNumberByPID failed,err=%s", err.Error())
		return nil, err
	}
	admin := middlewares.CheckUserHasPermission(ctx, constanct.ProblemAdmin)
	/*1 可视 -1 不可见*/
	if problem.Visible == -1 && !admin {
		return response.CreateResponse(constanct.PROBLEM_GET_PIDNotExistCode), nil
	}
	return response.ProblemInfoResp{
		Response:       response.CreateResponse(constanct.SuccessCode),
		ProblemResp:    response.ProblemResp(problem),
		SolutionNumber: number,
	}, nil
}

func DownloadProblemFromJson(ctx *gin.Context, PIDs string) (interface{}, error) {
	logger := utils.GetLogInstance()
	PIDArray := strings.Split(PIDs, ",")
	problems := make(mapping.JsonProblems, 0)
	for _, PID := range PIDArray {
		ok := models.IsProblemExistByPID(ctx, &dao.Problem{PID: PID})
		if !ok {
			logger.Errorf("the problem not exist, PID:%v", PID)
			return response.CreateResponse(constanct.PROBLEM_DOWNLOADPROBLE_PIDNoteExistCode), nil
		}
		Jproblem, err := models.ParseProblemToJsonProblem(ctx, PID)
		if err != nil {
			logger.Errorf("call ParseDBToJson failed, req=%+v, err=%s", utils.Sdump(PID), err)
			response.ResponseError(ctx, constanct.PROBLEM_DOWNLOADPROBLE_FAILEDCode)
			return nil, err
		}
		problems = append(problems, Jproblem)
	}
	str, err := models.ParseJsonProblemToJson(ctx, problems)
	if err != nil {
		logger.Errorf("call ParseJsonProblemToJson failed, req=%+v, err=%s", utils.Sdump(PIDs), err)
		response.ResponseError(ctx, constanct.PROBLEM_DOWNLOADPROBLE_FAILEDCode)
	}
	return str, nil
}

func UpProblemFile(ctx *gin.Context, file *multipart.FileHeader) (interface{}, error) {
	logger := utils.GetLogInstance()
	suffix := strings.ToUpper(utils.GetFileSuffix(file.Filename))
	fd, _ := file.Open()
	reader, _ := ioutil.ReadAll(fd)

	data := string(reader)
	parseProblemFunc, ok := ParseProblemFuncMap[suffix]
	if !ok {
		logger.Errorf("Uploading %v files is prohibited", suffix)
		return response.CreateResponseStr(constanct.AUTH_LOGIN_PassEmptyCode, fmt.Sprintf("Uploading %v files is prohibited", suffix), response.ERROR), nil
	}
	JsonProblems, err := parseProblemFunc(ctx, data)
	if err != nil {
		logger.Errorf("call parseProblemFunc failed err=%v", err.Error())
		return nil, err
	}
	for _, jsonProblem := range JsonProblems {
		err = mapping.SaveJproblemImage(&jsonProblem)
		if err != nil {
			return nil, err
		}
		problem := mapping.JsonProblemToProblem(jsonProblem)
		idx, err := models.GetNextProblemPID(ctx)
		PID := "P" + idx
		if err != nil {
			return nil, err
		}
		problem.PID = PID
		err = models.CreateProblem(ctx, &problem)
		if err != nil {
			return nil, err
		}
		if PID != "" {
			err = rediscache.UpdateNextPID(ctx, PID)
			if err != nil {
				logger.Errorf("call UpdateNextPID failed,err:%v", err.Error())
			}
		}
		err = models.CreateTestFile(ctx, PID, jsonProblem.Data)
		if err != nil {
			return nil, err
		}
	}
	return response.CreateResponse(constanct.SuccessCode), nil
}
