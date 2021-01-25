package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/**
 * @Description: 默认服务器
 */
func serve1() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello World")
	})
	router.Run(":8000")
}

/**
 * @Description: http 服务器
 */
func serve2() {
	router := gin.Default()
	http.ListenAndServe(":8080", router)
}

/**
 * @Description: 自定义 http 服务配置
 */
func serve3() {
	router := gin.Default()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

/**
 * @Description: 路由参数: API 参数 + URL 参数
 */
func router1() {
	router := gin.Default()

	// 注意: /user/:name/ 会被重定向到 /user/:name
	router.GET("/user/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, name)
	})

	router.GET("/user/:name/*action", func(ctx *gin.Context) {
		name := ctx.Param("name")
		action := ctx.Param("action")
		message := name + " is " + action
		ctx.String(http.StatusOK, message)
	})

	router.GET("/welcome", func(ctx *gin.Context) {
		name := ctx.DefaultQuery("name", "Guest")
		ctx.String(http.StatusOK, fmt.Sprintf("Hello %s ", name))
	})
	router.Run(":8000")
}

/**
 * @Description: 表单参数
 */
func router2() {
	router := gin.Default()
	router.POST("/form", func(ctx *gin.Context) {
		type1 := ctx.DefaultPostForm("type", "alert")
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")

		hobbys := ctx.PostFormArray("hobby")

		ctx.String(http.StatusOK, fmt.Sprintf("type is %v, username is %v, password is %v, hobby is %v",
			type1, username, password, hobbys))
	})

	router.Run(":9527")
}

/**
 * @Description: 单个文件上传
 */
func router3() {
	router := gin.Default()
	// 设置最大内存限制为 8MB，默认为 32MB
	router.MaxMultipartMemory = 8 << 20 // 8MB

	router.POST("/upload", func(ctx *gin.Context) {
		file, _ := ctx.FormFile("file") // 单文件
		log.Println(file.Filename)
		ctx.SaveUploadedFile(file, file.Filename) // 上传文件到具体的位置
		/*
			也可以直接使用 io 操作，拷贝文件数据
			out, err := os.Create(filename)
			defer out.Close()
			_, err := io.Copy(out, file)
		*/
		ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	router.Run(":8080")
}

/**
 * @Description: 多个文件上传
 */
func router4() {
	router := gin.Default()
	// 设置最大内存限制为 8MB，默认为 32MB
	router.MaxMultipartMemory = 8 << 20 // 8MB

	router.POST("/upload", func(ctx *gin.Context) {
		form, err := ctx.MultipartForm() // 多文件
		if err != nil {
			ctx.String(http.StatusBadRequest, fmt.Sprintf("get form err: %v", err))
			return
		}
		files := form.File["files"]
		for _, file := range files {
			if err := ctx.SaveUploadedFile(file, file.Filename); err != nil {
				ctx.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %v", err))
				return
			}
		}
		ctx.String(http.StatusOK, fmt.Sprintf("Uploaded successfully %v files ", len(files)))
	})
	router.Run(":8080")
}

/**
 * @Description: 分组路由
 */
func router5() {
	router := gin.Default()

	// group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/login", loginEndpoint)
		v1.GET("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}

	// group: v2
	v2 := router.Group("/v2")
	{
		v2.GET("/login", loginEndpoint)
		v2.GET("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}

	router.Run(":8080")
}

func loginEndpoint(ctx *gin.Context) {
	name := ctx.DefaultQuery("name", "Guest")
	ctx.String(http.StatusOK, fmt.Sprintf("Hello %v \n", name))
}

func submitEndpoint(ctx *gin.Context) {
	name := ctx.DefaultQuery("name", "Guest")
	ctx.String(http.StatusOK, fmt.Sprintf("Hello %v \n", name))
}

func readEndpoint(ctx *gin.Context) {
	name := ctx.DefaultQuery("name", "Guest")
	ctx.String(http.StatusOK, fmt.Sprintf("Hello %v \n", name))
}
