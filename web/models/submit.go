package models

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	redisdao "ahutoj/web/dao/redisDao"
	"ahutoj/web/utils"
	"context"

	"github.com/bytedance/gopkg/util/logger"
)

func GetSubmitByCIDFromDB(ctx context.Context, CID, fb int64) ([]dao.Submit, error) {
	res, err := mysqldao.SelectSubmitByCID(ctx, CID, fb)
	if err != nil {
		// response.CreateResponse(constanct.GetResCode(constanct.Submit, constanct.Models, constanct.MysqlQuery))
	}
	return res, err
}

func CreateSubmit(ctx context.Context, submit dao.Submit) error {
	err := mysqldao.InsertSubmit(ctx, submit)
	if err != nil {
		// response.CreateResponse(constanct.GetResCode(constanct.Submit, constanct.Models, constanct.MysqlAdd))
	}
	return err
}

func GetSubmitList(ctx context.Context, submit dao.Submit, offset, limit int) ([]dao.Submit, error) {
	res, err := mysqldao.SelectSubmitList(ctx, submit, offset, limit)
	if err != nil {
		// response.CreateResponse(constanct.GetResCode(constanct.Submit, constanct.Models, constanct.MysqlQuery))
	}
	return res, err
}

func RejudgeSubmit(ctx context.Context, submit dao.Submit) error {
	err := mysqldao.RejudgeSubmit(ctx, submit)
	if err != nil {
		// response.CreateResponse(constanct.GetResCode(constanct.Submit, constanct.Models, constanct.MysqlUpdate))
	}
	return err
}

func GetSubmitListCount(ctx context.Context, submit dao.Submit) (int64, error) {
	res, err := mysqldao.SelectCountSubmit(ctx, submit)
	if err != nil {
		// response.CreateResponse(constanct.GetResCode(constanct.Submit, constanct.Models, constanct.MysqlQuery))
	}
	return res, err
}
func GetUserStatusInfo(ctx context.Context, submit dao.Submit, lastTime int64) ([]dao.Submit, error) {
	res, err := mysqldao.SelectSubmitByUID(ctx, submit, lastTime)
	if err != nil {
		// response.CreateResponse(constanct.GetResCode(constanct.Submit, constanct.Models, constanct.MysqlQuery))
	}
	return res, err
}

func UpdateSubmit(ctx context.Context, submit dao.Submit) error {
	err := mysqldao.UpdateSubmit(ctx, submit)
	if err != nil {
		// response.CreateResponse(constanct.GetResCode(constanct.Submit, constanct.Models, constanct.MysqlUpdate))
	}
	return err
}

func GetOriginJudgeSubmit(ctx context.Context) ([]dao.Submit, error) {
	res, err := mysqldao.SelectSubmitIsOriginJudge(ctx)
	if err != nil {
		// response.CreateResponse(constanct.GetResCode(constanct.Submit, constanct.Models, constanct.MysqlQuery))
	}
	return res, err
}

func FindLastSIDByUID(ctx context.Context, UID string) (dao.Submit, error) {
	res, err := mysqldao.FindLastSIDByUID(ctx, UID)
	if err != nil {
		// response.CreateResponse(constanct.GetResCode(constanct.Submit, constanct.Models, constanct.MysqlQuery))
	}
	return res, err
}

func EqualLastSource(ctx context.Context, UID string, PID int64, Source string) bool {
	SourceMD5, err := utils.MD5EnCodeStr(Source)
	if err != nil {
		logger.Errorf("call MD5EnCodeStr failed. Source:%s", Source)
		return false
	}
	Lastsource := redisdao.GetLastSource(ctx, UID, PID)
	if Lastsource == SourceMD5 {
		return true
	}
	redisdao.SetLastSource(ctx, UID, PID, SourceMD5)
	return false
}
