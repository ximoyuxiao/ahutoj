package originJudged

import (
	"ahutoj/web/dao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/utils"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
)

type CFJudgeUser struct {
	OriginJudgeUser
	CsrfToken string
}

var logger *logrus.Logger = utils.GetLogInstance()
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
	"Accepted":                   constanct.OJ_AC,
	"Happy New Year!":            constanct.OJ_AC,
	"Compilation error(.*?)":     constanct.OJ_CE,
	"Memory limit exceeded(.*?)": constanct.OJ_MLE,
	"O(.*?)":                     constanct.OJ_OLE,
	"P(.*?)":                     constanct.OJ_PE,
	"Runtime error(.*?)":         constanct.OJ_RE,
	"Time limit exceeded(.*?)":   constanct.OJ_TLE,
	"Wrong answer(.*?)":          constanct.OJ_WA,
	"Running":                    constanct.OJ_JUDGE,
	"Inqueue(.*?)":               constanct.OJ_JUDGE,
	"Denial of judgement":        constanct.OJ_DENIAL,
	"In queue":                   constanct.OJ_JUDGE,
	"Running on test (.*?)":      constanct.OJ_JUDGE,
	"Pretests passed":            constanct.OJ_AC,
	"Pendding":                   constanct.OJ_JUDGE,
}

var CfHeaders = map[string]string{
	"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:105.0) Gecko/20100101 Firefox/105.0",
	"Accept":     "*/*",
	// "Accept-Encoding": "gzip, deflate, br",
	"Origin":       "https://codeforces.com",
	"Connection":   "keep-alive",
	"Content-Type": "application/x-www-form-urlencoded",
	"Host":         "www.codeforces.com",
}

type CodeForceJudge struct {
	Headers      map[string]string
	JudgeUser    *CFJudgeUser
	loginSuccess bool
	OriginJudge
}

func (p CodeForceJudge) Judge(ctx context.Context, submit dao.Submit, PID string) error {
	err := p.InitCodeForceJudge()
	p.PID = PID
	if err != nil {
		logger.Errorf("Call InitCodeForceJudge failed,err=%v", err.Error())
		return fmt.Errorf("call InitCodeForceJudge failed,err=%v", err.Error())
	}
	defer p.retRangeUser()
	defer p.CommitResult()
	p.Submit = submit
	err = p.Login()
	if err != nil {
		logger.Errorf("Codeforces Call login failed,err=%v, JdugedID:%v", err.Error(), p.JudgeUser.OriginJudgeUser.ID)
		return fmt.Errorf("call login failed,err=%v", err.Error())
	}
	p.SetJudeID(ctx)
	if !p.submit() {
		logger.Errorf("Call submit failed,submit=%v, JudgeID:%v", submit.SID, p.JudgeUser.OriginJudgeUser.ID)
		return fmt.Errorf("call submit failed,submit=%v", submit.SID)
	}
	err = p.getResult()
	if err != nil {
		logger.Errorf("Call getResult failed,submit:%v, err:%v", submit.SID, err.Error())
		return err
	}
	fmt.Println("判题结束:结果为" + utils.Sdump(p.Submit))
	return nil
}

func initUserCount() {
	cfLock.Lock()
	defer cfLock.Unlock()
	if JudgeUsers != nil {
		return
	}
	config := utils.GetConfInstance().CodeForceJudge
	for i := 1; i <= int(config.Count); i++ {
		JudgeUsers = append(JudgeUsers, CFJudgeUser{
			OriginJudgeUser: OriginJudgeUser{
				Status:   JUDGE_FREE,
				Cookies:  make(map[string]string, 0),
				ID:       fmt.Sprintf("%v%02d", config.Prefix, i),
				Password: config.Password,
			},
			CsrfToken: "",
		})
	}
}

func (p *CodeForceJudge) initCodeforceHead() {
	p.Headers = CfHeaders
}

func (p *CodeForceJudge) getCsrfToekn() (string, error) {
	url := "https://codeforces.com"
	p.JudgeUser.Cookies["RCPC"] = "cb8d001c7d179c2536d275752653bc56"
	resp, err := utils.DoRequest(utils.GET, url, p.Headers, p.JudgeUser.Cookies, nil, true)
	if err != nil {
		logger.Errorf("call DoRequest failed,url:%v,err=%v", url, err.Error())
		return "", err
	}
	SetCookies(resp, &p.JudgeUser.OriginJudgeUser)
	bodyText, err := utils.ParseRespToByte(resp)
	defer resp.Body.Close()
	if err != nil {
		logger.Errorf("call ReadAll failed,resp=%+v,err=%v", utils.Sdump(resp), err.Error())
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
	pos := rand.Int() % len(JudgeUsers)
	for idx := range JudgeUsers {
		user := &JudgeUsers[(pos+idx)%len(JudgeUsers)]
		if user.Status == JUDGE_FREE {
			user.Status = JUDGE_BUSY
			return user, nil
		}
	}
	return nil, fmt.Errorf("not find free judgedUser")
}

func (p *CodeForceJudge) SetJudeID(ctx context.Context) {
	for idx, judge := range JudgeUsers {
		if judge.ID == p.JudgeUser.ID {
			p.Submit.JudgeID = int64(idx + 1)
		}
	}
}
func (p *CodeForceJudge) retRangeUser() {
	if p.JudgeUser == nil {
		return
	}
	cfLock.Lock()
	defer cfLock.Unlock()
	p.JudgeUser.Status = JUDGE_FREE
	p.JudgeUser = nil
}

// 初始化一个判题机
func (p *CodeForceJudge) InitCodeForceJudge() error {
	// ()
	if JudgeUsers == nil {
		initUserCount()
	}
	// 必须得到一个运行中的判题机
	for {
		p.JudgeUser, _ = getRangeUser()
		if p.JudgeUser != nil {
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
	if p.loginSuccess {
		return true
	}
	resp, err := utils.DoRequest(utils.GET, cfurl, p.Headers, p.JudgeUser.Cookies, nil, true)
	if err != nil {
		logger.Errorf("call DoRequest failed")
		return false
	}
	Text, err := utils.ParseRespToByte(resp)
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

func (p *CodeForceJudge) Login() error {

	url := "https://codeforces.com/enter?locale=en"
	if p.JudgeUser == nil {
		p.JudgeUser, _ = getRangeUser()
	}
	userCount := p.JudgeUser
	logger.Debugf("use user:%+v:", utils.Sdump(userCount))
	if p.checkLoginSuccess() {
		return nil
	}
	SetCookies(nil, &p.JudgeUser.OriginJudgeUser)
	/*没有登录的情况下  需要重新做一次登录*/
	p.JudgeUser.CsrfToken, _ = p.getCsrfToekn()
	ftaa := getFtaa()
	bfaa := "f1b3f18c715565b589b7823cda7448ce"
	var data = fmt.Sprintf("csrf_token=%v&action=enter&handleOrEmail=%v&password=%v&remember=on&ftaa=%v&bfaa=%v&_taa=176", userCount.CsrfToken, userCount.ID, userCount.Password, ftaa, bfaa)
	resp, err := utils.DoRequest(utils.POST, url, p.Headers, p.JudgeUser.Cookies, &data, false)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	//填充 cookie
	SetCookies(resp, &p.JudgeUser.OriginJudgeUser)
	if p.checkLoginSuccess() {
		return nil
	}
	logger.Errorf("login failed,data=%v", data)
	return errors.New("login failed")
}

func (p *CodeForceJudge) GetCFLang() string {
	return cflang[p.Submit.Lang]
}

func (p *CodeForceJudge) ParsePID() (string, string, error) {
	if len(p.PID) < 1 {
		return "", "", errors.New("")
	}
	re, _ := regexp.Compile("([0-9]*)([A-Z]*[0-9]*)")
	strs := re.FindStringSubmatch(p.PID)
	if len(strs) < 3 {
		return "", "", fmt.Errorf("ParsePID failed ,pid:%v", p.PID)
	}
	return strs[1], strs[2], nil
}

func checkCFSubmitResp(resp *http.Response, CID string) bool {
	if resp.StatusCode != 302 {
		logger.Errorf("submit failed,should Status code 302 but now statuscode:%v", resp.StatusCode)
		return false
	}
	Nexturl := "https://codeforces.com/" + GetContest(CID) + "/" + CID + "/my"
	Location := resp.Header.Get("Location")
	if Nexturl != Location {
		logger.Errorf("submit failed,execpt NextURL:%v but now URL:%v", Nexturl, Location)
		return false
	}
	return true
}

func GetContest(CID string) string {
	CIDInt, _ := strconv.Atoi(CID)
	if CIDInt >= 100000 {
		return "gym"
	}
	return "contest"
}

func (p *CodeForceJudge) submit() bool {
	CID, idx, _ := p.ParsePID()
	url := cfurl + `/` + GetContest(CID) + `/` + CID + `/submit?csrf_token=` + p.JudgeUser.CsrfToken
	lang := p.GetCFLang()
	var dataMap = map[string]string{
		"csrf_token":            p.JudgeUser.CsrfToken,
		"action":                "submitSolutionFormSubmitted",
		"ftaa":                  getFtaa(),
		"bfaa":                  "f1b3f18c715565b589b7823cda7448ce",
		"submittedProblemIndex": idx,
		"programTypeId":         lang,
		"contestId":             CID,
		"source":                fmt.Sprintf("%v//submitTime:%v\n", p.Submit.Source, time.Now().UnixMilli()),
		"sourceFile":            "",
		"tabSize":               "4",
		"_tta":                  "493",
		"sourceCodeConfirmed":   "true",
	}
	data := utils.MapToFormStrings(dataMap, "&")
	resp, err := utils.DoRequest(utils.POST, url, p.Headers, p.JudgeUser.Cookies, &data, false)
	if err != nil {
		logger.Errorf("Call DoRequest failed,err=%v, data:%v", err.Error(), data)
		return false
	}
	return checkCFSubmitResp(resp, CID)
}

func (p *CodeForceJudge) GetSubmitID() (string, error) {
	CID, _, _ := p.ParsePID()
	url := cfurl + "/" + GetContest(CID) + "/" + CID + "/my"
	resp, err := utils.DoRequest(utils.GET, url, p.Headers, p.JudgeUser.Cookies, nil, true)
	if err != nil {
		logger.Errorf("call DoRequest failed,url:%v, err=%v", url, err.Error())
		return "", err
	}
	Text, err := utils.ParseRespToByte(resp)
	if err != nil {
		return "", err
	}
	re, _ := regexp.Compile(`submissionId="(.*?)"`)
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
	return constanct.OJ_TIMEOUT
}

type ResultObj struct {
	Verdict                string `json:"verdict"`
	Href                   string `json:"href"`
	CheckerStdoutAndStderr string `json:"checkerStdoutAndStderr#1"`
}

func (p *CodeForceJudge) getCEinfo(submissionID string) error {
	ResultURL := "https://codeforces.com/data/submitSource"
	var resp *http.Response
	csrftoken := ""
	if p.JudgeUser != nil {
		csrftoken = p.JudgeUser.CsrfToken
	}
	DataMap := map[string]string{
		"submissionId": submissionID,
		"csrf_token":   csrftoken,
	}
	data := utils.MapToFormStrings(DataMap, "&")
	resp, err := utils.DoRequest(utils.POST, ResultURL, p.Headers, p.JudgeUser.Cookies, &data, false)
	if err != nil {
		logger.Errorf("call DoRequest(getResult2) failed, err=%v", err.Error())
	}
	ret, err := utils.ParseRespToByte(resp)
	if err != nil {
		logger.Errorf("call ParseRespToByte(getResult2) failed, err=%v", err.Error())
	}
	var result = new(ResultObj)
	json.Unmarshal(ret, result)
	if result.CheckerStdoutAndStderr != "" {
		p.CEInfo = result.CheckerStdoutAndStderr
	}
	fmt.Println(result)
	return nil
}

func (p *CodeForceJudge) getResult() error {
	// https://codeforces.com/contest/1003/submission/174882990
	CID, _, err := p.ParsePID()
	if err != nil {
		logger.Errorf("call ParsePID failed,PID:%v, err=%v", p.PID, err.Error())
		return err
	}
	submissionID, err := p.GetSubmitID()
	if err != nil {
		logger.Errorf("call GetSubmitID failed")
		return err
	}
	p.Submit.JudgeID, _ = strconv.ParseInt(fmt.Sprintf("%v%v", p.Submit.JudgeID, submissionID), 10, 64)
	url := cfurl + "/" + GetContest(CID) + "/" + CID + "/submission/" + submissionID
	var resp *http.Response
	// 死循环去获取 最后肯定有一个结果的
	for {
		if p.JudgeUser != nil {
			resp, err = utils.DoRequest(utils.GET, url, p.Headers, p.JudgeUser.Cookies, nil, false)
		} else {
			resp, err = utils.DoRequest(utils.GET, url, p.Headers, nil, nil, false)
		}
		if err != nil {
			return err
		}
		Reader := utils.ParseRespToReader(resp)
		doc, err := goquery.NewDocumentFromReader(Reader)
		if err != nil {
			return err
		}
		doc.Find("table").Children().Find("tr").Next().Find("td").Each(func(i int, s *goquery.Selection) {
			/*判题结果*/
			if i == 4 {
				result := utils.DealStrings(s.Text())
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
				time := utils.DealStrings(s.Text())
				re, _ := regexp.Compile(`([0-9]*) ms`)
				p.Submit.Usetime, _ = strconv.ParseInt(re.FindStringSubmatch(time)[1], 10, 64)
			}
			/*内存KB*/
			if i == 6 {
				Memory := utils.DealStrings(s.Text())
				re, _ := regexp.Compile(`([0-9]*) KB`)
				p.Submit.UseMemory, _ = strconv.ParseInt(re.FindStringSubmatch(Memory)[1], 10, 64)
				p.Submit.UseMemory *= 1024
			}
		})
		if p.Submit.Result != constanct.OJ_JUDGE {
			break
		}
		time.Sleep(time.Second)
	}
	if p.Submit.Result == constanct.OJ_CE {
		err := p.getCEinfo(submissionID)
		if err != nil {
			logger.Errorf("call getCEinfo failed,submissionID:%v, err=%v", submissionID, err.Error())
			return err
		}
		return nil
	}
	if p.Submit.Result == constanct.OJ_JUDGE {
		p.Submit.Result = constanct.OJ_TIMEOUT // 判题超时
		return fmt.Errorf("codeforeces judge timeout submissionID:%v", submissionID)
	}
	return nil
}
