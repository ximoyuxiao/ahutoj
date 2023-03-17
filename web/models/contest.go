package models

import (
	rediscache "ahutoj/web/cache/redis"
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/utils"
	"context"
	"time"
)

func AddContestToDb(ctx context.Context, contest dao.Contest) error {
	return mysqldao.InserContest(ctx, contest)
}
func GetCurrentCID(ctx context.Context, contest dao.Contest) (int64, error) {
	return mysqldao.SelectContestByUID(ctx, contest.UID)
}
func SaveContestDB(ctx context.Context, contest dao.Contest) error {
	logger := utils.GetLogInstance()
	/*对mysql先操作*/
	err := mysqldao.UpdateContest(ctx, contest)
	if err != nil {
		logger.Errorf("call mysqldao.UpdateContest faile,contest=%v ,err=%s", utils.Sdump(contest), err.Error())
		return err
	}
	/*去redis获取缓存 查看有没有缓存*/
	contestTmp, err := rediscache.GetContestFromDB(ctx, contest.CID)
	if err != nil {
		logger.Errorf("call rediscache.GetContestFromDB faile,contest=%v ,err=%s", utils.Sdump(contest), err.Error())
		return err
	}
	/*redis更新数据库内容*/
	if contestTmp != nil {
		return rediscache.SaveContestToRDB(ctx, contest)
	}
	return nil
}

func DeleteContestDB(ctx context.Context, CID int64) error {
	// logger := utils.GetLogInstance()
	return mysqldao.DeleteContest(ctx, CID)
}

func GetContestListFromDb(ctx context.Context, offset, pagesize int) ([]dao.Contest, error) {
	return mysqldao.SelectContests(ctx, offset, pagesize)
}

/*
如果 当前这竞赛ID 是在进行中 他在这段时间的访问的频率就会比较高
那我们就对这个竞赛的访问做一层缓存 到redis数据库当中
中间 有人去修改了 竞赛信息
redis的数据 和mysql 数据  要保持一致性

	先 改mysql  我查看redis数据库 当中 是否有缓存  有的话 我就把缓存改掉。。
*/
func GetContestFromDB(ctx context.Context, CID int64) (dao.Contest, error) {
	/*比赛，进行中一般不会修改*/
	logger := utils.GetLogInstance()
	/* 从redis当中尝试获取一个缓存*/
	contest, err := rediscache.GetContestFromDB(ctx, CID)
	if err != nil || contest == nil {
		/*从mysql当中获取*/
		contest, err = mysqldao.SelectContestByCID(ctx, CID)
		if err != nil {
			logger.Errorf("call SelectContestByCID failed,CID=%v, err=%v", CID, err.Error())
			return dao.Contest{}, err
		}
		/*去做缓存 策略  判断需不需要作缓存 需要的就去给他做一个缓存*/
		now := time.Now().UnixMilli()
		if contest.End_time >= now && contest.Begin_time <= now {
			logger.Debugf("now=%v begin=%v endtime=%v", now, contest.Begin_time, contest.End_time)
			err = rediscache.SaveContestToRDB(ctx, *contest)
		}
	}
	return *contest, err
}

func GetContestCountFromDB(ctx context.Context) (int64, error) {
	return mysqldao.SelectContestCount(ctx)
}
