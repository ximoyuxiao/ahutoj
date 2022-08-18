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

func AddTraining(req *request.List, req2 *request.ListProblem, c *gin.Context) (interface{}, error) {

	list := dao.List{
		Lid:   req.Lid,
		Uid:   req.Uid,
		Title: req.Title,
		Stime: req.Stime,
	}
	listproblem := dao.ListProblem{
		Lid: req2.Lid,
		Pid: req2.Pid,
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
