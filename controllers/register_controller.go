package controllers

import (
	"learnGin/dal"
	"learnGin/models"
	"learnGin/utils"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{"title": "注册页"})
}

func RegisterUser(c *gin.Context) {
	/* 获取表单信息 */
	userName := c.PostForm("user_name")
	password := c.PostForm("password")
	rePassword := c.PostForm("repassword")
	logrus.Infof("user name: %v, password: %v, re password: %v", userName, password, rePassword)

	/* 参数校验 */
	if utils.StrIsBlank(userName) || utils.StrIsBlank(password) || utils.StrIsBlank(rePassword) {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "参数不合法"})
		return
	}
	if password != rePassword {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "两次密码不一致"})
		return
	}

	/* 判断该用户名是否已经注册 */
	id, err := dal.QueryUserWithName(userName)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": err.Error()})
		return
	}
	logrus.Infof("[RegisterUser] query result by name(%v) is id = %v", userName, id)
	if id > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "用户名已经存在"})
		return
	}

	password = utils.MD5(password)
	user := &models.User{0, userName, password, 0, time.Now().Unix()}
	_, err = dal.InsertUser(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "注册失败"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "注册成功"})
	}
}
