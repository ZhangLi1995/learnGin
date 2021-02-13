package main

import (
	"learnGin/routers"
	"learnGin/utils"
)

func main() {
	utils.InitMysql()
	router := routers.InitRouter()
	/* 静态资源 */
	router.Static("/static", "./static")
	router.Run(":8081")
}
