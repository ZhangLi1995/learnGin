package routers

import (
	"learnGin/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")
	router.GET("/register", controllers.Register)
	return router
}
