package routers

import (
	"ahutoj/web/controller"
	"ahutoj/web/middlewares"
	"ahutoj/web/utils"
	"fmt"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func InitForumServer() {
	conf := utils.GetConfInstance()

	switch conf.Mode {
	case gin.DebugMode:
		gin.SetMode(gin.DebugMode)
	case gin.ReleaseMode:
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(middlewares.Cors())
	regeisterForumRouters(router)

	// 注册 pprof 监控 仅仅在 开发阶段可看
	if conf.Mode == gin.DebugMode {
		pprof.Register(router)
	}

	InitRouters(router, fmt.Sprintf("%v:%v", conf.ForumConfig.Host, conf.ForumConfig.Port))
	// 404
	router.NoRoute(NotFindRegister)

	router.Run(fmt.Sprintf(":%v", conf.ForumConfig.Port))
}
func regeisterForumRouters(router *gin.Engine) {
	// 相当于接口 /api/ 这组路径
	// 1. 帖子
	apiRouter := router.Group("/api")
	{
		forumGroup := apiRouter.Group("/solution").Use(middlewares.JwtVerify)
		forumGroup.GET("/ping", PingTest)
		forumGroup.GET("/:id", controller.GetSolution)
		forumGroup.GET("/solutions", controller.GetSoulutions)
		forumGroup.POST("/", controller.SolutionOperator)
		forumGroup.PUT("/:id", controller.SolutionOperator)
		forumGroup.DELETE("/:id", controller.SolutionOperator)
		// 2. 点赞
		favoriteGropu := apiRouter.Group("/favorite").Use(middlewares.JwtVerify)
		favoriteGropu.GET("/:id", controller.GetFaviroate)
		favoriteGropu.POST("/", controller.FavoriteAction)
		// 3. 评论
		commentGroup := apiRouter.Group("/comment").Use(middlewares.JwtVerify)
		commentGroup.GET("/:id", controller.GetComment)
		commentGroup.GET("/comments", controller.GetComments)
		commentGroup.POST("/", controller.CommentOperator)
		commentGroup.DELETE("/:id", controller.CommentOperator)
	}

}