package main

import (
	"learnGin/ginlearn/middleware/middleware"

	"github.com/gin-gonic/gin"
)

func register(r *gin.Engine) {

	/* 单个路由中间件 */
	r.GET("/before", middleware.MiddleWare(), singleMiddlewareTest)

	//r.Use(middleware.MiddleWare())
	r.GET("/middleware", middleware.MiddleWare(), middlewareTest)

	authorized := r.Group("/admin")
	authorized.Use(
		middleware.Auth(),
	)
	authorized.GET("/secrets", getSecrets)
}
