package main

import (
	"fmt"
	"learnGin/ginlearn/middleware/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func middlewareTest(c *gin.Context) {
	request := c.MustGet("request")
	req, _ := c.Get("request")
	fmt.Println("request: ", request)
	c.JSON(http.StatusOK, gin.H{
		"middle_request": request,
		"request":        req,
	})
}

func singleMiddlewareTest(c *gin.Context) {
	request := c.MustGet("request").(string)
	c.JSON(http.StatusOK, gin.H{
		"middile_request": request,
	})
}

func getSecrets(c *gin.Context) {
	/* 获取提交的用户名 */
	user := c.MustGet(gin.AuthUserKey).(string)
	if secret, ok := middleware.Secrets[user]; ok {
		c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
	}
}
