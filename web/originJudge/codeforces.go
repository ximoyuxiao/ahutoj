package originjudge

import (
	"ahutoj/web/dao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/models"
	"ahutoj/web/utils"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/bytedance/gopkg/util/logger"
)

type CFJudgeUser struct {
	OriginJudgeUser
	CsrfToken string
}

var cfLock sync.Mutex

type CFJudgeUsers []CFJudgeUser

var JudgeUsers []CFJudgeUser = nil
var Cfoj OJPlatform = 1

var cfurl = "https://codeforces.com"

var cflang = map[constanct.LANG]string{
	constanct.C:       "43",
	constanct.CPP:     "50",
	constanct.CPP11:   "50",
	constanct.CPP17:   "54",
	constanct.JAVA:    "36",
	constanct.PYTHON3: "31",
}

var CFResultMap = map[string]constanct.OJResult{
	"Accepted":                 constanct.OJ_AC,
	"Compilationerror(.*?)":    constanct.OJ_CE,
	"Memorylimitexceeded(.*?)": constanct.OJ_MLE,
	"O(.*?)":                   constanct.OJ_OLE,
	"P(.*?)":                   constanct.OJ_PE,
	"Runtimeerror(.*?)":        constanct.OJ_RE,
	"Timelimitexceeded(.*?)":   constanct.OJ_TLE,
	"Wronganswer(.*?)":         constanct.OJ_WA,
	"Running":                  constanct.OJ_JUDGE,
	"Inqueue(.*?)":             constanct.OJ_JUDGE,
}

var CfHeaders = map[string]string{
	"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:105.0) Gecko/20100101 Firefox/105.0",
	"Accept":          "*/*",
	"Accept-Encoding": "gzip, deflate, br",
	"Origin":          "https://codeforces.com",
	"Connection":      "keep-alive",
	"Content-Type":    "application/x-www-form-urlencoded",
	"Host":            "www.codeforces.com",
}

type CodeForceJudge struct {
	Headers      map[string]string
	judgeUser    *CFJudgeUser
	loginSuccess bool
	OriginJudge
}

func (p CodeForceJudge) Judge(ctx context.Context, submit dao.Submit, PID string) error {
	err := p.InitCodeForceJudge()
	p.Submit = submit
	p.PID = PID
	if err != nil {
		logger.Errorf("Call InitCodeForceJudge failed,err=%s", err.Error())
		return fmt.Errorf("call InitCodeForceJudge failed,err=%s", err.Error())
	}
	err = p.login()
	if err != nil {
		logger.Errorf("Call login failed,err=%s", err.Error())
		return fmt.Errorf("call login failed,err=%s", err.Error())
	}
	if !p.submit() {
		logger.Errorf("Call submit failed,submit=%s", utils.Sdump(submit))
		return fmt.Errorf("call submit failed,submit=%s", utils.Sdump(submit))
	}
	p.getResult()
	p.commitToDB()
	return nil
}

func initUserCount() {
	cfLock.Lock()
	defer cfLock.Unlock()
	if JudgeUsers != nil {
		return
	}
	for i := 1; i <= 20; i++ {
		JudgeUsers = append(JudgeUsers, CFJudgeUser{
			OriginJudgeUser: OriginJudgeUser{
				Status:   JUDGE_FREE,
				Cookies:  make(map[string]string, 0),
				ID:       fmt.Sprintf("AOJjudge%02d", i),
				Password: "AhutAcm@108",
			},
			CsrfToken: "",
		})
	}
}

func (p *CodeForceJudge) initCodeforceHead() {
	p.Headers = CfHeaders
}

func (p *CodeForceJudge) getCsrfToekn() (string, error) {
	logger := utils.GetLogInstance()
	url := "https://codeforces.com"
	resp, err := DoRequest(GET, url, nil, nil, nil, true)
	if err != nil {
		logger.Errorf("call DoRequest failed,url:%s,err=%s", url, err.Error())
		return "", err
	}
	SetCookies(resp, &p.judgeUser.OriginJudgeUser)
	bodyText, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		logger.Errorf("call ReadAll failed,resp=%+v,err=%s", utils.Sdump(resp), err.Error())
		return "", err
	}
	re, _ := regexp.Compile(`<meta name="X-Csrf-Token" content="(?s:([0-9a-f]*))"`)
	ret := re.FindSubmatch(bodyText)
	if ret == nil {
		return "", errors.New("find csrf failed")
	}
	// regexp.find(`<meta name="X-Csrf-Token" content="([0-9a-f]*)"`, *body)
	return string(ret[1]), nil
}

func getRangeUser() (*CFJudgeUser, error) {
	cfLock.Lock()
	defer cfLock.Unlock()
	for idx := range JudgeUsers {
		user := &JudgeUsers[idx]
		if user.Status == JUDGE_FREE {
			user.Status = JUDGE_BUSY
			return user, nil
		}
	}
	return nil, nil
}

func retRangeUser(judgeUser *CFJudgeUser) {
	cfLock.Lock()
	defer cfLock.Unlock()
	judgeUser.Status = JUDGE_FREE
}

// 初始化一个判题机
func (p *CodeForceJudge) InitCodeForceJudge() error {
	// logger := utils.GetLogInstance()
	if JudgeUsers == nil {
		initUserCount()
	}
	// 必须得到一个运行中的判题机
	for {
		p.judgeUser, _ = getRangeUser()
		if p.judgeUser != nil {
			break
		}
		time.Sleep(time.Second)
	}
	p.loginSuccess = false
	p.initCodeforceHead()
	return nil
}

func getFtaa() string {
	str := "0123456789qwertyuiopasdfghjklzxcvbnm"
	var ret = ""
	for i := 0; i < 18; i++ {
		ret += string(str[int(math.Floor(rand.Float64()*float64(len(str))))])
	}
	return ret
}

func (p *CodeForceJudge) checkLoginSuccess() bool {
	logger := utils.GetLogInstance()
	if p.loginSuccess {
		return true
	}
	resp, err := DoRequest(GET, cfurl, p.Headers, p.judgeUser.Cookies, nil, true)
	if err != nil {
		logger.Errorf("call DoRequest failed")
		return false
	}
	Text, err := ParseRespToByte(resp)
	if err != nil {
		logger.Error("call ReadAll failed")
		return false
	}
	ret := strings.Index(string(Text), "logout")
	if ret != -1 {
		p.loginSuccess = true
	}
	return p.loginSuccess
}

func (p *CodeForceJudge) login() error {
	logger := utils.GetLogInstance()
	url := "https://codeforces.com/enter?locale=en"
	userCount, _ := getRangeUser()
	p.judgeUser = userCount
	SetCookies(nil, &p.judgeUser.OriginJudgeUser)
	if p.checkLoginSuccess() {
		return nil
	}
	/*没有登录的情况下  需要重新做一次登录*/
	p.judgeUser.CsrfToken, _ = p.getCsrfToekn()
	ftaa := getFtaa()
	bfaa := "f1b3f18c715565b589b7823cda7448ce"
	var data = fmt.Sprintf("csrf_token=%s&action=enter&handleOrEmail=%s&password=%s&remember=on&ftaa=%s&bfaa=%s&_taa=176", userCount.CsrfToken, userCount.ID, userCount.ID, ftaa, bfaa)
	resp, err := DoRequest(POST, url, p.Headers, p.judgeUser.Cookies, &data, false)
	if err != nil {
		logger.Error(err.Error())
	}
	//填充 cookie
	SetCookies(resp, &p.judgeUser.OriginJudgeUser)
	if p.checkLoginSuccess() {
		return nil
	}
	return errors.New("login failed")
}

func (p *CodeForceJudge) GetCFLang() string {
	return cflang[p.Submit.Lang]
}

func (p *CodeForceJudge) ParsePID() (string, string, error) {
	if len(p.PID) < 1 {
		return "", "", errors.New("")
	}
	re, _ := regexp.Compile("([0-9]*)([A-Z]*)")
	strs := re.FindStringSubmatch(p.PID)
	if len(strs) < 3 {
		return "", "", fmt.Errorf("ParsePID failed ,pid:%s", p.PID)
	}
	return strs[1], strs[2], nil
}

func checkCFSubmitResp(resp *http.Response) bool {
	Text, err := ParseRespToByte(resp)
	if err != nil {
		return false
	}
	re, _ := regexp.Compile(`submissionid="([0-9]*)"`)
	ret := re.FindSubmatch(Text)
	return ret != nil
}

func GetContest(CID string) string {
	CIDInt, _ := strconv.Atoi(CID)
	if CIDInt >= 100000 {
		return "gym"
	}
	return "contest"
}

func (p *CodeForceJudge) submit() bool {
	logger := utils.GetLogInstance()
	CID, idx, _ := p.ParsePID()
	url := cfurl + `/` + GetContest(CID) + `/` + CID + `/submit?csrf_token=` + p.judgeUser.CsrfToken
	lang := p.GetCFLang()
	var dataMap = map[string]string{
		"csrf_token":            p.judgeUser.CsrfToken,
		"action":                "submitSolutionFormSubmitted",
		"ftaa":                  getFtaa(),
		"bfaa":                  "f1b3f18c715565b589b7823cda7448ce",
		"submittedProblemIndex": idx,
		"programTypeId":         lang,
		"contestId":             CID,
		"source":                p.Submit.Source,
		"sourceFile":            "",
		"tabSize":               "4",
		"_tta":                  "493",
		"sourceCodeConfirmed":   "true",
	}
	data := MapToStrings(dataMap, "&")
	resp, err := DoRequest(POST, url, p.Headers, p.judgeUser.Cookies, &data, true)
	if err != nil {
		logger.Errorf("Call DoRequest failed,err=%s", err.Error())
		return false
	}
	return checkCFSubmitResp(resp)
}

func (p *CodeForceJudge) GetSubmitID() (string, error) {
	CID, _, _ := p.ParsePID()
	url := cfurl + "/" + GetContest(CID) + "/" + CID + "/my"
	resp, err := DoRequest(GET, url, p.Headers, p.judgeUser.Cookies, nil, false)
	if err != nil {
		logger.Errorf("call DoRequest failed,url:%s, err=%s", url, err.Error())
		return "", err
	}
	Text, err := ParseRespToByte(resp)
	if err != nil {
		return "", err
	}
	re, _ := regexp.Compile(`submissionid="([0-9]*)"`)
	ret := re.FindSubmatch(Text)
	if ret == nil {
		return "", errors.New("not find submission")
	}
	return string(ret[1]), nil
}

func CheckResult(result string) constanct.OJResult {
	for cfStatus, WStatus := range CFResultMap {
		re := regexp.MustCompile(string(cfStatus))
		ret := re.FindString(result)
		if ret != "" {
			return WStatus
		}
	}
	return constanct.OJ_JUDGE
}

func (p *CodeForceJudge) getResult() error {
	logger := utils.GetLogInstance()
	// https://codeforces.com/contest/1003/submission/174882990
	CID, _, err := p.ParsePID()
	if err != nil {
		logger.Errorf("call ParsePID failed,PID:%s, err=%s", p.PID, err.Error())
		return err
	}
	submissionID, err := p.GetSubmitID()
	if err != nil {
		logger.Errorf("call GetSubmitID failed")
		return err
	}
	retRangeUser(p.judgeUser)
	url := cfurl + "/" + GetContest(CID) + "/" + CID + "/submission/" + submissionID
	for {
		resp, err := DoRequest(GET, url, p.Headers, nil, nil, false)
		if err != nil {
			return err
		}
		Reader := ParseRespToReader(resp)
		doc, err := goquery.NewDocumentFromReader(Reader)
		if err != nil {
			return err
		}
		doc.Find("table").Children().Find("tr").Next().Find("td").Each(func(i int, s *goquery.Selection) {
			/*判题结果*/
			if i == 4 {
				result := DealStrings(s.Text())
				res := CheckResult(result)
				if res == constanct.OJ_JUDGE {
					return
				}
				p.Submit.Result = res
			}
			if p.Submit.Result == constanct.OJ_JUDGE {
				return
			}
			/*时间ms*/
			if i == 5 {
				time := DealStrings(s.Text())
				re, _ := regexp.Compile(`([0-9]*) ms`)
				p.Submit.Usetime, _ = strconv.ParseInt(re.FindStringSubmatch(time)[1], 10, 64)
			}
			/*内存KB*/
			if i == 6 {
				Memory := DealStrings(s.Text())
				re, _ := regexp.Compile(`([0-9]*) KB`)
				p.Submit.UseMemory, _ = strconv.ParseInt(re.FindStringSubmatch(Memory)[1], 10, 64)
			}
		})
		if p.Submit.Result != constanct.OJ_JUDGE {
			return nil
		}
		time.Sleep(time.Second)
	}
	// return nil
}

func (p *CodeForceJudge) commitToDB() error {
	return models.UpdateSubmit(context.Background(), p.Submit)
}
