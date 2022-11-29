package main

import (
	flowcontrol "ahutoj/gatway/flowControl"
	"ahutoj/gatway/parsejwt"
	originJudged "ahutoj/originJudge/originjudged"
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/middlewares"
	"ahutoj/web/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var appconf = "./config.yaml"

var conf struct {
	Mode    string `mapstructure:"mode"`
	Sign    string `mapstructure:"sign"`
	Version string `mapstructure:"version"`
	GatWay  struct {
		Port   string `mapstructure:"port"`
		MaxQPS int    `mapstructure:"maxQPS"`

		BlackHost []string `mapstructure:"blackReq"`
	} `mapstructure:"GatWay"`
}

var PrefixAndRouter = make(map[string]*Router, 0)

type Router struct {
	From       string
	Method     string
	To         []Target
	Permission middlewares.VerfiyLevel
}

type Target struct {
	Host       string
	Weight     int64
	Use        int64
	Connection bool
}

func main() {
	if len(os.Args) >= 2 {
		appconf = os.Args[1]
	}
	err := initGatWay(appconf)
	fmt.Printf("error server down!,err:%v\n", err.Error())
}

var router *gin.Engine = nil
var tb *flowcontrol.TokenBucket = nil

func initGatWay(config string) error {
	err := utils.InitAppConfig(config, &conf)
	if err != nil {
		fmt.Println("call InitAppConfig failed")
	}
	parsejwt.InitJwt(conf.Sign)

	switch conf.Mode {
	case gin.DebugMode:
		gin.SetMode(gin.DebugMode)
	case gin.ReleaseMode:
		gin.SetMode(gin.ReleaseMode)
	}
	tb = flowcontrol.InitTokenBucket(1000, 50)
	GatWay := &conf.GatWay
	router = gin.Default()
	router.Use(middlewares.Cors())
	router.Handle("POST", "/inner/addrouter", AddRouter)
	router.Handle("DELETE", "/inner/delrouter", DelRouter)
	router.NoRoute(HandleRouter)
	err = router.Run(GatWay.Port)
	return fmt.Errorf("gat way use error:%v", err.Error())
}

func HandleRouter(ctx *gin.Context) {
	url := ctx.FullPath()
	host := ctx.ClientIP()
	// 保证不再黑名单当中
	for _, blackhost := range conf.GatWay.BlackHost {
		if blackhost == host {
			ctx.JSON(403, response.Response{
				StatusCode: 0,
				StatusMsg:  "请求失败，请重试",
			})
			return
		}
	}
	// 从路由表中 获得
	fmt.Printf("route url:%v method:%v\n", url, ctx.Request.Method)
	key := strings.ToUpper(ctx.Request.Method) + " " + url
	Mrouter, ok := PrefixAndRouter[key]
	if !ok {
		response.ResponseError(ctx, constanct.PageNotFoundCode)
		return
	}
	res, ok := parsejwt.JwtVerify(ctx)
	if !ok {
		response.ResponseOK(ctx, res)
		return
	}
	var fetchToken int64 = 0
	if tb != nil {
		if fetchToken <= 0 {
			fetchToken = tb.FetchToken(1)
		}
		if fetchToken <= 0 {
			response.ResponseError(ctx, constanct.ServerBusyCode)
			return
		}
	}
	resp := switchRouter(ctx, Mrouter, DefaultGetHost)
	if resp == nil {
		response.ResponseError(ctx, constanct.ServerErrorCode)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var data1 map[string]interface{}
	json.Unmarshal(body, &data1)
	ctx.JSON(resp.StatusCode, data1)
}

// 负载均衡 策略
func DefaultGetHost(ctx *gin.Context, target []Target) *Target {
	start := rand.Int() % len(target)
	fmt.Printf("%v", utils.Sdump(target))
	for idx := range target {
		i := (idx + start) % len(target)
		if target[i].Connection {
			fmt.Printf("use Connection ID:%v\n", i)
			return &target[i]
		}
	}
	return nil
}

// 路由转发
func switchRouter(ctx *gin.Context, router *Router, GetHost func(*gin.Context, []Target) *Target) *http.Response {
	method := ctx.Request.Method
	Host := GetHost(ctx, router.To)
	if Host == nil {
		return nil
	}
	url := "http://" + Host.Host + ctx.Request.URL.String()
	header := make(map[string]string, 0)
	for k, v := range ctx.Request.Header {
		header[k] = v[0]
	}
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	defer ctx.Request.Body.Close()
	bodyString := string(body)
	resp, err := originJudged.DoRequest(originJudged.HttpMethodType(method), url, header, nil, &bodyString, true)
	if err != nil {
		Host.Connection = false
		return nil
	}
	return resp
}

// 添加路由
func AddRouter(ctx *gin.Context) {
	req := new(request.AddRouterReq)
	if err := ctx.ShouldBindBodyWith(req, binding.JSON); err != nil {
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	key := strings.ToUpper(req.Method) + " " + req.FromURL
	_, ok := PrefixAndRouter[key]
	// 如果不存在这个key
	fmt.Printf("url:%v level:%v\n", req.FromURL, req.VerfiyLevel)
	if !ok {
		PrefixAndRouter[key] = &Router{
			From:       req.FromURL,
			Method:     strings.ToUpper(req.Method),
			To:         make([]Target, 0),
			Permission: middlewares.VerfiyLevel(req.VerfiyLevel),
		}
		router.Handle(req.Method, req.FromURL, HandleRouter)
		parsejwt.VerifyMap[req.FromURL] = req.VerfiyLevel
	}
	// 存在这个key 保证不重复添加 找找有没有重复的
	for idx := range PrefixAndRouter[key].To {
		if PrefixAndRouter[key].To[idx].Host == req.ToHost {
			PrefixAndRouter[key].To[idx].Weight = req.Weight
			PrefixAndRouter[key].To[idx].Connection = true
			parsejwt.VerifyMap[req.FromURL] = req.VerfiyLevel
			response.ResponseOK(ctx, constanct.SuccessCode)
			return
		}
	}
	PrefixAndRouter[key].To = append(PrefixAndRouter[key].To, Target{
		Host:       req.ToHost,
		Weight:     req.Weight,
		Use:        0,
		Connection: true,
	})
	response.ResponseOK(ctx, constanct.SuccessCode)
}

func DelRouter(ctx *gin.Context) {
	req := new(request.DelRouterReq)
	if err := ctx.ShouldBindBodyWith(req, binding.JSON); err != nil {
		response.ResponseError(ctx, constanct.InvalidParamCode)
		return
	}
	key := strings.ToUpper(req.Method) + " " + req.FromURL
	_, ok := PrefixAndRouter[key]
	if !ok {
		response.ResponseOK(ctx, constanct.SuccessCode)
		return
	}
	for idx := range PrefixAndRouter[key].To {
		if PrefixAndRouter[key].To[idx].Host == req.ToHost {
			PrefixAndRouter[key].To[idx].Weight = 0
			PrefixAndRouter[key].To[idx].Connection = false
			response.ResponseOK(ctx, constanct.SuccessCode)
			return
		}
	}
}
