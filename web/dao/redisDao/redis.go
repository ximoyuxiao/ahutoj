package redisdao

import (
	"ahutoj/web/utils"
)

func InitRedis() error {
	logger := utils.GetLogInstance()
	logger.Info("Redis cann't Init")
	return nil
}
