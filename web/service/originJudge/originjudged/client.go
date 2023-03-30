package originJudged

import "net/http"

type UserStatus int64

const (
	JUDGE_FREE UserStatus = 1
	JUDGE_BUSY UserStatus = 2
)

func SetCookies(resp *http.Response, p *OriginJudgeUser) error {
	if resp != nil && p != nil {
		cookies := resp.Cookies()
		for _, cookie := range cookies {
			p.Cookies[cookie.Name] = cookie.Value
		}
	}
	return nil
}
