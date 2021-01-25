package middleware

import "github.com/gin-gonic/gin"

/* 模拟私有数据 */
var Secrets = gin.H{
	"zhangsan": gin.H{"email": "zhangsan@163.com", "phone": "123456"},
	"lisi":     gin.H{"email": "lisi@163.com", "phone": "666"},
	"wangwu":   gin.H{"email": "wangwu@163.com", "phone": "555"},
}

func Auth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"zhangsan": "zhangsan123",
		"lisi":     "1234",
		"wangwu":   "hello2",
		"dingliu":  "4321",
	})
}
