package logic

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	redisdao "ahutoj/web/dao/redisDao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/mapping"
	"ahutoj/web/middlewares"
	"ahutoj/web/models"
	"ahutoj/web/utils"
	"sync"

	"github.com/gin-gonic/gin"
)

var AddProblemLock sync.Mutex

func AddProblem(req *request.Problem, c *gin.Context) (interface{}, error) {
	var nextPID string = ""
	var err error
	AddProblemLock.Lock()
	defer AddProblemLock.Unlock()
	logger := utils.GetLogInstance()
	problem := mapping.ProblemReqToDao(*req)
	if !models.ChekckProblemType(c, req.PType) {
		return response.CreateResponse(constanct.Problem_ADD_PTYPEERR_CODE), nil
	}
	if req.PType == "" {
		req.PType = constanct.LOCALTYPE
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
		err = redisdao.UpdateNextPID(c, nextPID)
		if err != nil {
			logger.Errorf("call UpdateNextPID failed,err:%v", err.Error())
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
	admin := middlewares.CheckUserHasPermission(ctx, middlewares.ProblemAdmin)
	problem := dao.Problem{}
	if !admin {
		problem.Visible = 1
	}
	problems, err := mysqldao.SelectProblemByLists(ctx, offset, size, problem)
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
	if !models.IsProblemExistByPID(ctx, &dao.Problem{PID: PID}) {
		return response.CreateResponse(constanct.PROBLEM_GET_PIDNotExistCode), nil
	}
	problem, err := models.GetProblemByPID(ctx, PID)
	if err != nil {
		return nil, err
	}
	admin := middlewares.CheckUserHasPermission(ctx, middlewares.ProblemAdmin)
	/*1 可视 -1 不可见*/
	if problem.Visible == -1 && !admin {
		return response.CreateResponse(constanct.PROBLEM_GET_PIDNotExistCode), nil
	}
	return response.ProblemInfoResp{
		Response:    response.CreateResponse(constanct.SuccessCode),
		ProblemResp: response.ProblemResp(problem),
	}, nil
}
