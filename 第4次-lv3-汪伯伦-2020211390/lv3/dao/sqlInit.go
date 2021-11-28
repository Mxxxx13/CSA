// @Title : sqlInit
// @Description :数据库初始化
// @Author : MX
// @Update : 2021/11/28 14:22 

package dao

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// MysqlInit 数据库初始化
func MysqlInit() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/blog?charset=utf8")
	if err != nil {
		fmt.Printf("mysql connect failed:%v", err)
		return nil
	}
	db.SetMaxOpenConns(1000)
	err = db.Ping()
	if err != nil {
		fmt.Printf("mysql connect failed:%v", err)
		return nil
	}
	DB = db
	return DB
}