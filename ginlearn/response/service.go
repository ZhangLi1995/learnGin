package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin/testdata/protoexample"

	"github.com/gin-gonic/gin"
)

func someJson(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
}

func moreJson(c *gin.Context) {
	var msg struct {
		Name    string `json:"user"`
		Message string `json:"message"`
		Number  int    `json:"number"`
	}
	msg.Name = "yuyu"
	msg.Message = "hey"
	msg.Number = 123
	c.JSON(http.StatusOK, msg)
}

func someXml(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{"user": "yuyu", "message": "hey", "status": http.StatusOK})
}

func someYaml(c *gin.Context) {
	c.YAML(http.StatusOK, gin.H{"user": "yuyu", "message": "hey", "status": http.StatusOK})
}

func someProtoBuf(c *gin.Context) {
	reps := []int64{int64(1), int64(2)}
	label := "test"
	data := &protoexample.Test{
		Label: &label,
		Reps:  reps,
	}
	c.ProtoBuf(http.StatusOK, data)
}

func htmlRender(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{"titile": "Main website"})
}

func htmlRenderV2(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Posts",
	})
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Users",
	})
}

func redirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
}

func async(c *gin.Context) {
	/* goroutine 只能使用只读的上下文 c.Copy() */
	cCp := c.Copy()
	go func() {
		time.Sleep(5 * time.Second)
		log.Println("Done! in path" + cCp.Request.URL.Path)
	}()
}

func sync(c *gin.Context) {
	time.Sleep(5 * time.Second)
	log.Println("Done! in path" + c.Request.URL.Path)
}
