package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerRouter(r *gin.Engine) {

	/* Json/Xml/Yaml 渲染 */
	r.GET("/someJSON", someJson)
	r.GET("/moreJSON", moreJson)
	r.GET("/someXML", someXml)
	r.GET("/someYAML", someYaml)
	r.GET("/someProtoBuf", someProtoBuf)

	/* html 模板渲染 */
	//r.LoadHTMLGlob("templates/*")
	//r.GET("/index", htmlRender)
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/posts/index", htmlRenderV2)

	/* 文件响应 */
	// 静态文件服务：显示当前文件目录下的所有文件或指定文件
	r.StaticFS("/showDir", http.Dir("."))
	//r.StaticFS("/files", http.Dir("/bin"))
	// Static 方法提供给定文件系统根目录中的文件
	r.Static("/files", "/bin")
	r.StaticFile("/image", "./assets/miao.jpg")

	/* 重定向 */
	r.GET("/redirect", redirect)

	/* 同步异步 */
	r.GET("/long_async", async)
	r.GET("/long_sync", sync)
}
