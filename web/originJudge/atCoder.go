package originjudge

import (
	"ahutoj/web/dao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/models"
	"ahutoj/web/utils"
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/bytedance/gopkg/util/logger"
)

type ATcoderJudgeUser struct {
	OriginJudgeUser
	CsrfToken string
}

var atlock sync.Mutex

type ATcoderJudgeUsers []ATcoderJudgeUser

var atcoderJudgeUsers ATcoderJudgeUsers = nil

const ATcoderoj OJPlatform = 2

var atcoderUrl = "https://atcoder.jp"

var atcoderLang = map[constanct.LANG]string{
	constanct.C:       "43",
	constanct.CPP:     "50",
	constanct.CPP11:   "50",
	constanct.CPP17:   "54",
	constanct.JAVA:    "36",
	constanct.PYTHON3: "31",
}

var atcoderResultMap = map[string]constanct.OJResult{
	"AC":                       constanct.OJ_AC,
	"Compilationerror(.*?)":    constanct.OJ_CE,
	"Memorylimitexceeded(.*?)": constanct.OJ_MLE,
	"O(.*?)":                   constanct.OJ_OLE,
	"P(.*?)":                   constanct.OJ_PE,
	"Runtimeerror(.*?)":        constanct.OJ_RE,
	"TLE":                      constanct.OJ_TLE,
	"Wronganswer(.*?)":         constanct.OJ_WA,
	"Running":                  constanct.OJ_JUDGE,
	"Inqueue(.*?)":             constanct.OJ_JUDGE,
}

var atcoderHeaders = map[string]string{
	"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:105.0) Gecko/20100101 Firefox/105.0",
	"Accept":          "*/*",
	"Accept-Encoding": "gzip, deflate, br",
	"Origin":          "https://atcoder.jp",
	"Connection":      "keep-alive",
	"Content-Type":    "application/x-www-form-urlencoded",
	"Host":            "atcoder.jp",
}

type AtCoderJudge struct {
	Headers      map[string]string
	JudgeUser    *ATcoderJudgeUser
	LoginSuccess bool
	SubmissionID string
	OriginJudge
}

func (p AtCoderJudge) Judge(ctx context.Context, submit dao.Submit, PID string) error {
	err := p.InitAtcoderJudge(ctx)
	p.Submit = submit
	p.PID = PID
	if err != nil {
		logger.Errorf("Call InitAtcoderJudge failed,err=%s", err.Error())
		return fmt.Errorf("call InitAtcoderJudge failed,err=%s", err.Error())
	}
	err = p.login(ctx)
	if err != nil {
		logger.Errorf("Call login failed,err=%s", err.Error())
		return fmt.Errorf("call login failed,err=%s", err.Error())
	}
	if err = p.submit(ctx); err != nil {
		logger.Errorf("Call submit failed,submit=%s", utils.Sdump(submit))
		return fmt.Errorf("call submit failed,submit=%s", utils.Sdump(submit))
	}
	p.getResult(ctx)
	p.commitToDB(ctx)
	return nil
}

func getAtcoderUser(ctx context.Context) *ATcoderJudgeUser {
	return &ATcoderJudgeUser{
		OriginJudgeUser: OriginJudgeUser{
			Status:   JUDGE_BUSY,
			Cookies:  make(map[string]string, 0),
			ID:       "XMYchen",
			Password: "chen199094212",
		},
		CsrfToken: "",
	}
}
func initAtcoderUserCount(ctx context.Context) {
	atlock.Lock()
	defer atlock.Unlock()
	if atcoderJudgeUsers != nil {
		return
	}
	for i := 1; i <= 20; i++ {
		atcoderJudgeUsers = append(atcoderJudgeUsers, ATcoderJudgeUser{
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
func (p *AtCoderJudge) InitAtcoderJudge(ctx context.Context) error {
	// logger := utils.GetLogInstance()
	if atcoderJudgeUsers == nil {
		initAtcoderUserCount(ctx)
	}
	//必须得到一个运行中的判题机
	for {
		p.JudgeUser = getAtcoderUser(ctx)
		if p.JudgeUser != nil {
			break
		}
		time.Sleep(time.Second)
	}
	p.LoginSuccess = false
	p.Headers = atcoderHeaders
	return nil
}
func (p *AtCoderJudge) checkAtcoderLogin(ctx context.Context) bool {
	if p.LoginSuccess {
		return true
	}
	resp, err := DoRequest(GET, atcoderUrl, p.Headers, p.JudgeUser.Cookies, nil, false)
	if err != nil {
		return false
	}
	Text, err := ParseRespToByte(resp)
	if err != nil {
		return false
	}
	idx := strings.Index(string(Text), "Sign Out")
	if idx != -1 {
		p.LoginSuccess = true
	}
	return p.LoginSuccess
}
func (p *AtCoderJudge) getCsrfToekn() (string, error) {
	resp, err := DoRequest(GET, atcoderUrl, p.Headers, p.JudgeUser.Cookies, nil, false)
	if err != nil {
		return "", err
	}
	SetCookies(resp, &p.JudgeUser.OriginJudgeUser)
	Text, _ := ParseRespToByte(resp)
	re := regexp.MustCompile(`var csrfToken = "(.*?)"`)
	ret := re.FindSubmatch(Text)
	if ret == nil {
		return "", errors.New("find Token failed")
	}
	return string(ret[1]), nil
}
func (p *AtCoderJudge) login(ctx context.Context) error {
	// logger := utils.GetLogInstance()
	userCount := getAtcoderUser(ctx)
	p.JudgeUser = userCount
	if p.checkAtcoderLogin(ctx) {
		return nil
	}
	/*没有登录信息  登录*/
	loginURL := "https://atcoder.jp/login"
	for i := 1; i < 10; i++ {
		p.JudgeUser.CsrfToken, _ = p.getCsrfToekn()
		var data = map[string]string{
			"username":   p.JudgeUser.ID,
			"password":   p.JudgeUser.Password,
			"csrf_token": p.JudgeUser.CsrfToken,
		}
		formdata := MapToStrings(data, "&")
		resp, err := DoRequest(POST, loginURL, p.Headers, p.JudgeUser.Cookies, &formdata, false)
		if err != nil {
			return err
		}
		SetCookies(resp, &p.JudgeUser.OriginJudgeUser)
		if p.checkAtcoderLogin(ctx) {
			return nil
		}
		p.JudgeUser.OriginJudgeUser.Cookies = make(map[string]string, 0)
	}
	return fmt.Errorf("try login failed")
}
func (p *AtCoderJudge) ParsePID(ctx context.Context) (string, string, error) {
	return "", "", nil
}

func (p *AtCoderJudge) CheckAndGetSubmissionID(ctx context.Context, resp *http.Response) string {
	if resp.StatusCode != 302 {
		return ""
	}
	CID, _, _ := p.ParsePID(ctx)
	redyLocal := "/contests/" + CID + "/submissions/me"
	local := resp.Header.Get("Location")
	if local != redyLocal {
		return ""
	}
	checkUrl := atcoderUrl + redyLocal
	resp, err := DoRequest(GET, checkUrl, p.Headers, p.JudgeUser.Cookies, nil, false)
	if err != nil {
		return ""
	}
	SetCookies(resp, &p.JudgeUser.OriginJudgeUser)
	Text, err := ParseRespToByte(resp)
	if err != nil {
		return ""
	}
	re := regexp.MustCompile(`<a href="/contests/.*?/submissions/([0-9]*)">Detail</a>`)
	ret := re.FindSubmatch(Text)
	if ret == nil {
		return ""
	}
	return string(ret[1])
}

func (p *AtCoderJudge) submit(ctx context.Context) error {
	CID, _, _ := p.ParsePID(ctx)
	submitUrl := atcoderUrl + "/contests/" + CID + "/submit"
	var dataMap = map[string]string{
		"data.TaskScreenName": p.PID,
		"data.LanguageId":     atcoderLang[p.Submit.Lang],
		"sourceCode":          p.Submit.Source,
		"csrf_token":          p.JudgeUser.CsrfToken,
	}
	data := MapToStrings(dataMap, "&")
	resp, err := DoRequest(POST, submitUrl, p.Headers, p.JudgeUser.Cookies, &data, false)
	if err != nil {
		return err
	}
	SetCookies(resp, &p.JudgeUser.OriginJudgeUser)
	p.SubmissionID = p.CheckAndGetSubmissionID(ctx, resp)
	if p.SubmissionID == "" {
		return errors.New("submit SourceCode failed")
	}
	p.RetJudgeUser(ctx)
	return nil
}
func (p *AtCoderJudge) RetJudgeUser(ctx context.Context) {
	atlock.Lock()
	defer atlock.Unlock()
	p.JudgeUser.Status = JUDGE_FREE
	p.JudgeUser = nil
}
func (p *AtCoderJudge) CheckResult(ctx context.Context, Text []byte) bool {
	re := regexp.MustCompile(`<span class='label label-warning' data-toggle='tooltip' data-placement='top' title=".*?">([A-Za-z]*)</span>`)
	ret := re.FindSubmatch(Text)
	if ret == nil {
		return false
	}
	result := string(ret[1])
	if atcoderResultMap[result] == constanct.OJ_JUDGE {
		return false
	}
	p.Submit.Result = atcoderResultMap[result]
	return true
}
func (p *AtCoderJudge) getResult(ctx context.Context) error {
	CID, _, _ := p.ParsePID(ctx)
	submissionUrl := atcoderUrl + "/contests/" + CID + "/submissions/" + p.SubmissionID
	for {
		resp, err := DoRequest(GET, submissionUrl, nil, nil, nil, false)
		if err != nil {
			return err
		}
		Text, _ := ParseRespToByte(resp)
		if !p.CheckResult(ctx, Text) {
			time.Sleep(time.Second)
			continue
		}
		/*解析内存和*/
		re := regexp.MustCompile(`<td class="text-center">(.*?)</td>`)
		ret := re.FindAllSubmatch(Text, 9)
		timere := regexp.MustCompile(`[0-9]*`)
		time := timere.FindString(string(ret[6][1]))
		p.Submit.Usetime, _ = strconv.ParseInt(time, 10, 64)
		memory := timere.FindString(string(ret[7][1]))
		p.Submit.UseMemory, _ = strconv.ParseInt(memory, 10, 64)
		return nil
	}
}
func (p *AtCoderJudge) commitToDB(ctx context.Context) error {
	return models.UpdateSubmit(ctx, p.Submit)
}
