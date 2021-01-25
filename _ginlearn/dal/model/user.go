package model

type User struct {
	Id       int    `json:"id" form:"id"`
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
}
