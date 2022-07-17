package mysqldao

import (
	"ahutoj/web/utils"
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	ConnectTimeout          = 30 * time.Second
	globalDB       *gorm.DB = nil
	err            error
)

func InitMysql() error {
	cfg := utils.GetConfInstance().MySQLConfig
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.UserName, cfg.Password, cfg.Host, cfg.Port, cfg.Dbname,
	)
	globalDB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}

	return nil
}

func GetDB(ctx context.Context) *gorm.DB {
	return globalDB.WithContext(ctx)
}