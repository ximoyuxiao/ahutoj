package routers

import (
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
		authRouter := apiRouter.Group("/Auth")
		{
			//相当于接口 /api/Auth/login
			authRouter.POST("/login/", service.LoginSerivce)
			authRouter.POST("/regeister/")
			authRouter.GET("/logout/")
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
