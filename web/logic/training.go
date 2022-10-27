package logic

import (
	"ahutoj/web/dao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/middlewares"
	"ahutoj/web/models"
	"ahutoj/web/utils"

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
		return response.CreateResponseStr(constanct.GetResCode(constanct.Training, constanct.Models, constanct.MysqlAdd), "创建题单失败", response.ERROR), err
	}
	list.LID, err = models.GetCurrentLID(c, list)
	if err != nil {
		//日志报错
		utils.GetLogInstance().Errorf("call GetLID in CreateList failed,err=%s", err.Error())
		return response.CreateResponseStr(constanct.GetResCode(constanct.Training, constanct.Models, constanct.MysqlQuery), "获取题单ID失败", response.ERROR), err
	}

	//添加提单题目
	err = models.CreateListProblem(c, &listproblem)
	if err != nil {
		//日志报错
		utils.GetLogInstance().Errorf("call CreateListProblem failed,err=%s", err.Error())
		return response.CreateResponseStr(constanct.GetResCode(constanct.Training, constanct.Models, constanct.MysqlAdd), "添加题单内题目失败", response.ERROR), err
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
		return response.CreateResponseStr(constanct.GetResCode(constanct.Training, constanct.Models, constanct.MysqlUpdate), "修改题单信息失败", response.ERROR), err
	}
	//编辑提单题目
	err2 := models.EditListProblem(c, &listproblem)
	if err2 != nil {
		//日志报错
		utils.GetLogInstance().Errorf("call EditListProblem failed,err=%s", err2.Error())
		return response.CreateResponseStr(constanct.GetResCode(constanct.Training, constanct.Models, constanct.MysqlUpdate), "修改题单题目信息失败", response.ERROR), err2
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
		return response.CreateResponseStr(constanct.GetResCode(constanct.Training, constanct.Models, constanct.MysqlDelete), "删除题单失败", response.ERROR), err
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
		return response.CreateResponseStr(constanct.GetResCode(constanct.Training, constanct.Models, constanct.MysqlQuery), "获取题单列表失败", response.ERROR), err
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
