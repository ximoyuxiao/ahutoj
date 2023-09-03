package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

var config = new(AppConfig)

type AppConfig struct {
	Host            string  `mapstructure:"host"`
	Port            string  `mapstructure:"port"`
	GatWayHost      string  `mapstructure:"gatWayHost"`
	Mode            string  `mapstructure:"mode"`
	Sign            string  `mapstructure:"sign"`
	Version         string  `mapstructure:"version"`
	StartTime       string  `mapstructure:"startTime"`
	MachineID       int64   `mapstructure:"machineID "`
	DataPath        string  `mapstructure:"dataPath"`
	JsonPath        string  `mapstructure:"jsonPath"`
	Terminal        float64 `mapstructure:"terminal"`
	OpenTime        float64 `mapstructure:"openTime"`
	ImagePath       string  `mapstructure:"imagePath"`
	HeadPath        string  `mapstructure:"headPath"`
	SpjPath         string  `mapstructure:"spjPath"`
	OpenRegisiter   bool    `mapstructure:"openRegisiter"`
	*MySQLConfig    `mapstructure:"mysql"`
	*RedisConfig    `mapstructure:"redis"`
	*LogConfig      `mapstructure:"log"`
	UseOriginJudge  bool `mapstructure:"useOriginJudge"`
	*AtCoderJudges  `mapstructure:"atCoderJudges"`
	*CodeForceJudge `mapstructure:"codeForceJudge"`
	*LuoguJudge     `mapstructure:"luoguJudge"`
	*RabbitMQ       `mapstructure:"rabbitmq"`
	*OssConfig      `mapstructure:"oss"`
}
type OssConfig struct {
	Host     string `mapstructure:"host"`
	Port     uint16 `mapstructure:"port"`
	BasePath string `mapstructure:"basepath"`
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
type AtCoderJudges struct {
	Count    int64  `mapstructure:"count"`
	Prefix   string `mapstructure:"prefix"`
	Password string `mapstructure:"password"`
}

type CodeForceJudge struct {
	Count    int64  `mapstructure:"count"`
	Prefix   string `mapstructure:"prefix"`
	Password string `mapstructure:"password"`
}
type LuoguJudge struct {
	Count    int64  `mapstructure:"count"`
	Prefix   string `mapstructure:"prefix"`
	Password string `mapstructure:"password"`
}
type RabbitMQ struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
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
func InitAppConfig(configPath string, Myconfig interface{}) error {
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("call ReadConfig Failed, err=%s\n", err.Error())
		return err
	}
	if err := viper.Unmarshal(Myconfig); err != nil {
		fmt.Printf("call Unmarshal Failed, err=%s\n", err.Error())
		return err
	}
	fmt.Println(Sdump(Myconfig))
	return nil

}

func GetConfInstance() *AppConfig {
	return config
}
