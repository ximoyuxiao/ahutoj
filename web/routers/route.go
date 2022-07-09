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
	conf := utils.GetInstance()

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
	router.GET("/ping", PingTest)
	// 相当于接口 /api/ 这组路径
	apiRouter := router.Group("/api")
	{
		//相当于接口/api/Auth/ 的这组路径
		authRouter := apiRouter.Group("/auth")
		{
			//相当于接口 /api/Auth/login
			authRouter.POST("/login/", service.LoginSerivce)
			authRouter.POST("/register/", service.RegisterService)
			authRouter.GET("/logout/")
		}

		userRouter := apiRouter.Group("/user").Use(middlewares.JwtVerify)
		{
			userRouter.GET("/info/")
			userRouter.POST("/edit/")
			userRouter.POST("/edit/pass/")
			userRouter.GET("/VjudgeBind/")
		}

		adminRouter := apiRouter.Group("/admin").Use(middlewares.JwtVerify)
		{
			adminRouter.PUT("permission/edit")
		}
		problemRouter := apiRouter.Group("/problem")
		{
			//->  /api/problem/problems
			problemRouter.GET("/problems/")
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
