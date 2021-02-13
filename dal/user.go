package dal

import (
	"fmt"
	"learnGin/models"
	"learnGin/utils"

	"github.com/sirupsen/logrus"
)

/**
 * @Description: 插入用户信息
 * @param user
 * @return int64
 * @return error
 */
func InsertUser(user *models.User) (int64, error) {
	return utils.ModifyDB("insert into users(user_name,password,status,create_time) values (?,?,?,?)",
		user.UserName, user.Password, user.Status, user.CreateTime)
}

/**
 * @Description: 按照条件查询 id
 * @param cond
 * @param id
 */
func QueryUserWithCond(cond string) (int, error) {
	sql := fmt.Sprintf("SELECT id FROM users %v", cond)
	logrus.Infof("[QueryUserWithCond] query sql: %v", sql)
	row := utils.QueryRowDB(sql)
	if row.Err() != nil {
		logrus.Errorf("[QueryUserWithCond] query db failed. err: %v", row.Err())
		return 0, row.Err()
	}
	var id int
	row.Scan(&id)
	return id, nil
}

/**
 * @Description: 根据名字查询 id
 * @param name
 * @param id
 */
func QueryUserWithName(name string) (int, error) {
	sql := fmt.Sprintf("WHERE user_name = '%v'", name)
	return QueryUserWithCond(sql)
}
