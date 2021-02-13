package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitMysql() {
	fmt.Println("InitMysql...")
	if db == nil {
		var err error
		db, err = sql.Open("mysql", "root:ci@tcp(127.0.0.1:3306)/ci?charset=utf8")
		if err != nil {
			log.Fatal(err)
		}
		CreateTableWithUser()
	}
}

func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
        id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
        user_name VARCHAR(64),
        password VARCHAR(64),
        status INT(4),
        create_time INT(10)
        );`
	ModifyDB(sql)
}

/**
 * @Description: 操作数据库
 * @param sql
 * @param args
 * @return int64
 * @return error
 */
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

/**
 * @Description: 查询数据库
 * @param sql
 * @return *sql.Row
 */
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}
