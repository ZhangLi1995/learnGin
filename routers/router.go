package routers

import (
	"learnGin/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	/* 注册登陆 */
	router.LoadHTMLGlob("views/*")
	router.GET("/register", controllers.RegisterPage)
	router.POST("/register", controllers.RegisterUser)
	router.GET("/login", controllers.LoginPage)
	router.POST("/login", controllers.Login)

	return router
}
