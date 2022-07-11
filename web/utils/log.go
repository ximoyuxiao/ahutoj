package utils

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func LogInit() {

	config := GetConfInstance()
	level, err := logrus.ParseLevel(config.LogConfig.Level)
	if err != nil {
		fmt.Printf("log level(%s) error, err =%s\n", (config.LogConfig.Level), err.Error())
	}
	logger.SetLevel(level)
	logger.SetOutput(os.Stdout)
	if config.LogConfig.FileName != "console" {
		writer, err := os.OpenFile(config.LogConfig.FileName, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			fmt.Printf("call OpenFile failed, logname=%s, err=%s", config.LogConfig.FileName, err.Error())
		}
		logger.SetOutput(writer)
	}
}
func GetLogInstance() *logrus.Logger {
	return logger
}
