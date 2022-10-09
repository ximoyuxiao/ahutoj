package models

import (
	mysqldao "ahutoj/web/dao/mysqlDao"
	"context"
)

func FindSubmitCeInfo(ctx context.Context, SID int64) string {
	ceinfo, _ := mysqldao.SelectCeinfoBySID(ctx, SID)
	return ceinfo.Info
}
