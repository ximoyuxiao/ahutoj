package main

import (
	originjudge "ahutoj/web/originJudge"
	"fmt"
	"regexp"
)

func main() {
	submissionUrl := "https://atcoder.jp/contests/abc267/submissions/35493843"
	for {
		resp, err := originjudge.DoRequest(originjudge.GET, submissionUrl, nil, nil, nil, false)
		if err != nil {
			continue
		}
		Text, _ := originjudge.ParseRespToByte(resp)
		re := regexp.MustCompile(`<td class="text-center">(.*?)</td>`)
		ret := re.FindAllSubmatch(Text, 9)
		fmt.Println("Time:" + string(ret[6][1]))
		fmt.Println("Memory:" + string(ret[7][1]))
		return
	}
}
