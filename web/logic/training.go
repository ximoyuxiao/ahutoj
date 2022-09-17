package logic

import (
	"ahutoj/web/dao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/models"
	"ahutoj/web/utils"

	"github.com/gin-gonic/gin"
)

func AddTraining(req *request.ListAll, c *gin.Context) (interface{}, error) {

	list := dao.List{
		LID:       req.LID,
		UID:       req.UID,
		Title:     req.Title,
		StartTime: req.Stime,
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
		return response.CreateResponse(constanct.MySQLErrorCode), err
	}
	//添加提单题目
	err2 := models.CreateListProblem(c, &listproblem)
	if err2 != nil {
		//日志报错
		utils.GetLogInstance().Errorf("call CreateListProblem failed,err=%s", err2.Error())
		return response.CreateResponse(constanct.MySQLErrorCode), err2
	}
	return response.CreateResponse(constanct.SuccessCode), nil
}
func EditTraining(req *request.ListAll, c *gin.Context) (interface{}, error) {
	list := dao.List{
		LID:       req.LID,
		UID:       req.UID,
		Title:     req.Title,
		StartTime: req.Stime,
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
		return response.CreateResponse(constanct.MySQLErrorCode), err
	}
	//编辑提单题目
	err2 := models.EditListProblem(c, &listproblem)
	if err2 != nil {
		//日志报错
		utils.GetLogInstance().Errorf("call EditListProblem failed,err=%s", err2.Error())
		return response.CreateResponse(constanct.MySQLErrorCode), err2
	}
	return response.CreateResponse(constanct.SuccessCode), nil
}
func DeleteTraining(req *request.List, c *gin.Context) (interface{}, error) {
	list := dao.List{
		LID:       req.LID,
		UID:       req.UID,
		Title:     req.Title,
		StartTime: req.Stime,
	}
	//删除题单
	err := models.DeleteTraining(c, &list)
	if err != nil {
		//日志报错
		utils.GetLogInstance().Errorf("call DeleteList failed,err=%s", err.Error())
		return response.CreateResponse(constanct.MySQLErrorCode), err
	}
	return response.CreateResponse(constanct.SuccessCode), nil
}
