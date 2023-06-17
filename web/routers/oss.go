package routers

import (
	"ahutoj/web/controller"
	"ahutoj/web/middlewares"
	"ahutoj/web/utils"
	"fmt"
	"io/ioutil"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func InitOssServer() {
	conf := utils.GetConfInstance()

	switch conf.Mode {
	case gin.DebugMode:
		gin.SetMode(gin.DebugMode)
	case gin.ReleaseMode:
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(middlewares.Cors())
	regeisterOSSRouters(router)

	// 注册 pprof 监控 仅仅在 开发阶段可看
	if conf.Mode == gin.DebugMode {
		pprof.Register(router)
	}

	InitRouters(router, fmt.Sprintf("%v:%v", conf.OssConfig.Host, conf.OssConfig.Port))
	// 404
	router.NoRoute(NotFindRegister)
	router.Run(fmt.Sprintf(":%v", conf.OssConfig.Port))
}
func SwitchStaticFile(router *gin.Engine, path string) {
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		if file.IsDir() {
			router.Static("/"+file.Name(), path+file.Name())
		}
	}
}
func regeisterOSSRouters(router *gin.Engine) {
	SwitchStaticFile(router, utils.GetConfInstance().OssConfig.BasePath)

	apiRouter := router.Group("api")
	{
		// 对象存储
		apiRouter.GET("/object", controller.GetObject)
		// 获取某个文件夹下面的所有文件
		apiRouter.GET("/object/files", controller.GetObjects)
		apiRouter.POST("/object/", controller.CreateObject)
		apiRouter.PUT("/object/", controller.ModifyObject)
		apiRouter.DELETE("/object", controller.DeleteObject)
		apiRouter.HEAD("/object", controller.GetObjectInfo)
	}
}
