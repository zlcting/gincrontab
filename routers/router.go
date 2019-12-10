package routers

import (
	"github.com/gin-gonic/gin"
	"gospiderkeeper/controllers"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	router.LoadHTMLGlob("views/*")
	//注册：
	router.GET("/register", controllers.RegisterGet)
	router.GET("/crontabrun", controllers.Crontabrun)
	return router

} 
