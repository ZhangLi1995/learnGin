package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	register(router)
	router.Run(":8080")
}
