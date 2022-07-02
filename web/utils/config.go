package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

var config = new(AppConfig)

type AppConfig struct {
	Port         string `mapstructure:"port"`
	Mode         string `mapstructure:"mode"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
	*LogConfig   `mapstructure:"log"`
}

type MySQLConfig struct {
	Driver   string `mapstructure:"driver"`
	Host     string `mapstructure:"host"`
	Port     int16  `mapstructure:"port"`
	UserName string `mapstructure:"username"`
	Dbname   string `mapstructure:"dbname"`
	Password string `password:"password"`
}
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int16  `mapstructure:"port"`
	Db       int16  `mapstructure:"db"`
	PoolSize int16  `mapstructure:"pool_size"`
}
type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int32  `mapstructure:"max_size"`
	MaxAge     int32  `mapstructure:"max_age"`
	MaxBackups int32  `mapstructure:"max_backups"`
}

func ConfigInit(configPath string) error {
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("call ReadConfig Failed, err=%s\n", err.Error())
		return err
	}
	if err := viper.Unmarshal(config); err != nil {
		fmt.Printf("call Unmarshal Failed, err=%s\n", err.Error())
		return err
	}
	return nil

}

func GetInstance() *AppConfig {
	return config
}
