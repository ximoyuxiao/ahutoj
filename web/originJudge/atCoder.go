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
)

type ATcoderJudgeUser struct {
	OriginJudgeUser
	LastSubmitTime int64
	SleepTime      int64
	CsrfToken      string
}

var atlock sync.Mutex

type ATcoderJudgeUsers []ATcoderJudgeUser

var atcoderJudgeUsers ATcoderJudgeUsers = nil

const ATcoderoj OJPlatform = 2

var atcoderUrl = "https://atcoder.jp"

var atcoderLang = map[constanct.LANG]string{
	constanct.C:       "4002",
	constanct.CPP:     "4004",
	constanct.CPP11:   "4004",
	constanct.CPP17:   "4004",
	constanct.JAVA:    "4005",
	constanct.PYTHON3: "4006",
}

var atcoderResultMap = map[string]constanct.OJResult{
	"AC":           constanct.OJ_AC,
	"CE":           constanct.OJ_CE,
	"MLE":          constanct.OJ_MLE,
	"OLE":          constanct.OJ_OLE,
	"PE":           constanct.OJ_PE,
	"RE":           constanct.OJ_RE,
	"TLE":          constanct.OJ_TLE,
	"WA":           constanct.OJ_WA,
	"WJ":           constanct.OJ_JUDGE,
	"Inqueue(.*?)": constanct.OJ_JUDGE,
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
	logger := utils.GetLogInstance()
	logger.Infof("use atcoder judgeing:SID:%v", submit.SID)
	err := p.InitAtcoderJudge(ctx)
	if err != nil {
		logger.Errorf("Call InitAtcoderJudge failed,err=%v", err.Error())
		return fmt.Errorf("call InitAtcoderJudge failed,err=%v", err.Error())
	}
	defer p.RetJudgeUser(ctx)
	defer p.commitToDB(ctx)
	defer logger.Infof("judge complete judgeing:SID:%v result:%v", p.Submit.SID, submit.Result)
	p.Submit = submit
	p.PID = PID
	err = p.login(ctx)
	if err != nil {
		logger.Errorf("Call login failed,err=%v", err.Error())
		return fmt.Errorf("call login failed,err=%v", err.Error())
	}
	p.SetJudeID(ctx)
	if err = p.submit(ctx); err != nil {
		logger.Errorf("Call submit failed,submit=%v", submit.PID)
		return fmt.Errorf("call submit failed,submit=%v", submit.PID)
	}
	p.getResult(ctx)
	return nil
}

func (p *AtCoderJudge) SetJudeID(ctx context.Context) {
	for idx, judge := range atcoderJudgeUsers {
		if judge.ID == p.JudgeUser.ID {
			p.Submit.JudgeID = int64(idx + 1)
		}
	}
}

func getAtcoderUser(ctx context.Context) *ATcoderJudgeUser {
	atlock.Lock()
	defer atlock.Unlock()
	for idx := range atcoderJudgeUsers {
		user := &atcoderJudgeUsers[idx]
		if user.Status == JUDGE_FREE && time.Now().UnixNano()-user.LastSubmitTime >= user.SleepTime {
			if user.SleepTime == 10*int64(time.Second) {
				user.SleepTime = 60 * int64(time.Second)
			}
			user.Status = JUDGE_BUSY
			return user
		}
	}
	return nil
}

func initAtcoderUserCount(ctx context.Context) {
	atlock.Lock()
	defer atlock.Unlock()
	if atcoderJudgeUsers != nil {
		return
	}
	config := utils.GetConfInstance().AtCoderJudges
	for i := 1; i <= int(config.Count); i++ {
		atcoderJudgeUsers = append(atcoderJudgeUsers, ATcoderJudgeUser{
			OriginJudgeUser: OriginJudgeUser{
				Status:   JUDGE_FREE,
				Cookies:  make(map[string]string, 0),
				ID:       fmt.Sprintf("%v%02d", config.Prefix, i),
				Password: config.Password,
			},
			SleepTime: 10 * int64(time.Second),
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

func (p *AtCoderJudge) checkAtcoderLogin(ctx context.Context) (bool, error) {
	// logger := utils.GetLogInstance()
	if p.LoginSuccess {
		return true, nil
	}
	resp, err := DoRequest(GET, atcoderUrl, p.Headers, p.JudgeUser.Cookies, nil, false)
	if err != nil {
		return false, fmt.Errorf("call DoRequest failed,err=%v", err.Error())
	}
	SetCookies(resp, &p.JudgeUser.OriginJudgeUser)
	Text, err := ParseRespToByte(resp)
	if err != nil {
		return false, fmt.Errorf("call ParseRespToByte failed,err=%v", err.Error())
	}
	idx := strings.Index(string(Text), "Sign Out")
	if idx != -1 {
		p.LoginSuccess = true
	}
	return p.LoginSuccess, nil
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
	logger := utils.GetLogInstance()
	/*没有登录信息  登录*/
	loginURL := "https://atcoder.jp/login"
	for i := 1; i < 10; i++ {
		for p.JudgeUser == nil {
			p.JudgeUser = getAtcoderUser(ctx)
		}
		ok, err := p.checkAtcoderLogin(ctx)
		if ok {
			return nil
		}
		if err != nil {
			logger.Debugf("call checkAtcoderLogin failed,err=%v", err.Error())
		}
		p.JudgeUser.CsrfToken, _ = p.getCsrfToekn()
		var data = map[string]string{
			"username":   p.JudgeUser.ID,
			"password":   p.JudgeUser.Password,
			"csrf_token": p.JudgeUser.CsrfToken,
		}
		formdata := MapToFormStrings(data, "&")
		resp, err := DoRequest(POST, loginURL, p.Headers, p.JudgeUser.Cookies, &formdata, false)
		if err != nil {
			return err
		}
		SetCookies(resp, &p.JudgeUser.OriginJudgeUser)
		ok, err = p.checkAtcoderLogin(ctx)
		if ok {
			return nil
		}
		if err != nil {
			logger.Debugf("call checkAtcoderLogin failed,err=%v", err.Error())
		}
		p.RetJudgeUser(ctx)
	}
	return fmt.Errorf("try login failed")
}

func (p *AtCoderJudge) ParsePID(ctx context.Context) (string, string, error) {
	logger := utils.GetLogInstance()
	re := regexp.MustCompile(`([A-Za-z0-9]*)_([A-Za-z]*)`)
	ret := re.FindStringSubmatch(p.PID)
	if ret == nil {
		logger.Errorf("problem not found!")
	}
	return ret[1], ret[2], nil
}

func (p *AtCoderJudge) CheckAndGetSubmissionID(ctx context.Context, resp *http.Response) (string, error) {
	if resp.StatusCode != 302 {
		Text, _ := ParseRespToByte(resp)
		re := regexp.MustCompile("<div class=\"alert alert-danger alert-dismissible col-sm-12 fade in\" role=\"alert\" >\n\t\t\t\t<button type=\"button\" class=\"close\" data-dismiss=\"alert\" aria-label=\"Close\"><span aria-hidden=\"true\">&times;</span></button>\n\t\t\t\t<span class=\"glyphicon glyphicon-exclamation-sign\" aria-hidden=\"true\"></span>(.*?)\n\t\t\t</div>")
		ret := re.FindSubmatch(Text)
		return "", fmt.Errorf("resp should 302 but statusCode:%v errInfo:%v", resp.StatusCode, string(ret[1]))
	}
	CID, _, _ := p.ParsePID(ctx)
	redyLocal := "/contests/" + CID + "/submissions/me"
	local := resp.Header.Get("Location")
	if local != redyLocal {
		return "", fmt.Errorf("resp Location should:%v,but Location:%v", redyLocal, local)
	}
	checkUrl := atcoderUrl + redyLocal
	resp, err := DoRequest(GET, checkUrl, p.Headers, p.JudgeUser.Cookies, nil, false)
	if err != nil {
		return "", err
	}
	SetCookies(resp, &p.JudgeUser.OriginJudgeUser)
	Text, err := ParseRespToByte(resp)
	if err != nil {
		return "", err
	}
	re := regexp.MustCompile(`<a href="/contests/.*?/submissions/([0-9]*)">Detail</a>`)
	ret := re.FindSubmatch(Text)
	if ret == nil {
		return "", fmt.Errorf("not found ret resp.status:%v and url:%v", resp.StatusCode, checkUrl)
	}
	return string(ret[1]), nil
}

func (p *AtCoderJudge) submit(ctx context.Context) error {
	logger := utils.GetLogInstance()
	CID, _, _ := p.ParsePID(ctx)
	submitUrl := atcoderUrl + "/contests/" + CID + "/submit"
	var dataMap = map[string]string{
		"data.TaskScreenName": p.PID,
		"data.LanguageId":     atcoderLang[p.Submit.Lang],
		"sourceCode":          p.Submit.Source,
		"csrf_token":          p.JudgeUser.CsrfToken,
	}
	data := MapToFormStrings(dataMap, "&")
	resp, err := DoRequest(POST, submitUrl, p.Headers, p.JudgeUser.Cookies, &data, false)
	if err != nil {
		return err
	}
	SetCookies(resp, &p.JudgeUser.OriginJudgeUser)
	p.SubmissionID, err = p.CheckAndGetSubmissionID(ctx, resp)
	if err != nil {
		logger.Errorf("submit SourceCode failed,submit:%+v,err:%v", p.Submit.PID, err.Error())
		return fmt.Errorf("submit SourceCode failed,submit:%+v,err:%v", p.Submit.PID, err.Error())
	}
	p.JudgeUser.LastSubmitTime = time.Now().UnixNano()
	return nil
}

func (p *AtCoderJudge) RetJudgeUser(ctx context.Context) {
	if p.JudgeUser == nil {
		return
	}
	atlock.Lock()
	defer atlock.Unlock()
	p.JudgeUser.LastSubmitTime = time.Now().UnixNano()
	p.JudgeUser.Status = JUDGE_FREE
	p.JudgeUser = nil
}

func (p *AtCoderJudge) CheckResult(ctx context.Context, Text []byte) bool {
	re := regexp.MustCompile(`<span class='label label-.*?' data-toggle='tooltip' data-placement='top' title=".*?">([A-Za-z]*)</span>`)
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
		// 此处不设置cookie
		// SetCookies(resp, &p.JudgeUser.OriginJudgeUser)
		Text, _ := ParseRespToByte(resp)
		if !p.CheckResult(ctx, Text) {
			time.Sleep(time.Second)
			continue
		}
		/*解析内存和*/
		re := regexp.MustCompile(`<td class="text-center">(.*?)</td>`)
		ret := re.FindAllSubmatch(Text, 9)
		timere := regexp.MustCompile(`[0-9]*`)
		time := "0"
		if len(ret) > 6 {
			time = timere.FindString(string(ret[6][1]))
		}
		p.Submit.Usetime, _ = strconv.ParseInt(time, 10, 64)
		memory := "0"
		if len(ret) > 7 {
			memory = timere.FindString(string(ret[7][1]))
		}
		p.Submit.UseMemory, _ = strconv.ParseInt(memory, 10, 64)
		p.Submit.UseMemory *= 1024
		return nil
	}
}

func (p *AtCoderJudge) commitToDB(ctx context.Context) error {
	if p.Submit.Result == constanct.OJ_JUDGE {
		p.Submit.Result = constanct.OJ_FAILED
	}
	return models.UpdateSubmit(context.Background(), p.Submit)
}
