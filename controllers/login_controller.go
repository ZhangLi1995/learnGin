package controllers

import (
	"learnGin/constants"
	"learnGin/dal"
	"learnGin/utils"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

/**
 * @Description: 获取登陆页
 * @param c
 */
func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{"title": "登陆页"})
}

/**
 * @Description: 用户登陆
 * @param c
 */
func Login(c *gin.Context) {
	userName := c.PostForm("username")
	password := c.PostForm("password")
	logrus.Infof("[Login] username: %v, password: %v", userName, password)

	/* 参数校验 */
	if utils.StrIsBlank(userName) || utils.StrIsBlank(password) {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "参数不合法"})
		return
	}

	id, err := dal.QueryUserWithParam(userName, utils.MD5(password))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": utils.ConvertBizErr(err).ErrCode(), "message": utils.ConvertBizErr(err).ErrMsg()})
		return
	} else if id > 0 {
		c.JSON(http.StatusOK, gin.H{"code": constants.SuccessCode, "message": "登陆成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": constants.FailedCode, "message": "登陆失败"})
	}

}
