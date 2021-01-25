package main

import "github.com/gin-gonic/gin"

func register(r *gin.Engine) {
	r.GET("/user", getUser)
	r.POST("/add", addUser)
	r.PUT("update", updateUser)
	r.DELETE("/delete/:id", deleteUser)
}
