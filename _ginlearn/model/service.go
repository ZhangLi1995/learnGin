package main

import (
	"net/http"

	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
)

type Login struct {
	User     string `form:"username" json:"user"     uri:"user"     xml:"user"     binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

/**
 * @Description: Json 绑定
 * @param c
 */
func loginJson(c *gin.Context) {
	var json Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if json.User != "yuyu" || json.Password != "yuyu123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "login success"})
}

/**
 * @Description: Form 绑定
 * @param c
 */
func loginForm(c *gin.Context) {
	var form Login
	// 方法一: 对于 FORM 数据直接使用 Bind 函数，默认使用 form 格式解析
	// 根据请求头中的 content-type 自动推断
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if form.User != "yuyu" || form.Password != "yuyu123" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "unauthorized"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "login success"})
}

/**
 * @Description: Form 绑定
 * @param c
 */
func loginFormV2(c *gin.Context) {
	var form Login
	// 方法二: 使用 BindWith 函数，如果明确知道数据类型，可以显示声明来绑定多媒体表单，或者使用自动推断
	if c.BindWith(&form, binding.Form) == nil {
		if form.User == "yuyu" && form.Password == "yuyu123" {
			c.JSON(http.StatusOK, gin.H{"status": "login success"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"status": "unauthorized"})
		}
	}
}

/**
 * @Description: uri 绑定
 * @param c
 */
func loginUri(c *gin.Context) {
	var login Login
	if err := c.ShouldBindUri(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}
	c.JSON(http.StatusOK, gin.H{"status": "login success"})
}
