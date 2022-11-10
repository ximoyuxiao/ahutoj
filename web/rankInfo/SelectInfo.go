package rankinfo

// import (
// 	"ahutoj/web/models"
// 	originjudge "ahutoj/web/originJudge"
// 	"context"
// 	"time"
// )

// func UpdateUserInfoThread(ctx context.Context) {
// 	/*每个小时 去更新一次  用户的做题数据*/
// 	for {
// 		UpdateCodeForceTable(ctx)
// 		time.Sleep(time.Hour)
// 	}
// }

//	func UpdateCodeForceTable(ctx context.Context) {
//		users, _ := models.FindAllUser(ctx)
//		for _, user := range users {
//			cfID := user.CodeForceUser
//			url := "https://codeforces.com/submissions/" + cfID
//			originjudge.DoRequest(originjudge.GET, url, originjudge.CfHeaders, nil, nil, false)
//		}
//	}
