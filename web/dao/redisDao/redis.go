package redisdao

import (
	"ahutoj/web/utils"
)

func InitRedis() error {
	logger := utils.GetLogInstance()
	logger.Info("Redis cnn't Init")
	return nil
}
