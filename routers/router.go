package routers

import (
	"github.com/gin-gonic/gin"
	"gospiderkeeper/controllers"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	router.LoadHTMLGlob("views/*")
	//注册：
	router.GET("/crontablist", controllers.CrontabList)
	router.GET("/crontabrun", controllers.Crontabrun)
	router.GET("/crontabStop", controllers.CrontabStopAction)
	return router

}
