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
	//SwitchStaticFile(router, utils.GetConfInstance().OssConfig.BasePath)

	apiRouter := router.Group("api").Use(middlewares.JwtVerify)
	{
		//-----------对象操作--------------

		apiRouter.POST("/object/", controller.GetObject) //获取base64
		// 获取某个桶下面的所有文件
		apiRouter.GET("/object/:bucket", controller.GetObjects)
		//获取还未上传完成的列表
		//apiRouter.GET("/object/files", controller.GetUpingObjects)
		//从本地上传，暂时不需要上传其他来源数据，以后做(
		//apiRouter.POST("/object/", controller.CreateObject)
		apiRouter.POST("/object/delete/", controller.DeleteObject)
		apiRouter.POST("/object/getfile/", controller.FGetObject) //下载本地
		apiRouter.POST("/object/putfile/", controller.FPutObject) //本地上传
		//apiRouter.PUT("/object/", controller.ModifyObject)
		apiRouter.POST("/object/info/", controller.GetObjectInfo)
		//apiRouter.POST("/object/unzip", controller.UnzipObject)
		//-----------桶操作--------------
		//获得所有桶的名称+创建日期
		apiRouter.GET("/bucket", controller.GetBuckets)
		apiRouter.POST("/bucket/add/", controller.CreateBucket)
		apiRouter.POST("/bucket/delete/", controller.RemoveBucket)
	}
}
