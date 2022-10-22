package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

var config = new(AppConfig)

type AppConfig struct {
	Port         string  `mapstructure:"port"`
	Mode         string  `mapstructure:"mode"`
	Sign         string  `mapstructure:"sign"`
	Version      string  `mapstructure:"version"`
	StartTime    string  `mapstructure:"startTime"`
	MachineID    int64   `mapstructure:"machineID "`
	DataPath     string  `mapstructure:"dataPath"`
	Terminal     float64 `mapstructure:"Terminal"`
	OpenTime     float64 `mapstructure:"openTime"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
	*LogConfig   `mapstructure:"log"`
}

type MySQLConfig struct {
	Driver   string `mapstructure:"driver"`
	Host     string `mapstructure:"host"`
	Port     uint16 `mapstructure:"port"`
	UserName string `mapstructure:"username"`
	Dbname   string `mapstructure:"dbname"`
	Password string `password:"password"`
}
type RedisConfig struct {
	Host      string `mapstructure:"host"`
	Port      string `mapstructure:"port"`
	Db        int    `mapstructure:"db"`
	PoolSize  int16  `mapstructure:"pool_size"`
	Password  string `mapstructure:"password"`
	KeepAlive int64  `mapstructure:"keppalive"` // 秒为单位
}
type LogConfig struct {
	FileName string `mapstructure:"filename"`
	MaxSize  int32  `mapstructure:"max_size"`
	Level    string `mapstructure:"level"`
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
	fmt.Println(Sdump(config))
	return nil

}

func GetConfInstance() *AppConfig {
	return config
}
