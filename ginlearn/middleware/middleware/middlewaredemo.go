package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("before middleware")
		/* 设置 request 变量到 Context 的 Key 中，通过 Get 等函数可以得到 */
		c.Set("request", "client_request")
		/* 发送 request 之前 */
		/* 请求过来后先执行到 Next() 之前，再执行业务函数，再执行 Next() 后面的部分 */
		c.Next()
		// 发送 request 之后
		status := c.Writer.Status() // 得到状态等信息
		fmt.Println("after middleware, ", status)
		t2 := time.Since(t)
		fmt.Println("time: ", t2)
	}
}
