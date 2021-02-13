package models

type User struct {
	Id         int
	UserName   string
	Password   string
	Status     int // 0: 正常状态，1: 已删除
	CreateTime int64
}
