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

func AddProblem(req *request.Problem, c *gin.Context) (interface{}, error) {
	problem := dao.Problem{
		Title:         req.Title,
		Description:   req.Description,
		Input:         req.Input,
		Output:        req.Output,
		Sample_input:  req.Sample_input,
		Sample_output: req.Sample_output,
		Hit:           req.Hit,
		Label:         req.Label,
		LimitTime:     req.LimitTime,
		LimitMemory:   req.LimitMemory,
	}
	//题目不存在 添加题目
	err := models.CreateProblem(c, &problem)
	if err != nil {
		//日志报错
		utils.GetLogInstance().Errorf("call CreateProblem failed,err=%s", err.Error())
		return response.CreateResponse(constanct.MySQLErrorCode), err
	}
	//成功返回
	return response.CreateResponse(constanct.SuccessCode), nil
}
func EditProblem(req *request.Problem, c *gin.Context) (interface{}, error) {
	problem := dao.Problem{
		Title:         req.Title,
		Description:   req.Description,
		Input:         req.Input,
		Output:        req.Output,
		Sample_input:  req.Sample_input,
		Sample_output: req.Sample_output,
		Hit:           req.Hit,
		Label:         req.Label,
		LimitTime:     req.LimitTime,
		LimitMemory:   req.LimitMemory,
	}
	err := models.EditProblem(c, &problem)
	if err != nil {
		//日志报错
		utils.GetLogInstance().Errorf("call EditProblem failed,err=%s", err.Error())
		return response.CreateResponse(constanct.MySQLErrorCode), err
	}
	return response.CreateResponse(constanct.SuccessCode), nil
}

func DeleteProblem(ctx *gin.Context, req *request.DeleteProblemReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	for _, pid := range req.Pids {
		err := models.DeleteProblem(ctx, pid)
		if err != nil {
			logger.Errorf("call DeleteProblem failed,err=%s", err.Error())
			return nil, err
		}
	}
	return response.CreateResponse(constanct.SuccessCode), nil
}

func GetProblemList(ctx *gin.Context, req *request.ProblemListReq) (interface{}, error) {
	var ret response.ProblemListResp
	offset, size := utils.GetPageInfo(req.Page, req.Limit)
	problems, err := mysqldao.SelectProblemByLists(ctx, offset, size)
	if err != nil {
		return nil, err
	}
	ret.Response = response.CreateResponse(constanct.SuccessCode)
	ret.Count, _ = models.GetProblemCount(ctx)
	ret.Data = make([]response.ProblemItemResp, 0, len(problems))
	for _, problem := range problems {
		ret.Data = append(ret.Data, response.ProblemItemResp{
			Pid:   problem.Pid,
			Title: problem.Title,
			Label: problem.Label,
		})
	}
	return ret, nil
}

func GetProblemInfo(ctx *gin.Context, pid int64) (interface{}, error) {
	if !models.IsProblemExistByPid(ctx, &dao.Problem{Pid: int(pid)}) {
		return response.CreateResponse(constanct.PIDNotExistCode), nil
	}
	problem, err := models.GetProblemByPID(ctx, pid)
	if err != nil {
		return nil, err
	}
	return response.ProblemInfoResp{
		Response:    response.CreateResponse(constanct.SuccessCode),
		ProblemResp: response.ProblemResp(problem),
	}, nil

}
