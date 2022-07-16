package routers

import (
	"ahutoj/web/middlewares"
	"ahutoj/web/service"
	"ahutoj/web/utils"
	"net/http"

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

	//404
	router.NoRoute(NotFindRegister)

	router.Run(conf.Port)
}

func regeisterRouters(router *gin.Engine) {
	router.GET("/ping", PingTest) //测试网络连通性
	// 相当于接口 /api/ 这组路径
	apiRouter := router.Group("/api")
	{
		//相当于接口/api/Auth/ 的这组路径
		authRouter := apiRouter.Group("/auth")
		{
			//相当于接口 /api/Auth/login
			authRouter.POST("/login/", service.Login)
			authRouter.POST("/register/", service.Register)
			authRouter.POST("/logout/", service.Logout)
		}

		userRouter := apiRouter.Group("/user").Use(middlewares.JwtVerify)
		{
			userRouter.GET("/info", service.UserInfo)
			userRouter.POST("/edit/", service.EditUserInfo)
			userRouter.POST("/edit/pass/", service.EditUserPass)
			userRouter.GET("/vjudgeBind", service.VjudgeBind)
		}

		adminRouter := apiRouter.Group("/admin").Use(middlewares.JwtVerify)
		{
			adminRouter.POST("/permission/edit/")
			adminRouter.POST("/permission/delete/")
			adminRouter.POST("/permission/add/")
			adminRouter.GET("/permission/list/")
			adminRouter.GET("/permission/:id")
		}

		problemRouter := apiRouter.Group("/problem")
		{
			//->  /api/problem/problems'
			problemRouter.POST("/add/", service.AddService)       //添加题目
			problemRouter.POST("/edit/", service.EditService)     //编辑题目
			problemRouter.POST("/delete/", service.DeleteService) //删除题目

			problemRouter.GET("/list", service.GetListService) //获取题目列表

			// param 可以获取id
			problemRouter.GET("/:id", service.GetService) //获取题目
		}

		trainingRouter := apiRouter.Group("/training")
		{
			trainingRouter.POST("/add/")
			trainingRouter.POST("/edit/")

			trainingRouter.POST("/delete/")

			trainingRouter.GET("/list")
			trainingRouter.GET("/:id")
			trainingRouter.GET("/:id/rank")
		}

		contestRouter := apiRouter.Group("/contest")
		{
			contestRouter.POST("/add/")
			contestRouter.POST("/edit/")
			contestRouter.POST("/delete/")

			contestRouter.GET("/list")
			contestRouter.GET("/:id")
			contestRouter.GET("/:id/rank")
		}

		fileRouter := apiRouter.Group("/file")
		{
			fileRouter.PUT("/add/", service.UpFile)
			fileRouter.DELETE("/delete/", service.RemoveFile)
			fileRouter.POST("/unzip/", service.UnzipFile)
		}
	}
}

func NotFindRegister(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "404",
	})
}

func PingTest(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"messgae": "pong",
	})
}
