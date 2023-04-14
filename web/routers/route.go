package routers

import (
	"ahutoj/web/controller"
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/middlewares"
	"ahutoj/web/utils"
	"net/http"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

func InitServer() {
	conf := utils.GetConfInstance()

	switch conf.Mode {
	case gin.DebugMode:
		gin.SetMode(gin.DebugMode)
	case gin.ReleaseMode:
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(middlewares.Cors())
	regeisterRouters(router)

	// 注册 pprof 监控 仅仅在 开发阶段可看
	if conf.Mode == gin.DebugMode {
		pprof.Register(router)
	}

	InitRouters(router)
	// 404
	router.NoRoute(NotFindRegister)

	router.Run(conf.Port)
}

func regeisterRouters(router *gin.Engine) {

	// 相当于接口 /api/ 这组路径
	apiRouter := router.Group("/api")
	{
		apiRouter.GET("/ping", PingTest) // 测试网络连通性
		apiRouter.GET("/now", serverTime)
		// 相当于接口/api/Auth/ 的这组路径
		authRouter := apiRouter.Group("/auth").Use(middlewares.JwtVerify)
		{
			// 相当于接口 /api/Auth/login
			authRouter.POST("/login/", controller.Login)
			authRouter.POST("/register/", controller.Register)
			authRouter.POST("/logout/", controller.Logout)
		}
		userRouter := apiRouter.Group("/user").Use(middlewares.JwtVerify)
		{
			userRouter.GET("/info", controller.UserInfo)
			userRouter.GET("/info/status", controller.UserStatusInfo)
			userRouter.POST("/edit/", controller.EditUserInfo)
			userRouter.POST("/edit/pass/", controller.EditUserPass)
			userRouter.POST("/vjudgeBind", controller.VjudgeBind)
			userRouter.POST("/CodeForceBind/", controller.CodeForceBind)
			userRouter.POST("/editHead/", controller.EditImage)
		}

		adminRouter := apiRouter.Group("/admin").Use(middlewares.JwtVerify)
		{
			adminRouter.POST("/permission/edit/", controller.EditPermission)
			adminRouter.POST("/permission/delete/", controller.DeletePermission)
			adminRouter.POST("/permission/add/", controller.AddPermission)
			adminRouter.GET("/permission/list/", controller.GetListPermission)
			adminRouter.GET("/permission/:id", controller.GetPermission)
			adminRouter.POST("/users/Range", controller.AddUsersRange)
			adminRouter.POST("/users", controller.AddUsers)

		}

		problemRouter := apiRouter.Group("/problem").Use(middlewares.JwtVerify)
		{
			// ->  /api/problems/add/'
			problemRouter.POST("/add/", controller.AddProblem)       // 添加题目
			problemRouter.POST("/edit/", controller.EditProblem)     // 编辑题目
			problemRouter.POST("/delete/", controller.DeleteProblem) // 删除题目
			problemRouter.GET("/list", controller.GetProblemList)    // 获取题目列表
			// param 可以获取id
			problemRouter.GET("/:id", controller.GetProblem) // 获取题目
		}

		trainingRouter := apiRouter.Group("/training").Use(middlewares.JwtVerify)
		{
			trainingRouter.POST("/add/", controller.AddTraining)
			trainingRouter.POST("/edit/", controller.EditTraining)
			trainingRouter.POST("/user/", controller.RegisterTraining)
			trainingRouter.GET("/user", controller.GetTrainUserInfo)
			trainingRouter.POST("/delete/", controller.DeleteTraining) // Lids []
			trainingRouter.GET("/list", controller.GetListTraining)
			trainingRouter.GET("/:id", controller.GetTraining)
			trainingRouter.GET("/:id/rank", controller.GetRankTraining)
		}

		contestRouter := apiRouter.Group("/contest").Use(middlewares.JwtVerify)
		{
			contestRouter.POST("/add/", controller.AddContest)
			contestRouter.POST("/edit/", controller.EditContest)
			contestRouter.POST("/delete/", controller.DeleteContest)

			contestRouter.GET("/list", controller.GetListContest)
			contestRouter.GET("/:id", controller.GetContest)
			contestRouter.GET("/:id/rank", controller.GteRankContest)
		}

		SubmitRouter := apiRouter.Group("/submit").Use(middlewares.JwtVerify)
		{
			SubmitRouter.POST("/commit/", controller.AddCommit)
			SubmitRouter.POST("/rejudge/", controller.RejudgeCommit)
			SubmitRouter.GET("/status", controller.StatusList)
			SubmitRouter.GET("/:id", controller.GetCommit)
		}

		fileRouter := apiRouter.Group("/file").Use(middlewares.JwtVerify)
		{
			// 上传判题文件
			fileRouter.POST("/:pid", controller.UpFile)
			// 获取判题文件列表
			fileRouter.GET("/:pid", controller.GetFileList)
			fileRouter.POST("/image/", controller.UpImagefile)
			// 删除文件
			fileRouter.DELETE("/:pid", controller.RemoveFile)
			// 解压文件
			fileRouter.POST("/unzip/:pid", controller.UnzipFile)
			// 上传并解析题目
			fileRouter.POST("/problem", controller.UpProblemFile)
			// 下载题目
			fileRouter.GET("/json/download", controller.DownloadProblemFromJson)
			// 上传题目
			fileRouter.POST("/problem/upfile/", controller.UpProblemFile)
		}
	}
}

func NotFindRegister(ctx *gin.Context) {
	response.ResponseError(ctx, constanct.PageNotFoundCode)
}

func PingTest(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    constanct.SuccessCode,
		"messgae": "",
	})
}

func serverTime(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    constanct.SuccessCode,
		"messgae": "success",
		"time":    time.Now().UnixMilli(),
	})
}

func InitRouters(router *gin.Engine) {
	conf := utils.GetConfInstance()
	for _, router := range router.Routes() {
		url := router.Path
		Method := router.Method
		req := request.AddRouterReq{
			FromURL:     url,
			Method:      Method,
			ToHost:      conf.Host + conf.Port,
			Weight:      10,
			VerfiyLevel: middlewares.GetVerifyUrl(url),
		}
		Header := make(map[string]string)
		Header["Content-Type"] = "application/json"
		dataByte, _ := json.Marshal(req)
		data := string(dataByte)
		utils.DoRequest(utils.POST, conf.GatWayHost+"inner/addrouter", Header, nil, &data, true)
	}
}
