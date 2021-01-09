package main

/*
数据解析绑定
*/
import "github.com/gin-gonic/gin"

func model() {
	router := gin.Default()
	router.POST("/loginJSON", loginJson)
	router.POST("/loginForm", loginForm)
	router.POST("/loginForm/v2", loginFormV2)
	router.POST("/:user/:password", loginUri)
	router.Run(":8080")
}
