package main

import (
	"fmt"
	"learnGin/ginlearn/dal/dal"
	"learnGin/ginlearn/dal/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getUser(c *gin.Context) {
	users, err := dal.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": users,
		"count":  len(users),
	})
}

func addUser(c *gin.Context) {
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		log.Fatal(err)
	}
	Id, err := dal.Add(user)
	fmt.Print("id=", Id)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s 插入成功", user.UserName),
	})
}

func updateUser(c *gin.Context) {
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		log.Fatal(err)
	}
	num, err := dal.Update(user)
	fmt.Print("num=", num)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("修改id: %d 成功", user.Id),
	})
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")

	Id, err := strconv.Atoi(id)

	if err != nil {
		log.Fatalln(err)
	}
	rows, err := dal.Delete(Id)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("delete rows ", rows)

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully deleted user: %s", id),
	})
}
