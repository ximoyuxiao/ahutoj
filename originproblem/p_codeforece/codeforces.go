package pcodeforece

import (
	originJudged "ahutoj/originJudge/originjudged"
	"ahutoj/web/dao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/models"
	"ahutoj/web/utils"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"time"
)

var mainURL = "https://www.luogu.com.cn/problem/list?type=CF&page="
var problemURL = "https://www.luogu.com.cn/problem/"

var startPID = "CF7C"
var page = 1

type CFProblemObj struct {
	Code            int    `json:"code"`
	CurrentTemplate string `json:"currentTemplate"`
	CurrentData     struct {
		Problem struct {
			Background   string     `json:"background"`
			Description  string     `json:"description"`
			InputFormat  string     `json:"inputFormat"`
			OutputFormat string     `json:"outputFormat"`
			Samples      [][]string `json:"samples"`
			Hint         string     `json:"hint"`
			Provider     struct {
				UID        int    `json:"uid"`
				Name       string `json:"name"`
				Slogan     string `json:"slogan"`
				Badge      string `json:"badge"`
				IsAdmin    bool   `json:"isAdmin"`
				IsBanned   bool   `json:"isBanned"`
				Color      string `json:"color"`
				CcfLevel   int    `json:"ccfLevel"`
				Background string `json:"background"`
			} `json:"provider"`
			Attachments []interface{} `json:"attachments"`
			CanEdit     bool          `json:"canEdit"`
			Limits      struct {
				Time   []int `json:"time"`
				Memory []int `json:"memory"`
			} `json:"limits"`
			StdCode string `json:"stdCode"`
			Vjudge  struct {
				Origin string `json:"origin"`
				Link   string `json:"link"`
				ID     string `json:"id"`
			} `json:"vjudge"`
			Translation      string `json:"translation"`
			Tags             []int  `json:"tags"`
			WantsTranslation bool   `json:"wantsTranslation"`
			TotalSubmit      int    `json:"totalSubmit"`
			TotalAccepted    int    `json:"totalAccepted"`
			Flag             int    `json:"flag"`
			Pid              string `json:"pid"`
			Title            string `json:"title"`
			Difficulty       int    `json:"difficulty"`
			FullScore        int    `json:"fullScore"`
			Type             string `json:"type"`
		} `json:"problem"`
		Contest     interface{} `json:"contest"`
		Discussions []struct {
			ID    int    `json:"id"`
			Title string `json:"title"`
			Forum struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Slug string `json:"slug"`
			} `json:"forum"`
		} `json:"discussions"`
		Bookmarked      bool          `json:"bookmarked"`
		VjudgeUsername  interface{}   `json:"vjudgeUsername"`
		Recommendations []interface{} `json:"recommendations"`
		LastLanguage    int           `json:"lastLanguage"`
		LastCode        string        `json:"lastCode"`
		PrivilegedTeams []interface{} `json:"privilegedTeams"`
		UserTranslation interface{}   `json:"userTranslation"`
	} `json:"currentData"`
	CurrentTitle string      `json:"currentTitle"`
	CurrentTheme interface{} `json:"currentTheme"`
	CurrentTime  int         `json:"currentTime"`
}

func parsePID(originPID string) (string, string, error) {
	if len(originPID) < 1 {
		return "", "", errors.New("")
	}
	re, _ := regexp.Compile("([0-9]*)([A-Z]*[0-9]*)")
	strs := re.FindStringSubmatch(originPID)
	if len(strs) < 3 {
		return "", "", fmt.Errorf("ParsePID failed ,pid:%v", originPID)
	}
	return strs[1], strs[2], nil
}
func GetCurrentPageProblems() ([]string, error) {
	url := fmt.Sprintf("%v%v", mainURL, page)
	resp, err := originJudged.DoRequest(originJudged.GET, url, nil, nil, nil, false)
	if err != nil {
		return nil, fmt.Errorf("call Dorequest Failed,err:%v", err.Error())
	}
	body, err := originJudged.ParseRespToByte(resp)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp.StatusCode)
	re := regexp.MustCompile(`<li>(.*?)&nbsp;<a href=".*?">.*?</a></li>`)
	ret := re.FindAllStringSubmatch(string(body), -1)
	ansstring := make([]string, 0)
	for _, ans := range ret {
		ansstring = append(ansstring, ans[1])
		fmt.Println(ans[1])
	}
	page++
	return ansstring, nil
}

func CFproblemToMyProblem(CFproblem CFProblemObj) dao.Problem {
	ret := dao.Problem{}
	ret.PID = CFproblem.CurrentData.Problem.Pid
	ret.PType = constanct.CODEFORCESTYPE
	ret.Title = CFproblem.CurrentData.Problem.Title
	SamipleInPutAndOutPut := ""
	for idx := range CFproblem.CurrentData.Problem.Samples {
		SamipleInPutAndOutPut += fmt.Sprintf("### 样例输入 #%v\n\n```\n%v\n```\n\n### 样例输出 #%v\n\n```\n%v\n```\n\n",
			idx,
			CFproblem.CurrentData.Problem.Samples[idx][0],
			idx,
			CFproblem.CurrentData.Problem.Samples[idx][1],
		)
	}
	description := fmt.Sprintf("## 题目描述\n\n%v\n\n## 输入格式\n\n%v\n\n## 输出格式\n\n%v\n\n题目样例\n\n%v\n\n## 提示\n\n%v",
		CFproblem.CurrentData.Problem.Description,
		CFproblem.CurrentData.Problem.InputFormat,
		CFproblem.CurrentData.Problem.OutputFormat,
		SamipleInPutAndOutPut,
		CFproblem.CurrentData.Problem.Hint,
	)
	ret.LimitTime = int64(CFproblem.CurrentData.Problem.Limits.Time[0])
	ret.LimitMemory = int64(CFproblem.CurrentData.Problem.Limits.Memory[0] / 1024)
	ret.ContentType = 1
	ret.Origin = int64(originJudged.Cfoj)
	ret.Label = "codeforce"
	ret.Visible = 1
	ret.OriginPID = ret.PID[2:]
	CFCID, CFCIDX, _ := parsePID(ret.OriginPID)
	ret.Description = description + "\n" + fmt.Sprintf(`<a href="https://codeforces.com/contest/%v/problem/%v" target="_blank">点击查看原题</a>`,
		CFCID,
		CFCIDX,
	)
	return ret
}

func ParseCodeForceProblemURL(PID string) (dao.Problem, error) {
	Purl := fmt.Sprintf("%v%v", problemURL, PID)
	resp, err := originJudged.DoRequest(originJudged.GET, Purl, nil, nil, nil, false)
	if err != nil {
		return dao.Problem{}, fmt.Errorf("call Dorequest Failed,err:%v", err.Error())
	}
	body, err := originJudged.ParseRespToByte(resp)
	if err != nil {
		return dao.Problem{}, err
	}
	// re := regexp.MustCompile(`<article>(?s:(.*?))</article>`)
	re := regexp.MustCompile(`JSON.parse\(decodeURIComponent\("(?s:(.*?))"\)`)
	jsonCode := re.FindSubmatch(body)
	if jsonCode == nil {
		return dao.Problem{}, fmt.Errorf("parse html failed")
	}
	Pjson, _ := url.QueryUnescape(string(jsonCode[1]))
	CFproblem := CFProblemObj{}
	json.Unmarshal([]byte(Pjson), &CFproblem)
	myproblem := CFproblemToMyProblem(CFproblem)
	return myproblem, nil
}

func DownLoadAllCodeForceProblem() {
	var canDown = false
	if startPID == "" {
		canDown = true
	}
	for {
		PIDs, _ := GetCurrentPageProblems()
		if len(PIDs) == 0 {
			println("PIDs==0")
			return
		}
		if !canDown {
			for _, PID := range PIDs {
				if canDown {
					fmt.Printf("正在获取%v的题目信息\n", PID)
					problem, _ := ParseCodeForceProblemURL(PID)
					fmt.Println(utils.Sdump(problem))
					models.CreateProblem(context.Background(), &problem)
					time.Sleep(time.Second)
				}
				if PID == startPID {
					canDown = true
				}
			}
			time.Sleep(1 * time.Second)
			continue
		}
		if canDown {
			for _, PID := range PIDs {
				fmt.Printf("正在获取%v的题目信息\n", PID)
				problem, err := ParseCodeForceProblemURL(PID)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(problem)
				models.CreateProblem(context.Background(), &problem)
				time.Sleep(1 * time.Second)
			}
		}
	}

}
