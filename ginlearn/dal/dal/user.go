package dal

import (
	"database/sql"
	"learnGin/ginlearn/dal/model"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	getAll = "GetAll"
)

/**
 * @Description: 查询所有 user 信息
 * @return users
 * @return err
 */
func GetAll() (users []model.User, err error) {

	/* 操作数据库 */
	db, err := sql.Open("mysql", "root:ci@tcp(127.0.0.1:3306)/ci?charset=utf8")
	if err != nil {
		log.Printf("[%v]open db failed. err: %v\n", getAll, err)
		return nil, err
	}
	defer db.Close()

	/* 查询 */
	rows, err := db.Query("SELECT id, user_name, password FROM user")
	if err != nil {
		log.Printf("[%v] query raws failed. err: %v", getAll, err)
		return nil, err
	}
	for rows.Next() {
		var user model.User
		rows.Scan(&user.Password, &user.UserName, &user.Password)
		users = append(users, user)
	}
	defer rows.Close()
	return
}

/**
 * @Description: 插入数据
 * @param user
 * @return id
 * @return err
 */
func Add(user model.User) (id int, err error) {

	db, err := sql.Open("mysql", "root:ci@tcp(127.0.0.1:3306)/ci?charset=utf8")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO user(id, user_name, password) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return
	}

	result, err := stmt.Exec(user.Id, user.UserName, user.Password)
	if err != nil {
		log.Fatal(err)
		return
	}

	idd, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return
	}
	id = int(idd)
	defer stmt.Close()
	return
}

/**
 * @Description: 更新数据
 * @param user
 * @return rowsAffected
 * @return err
 */
func Update(user model.User) (rowsAffected int64, err error) {

	db, err := sql.Open("mysql", "root:ci@tcp(127.0.0.1:3306)/ci?charset=utf8")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE user SET user_name=?, password=? WHERE id=?")
	if err != nil {
		return
	}

	rs, err := stmt.Exec(user.UserName, user.Password, user.Id)
	if err != nil {
		return
	}

	rowsAffected, err = rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()
	return
}

/**
 * @Description: 删除数据
 * @param id
 * @return rows
 * @return err
 */
func Delete(id int) (rows int, err error) {

	db, err := sql.Open("mysql", "root:ci@tcp(127.0.0.1:3306)/ci?charset=utf8")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM user WHERE id=?")
	if err != nil {
		log.Fatalln(err)
		return
	}

	rs, err := stmt.Exec(id)
	if err != nil {
		log.Fatalln(err)
		return
	}
	row, err := rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()
	rows = int(row)
	return
}
