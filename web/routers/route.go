package routers

import (
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/response"
	"ahutoj/web/middlewares"
	"ahutoj/web/service"
	"ahutoj/web/utils"
	"net/http"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
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

	// 404
	router.NoRoute(NotFindRegister)

	router.Run(conf.Port)
}

func regeisterRouters(router *gin.Engine) {
	router.GET("/ping", PingTest) // 测试网络连通性
	// 相当于接口 /api/ 这组路径
	apiRouter := router.Group("/api")
	{
		apiRouter.GET("/now", serverTime)
		// 相当于接口/api/Auth/ 的这组路径
		authRouter := apiRouter.Group("/auth")
		{
			// 相当于接口 /api/Auth/login
			authRouter.POST("/login/", service.Login)
			authRouter.POST("/register/", service.Register)
			authRouter.POST("/logout/", service.Logout)
		}

		userRouter := apiRouter.Group("/user").Use(middlewares.JwtVerify)
		{
			userRouter.GET("/info", service.UserInfo)
			userRouter.GET("/info/status", service.UserStatusInfo)
			userRouter.POST("/edit/", service.EditUserInfo)
			userRouter.POST("/edit/pass/", service.EditUserPass)
			userRouter.POST("/vjudgeBind", service.VjudgeBind)
			userRouter.POST("/CodeForceBind/", service.CodeForceBind)
		}

		adminRouter := apiRouter.Group("/admin").Use(middlewares.JwtVerify)
		{
			adminRouter.POST("/permission/edit/", service.EditPermission)
			adminRouter.POST("/permission/delete/", service.DeletePermission)
			adminRouter.POST("/permission/add/", service.AddPermission)
			adminRouter.GET("/permission/list/", service.GetListPermission)
			adminRouter.GET("/permission/:id", service.GetPermission)
			adminRouter.POST("/users/Range", service.AddUsersRange)
			adminRouter.POST("/users", service.AddUsers)

		}

		problemRouter := apiRouter.Group("/problem").Use(middlewares.JwtVerify)
		{
			// ->  /api/problems/add/'
			problemRouter.POST("/add/", service.AddProblem)       // 添加题目
			problemRouter.POST("/edit/", service.EditProblem)     // 编辑题目
			problemRouter.POST("/delete/", service.DeleteProblem) // 删除题目
			problemRouter.GET("/list", service.GetProblemList)    // 获取题目列表
			// param 可以获取id
			problemRouter.GET("/:id", service.GetProblem) // 获取题目
		}

		trainingRouter := apiRouter.Group("/training").Use(middlewares.JwtVerify)
		{
			trainingRouter.POST("/add/", service.AddTraining)
			trainingRouter.POST("/edit/", service.EditTraining)

			trainingRouter.POST("/delete/", service.DeleteTraining) // Lids []
			trainingRouter.GET("/list", service.GetListTraining)
			trainingRouter.GET("/:id", service.GetTraining)
			trainingRouter.GET("/:id/rank", service.GetRankTraining)
		}

		contestRouter := apiRouter.Group("/contest").Use(middlewares.JwtVerify)
		{
			contestRouter.POST("/add/", service.AddContest)
			contestRouter.POST("/edit/", service.EditContest)
			contestRouter.POST("/delete/", service.DeleteContest)

			contestRouter.GET("/list", service.GetListContest)
			contestRouter.GET("/:id", service.GetContest)
			contestRouter.GET("/:id/rank", service.GteRankContest)
		}

		SubmitRouter := apiRouter.Group("/submit").Use(middlewares.JwtVerify)
		{
			SubmitRouter.POST("/commit/", service.AddCommit)
			SubmitRouter.POST("/rejudge/", service.RejudgeCommit)
			SubmitRouter.GET("/status", service.StatusList)
			SubmitRouter.GET("/:id", service.GetCommit)
		}

		fileRouter := apiRouter.Group("/file").Use(middlewares.JwtVerify)
		{
			// 上传判题文件
			fileRouter.POST("/:pid", service.UpFile)
			// 获取判题文件列表
			fileRouter.GET("/:pid", service.GetFileList)
			// 删除文件
			fileRouter.DELETE("/:pid", service.RemoveFile)
			// 解压文件
			fileRouter.POST("/unzip/:pid", service.UnzipFile)
			// 上传并解析题目
			fileRouter.POST("/problem", service.UpProblemFile)
		}
	}
}

func NotFindRegister(ctx *gin.Context) {
	response.ResponseError(ctx, constanct.PageNotFoundCode)
}

func PingTest(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    constanct.SuccessCode,
		"messgae": "pong",
	})
}

func serverTime(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    constanct.SuccessCode,
		"messgae": "success",
		"time":    time.Now().UnixMilli(),
	})
}
