// @Title : user
// @Description :user 数据库相关操作
// @Author : MX
// @Update : 2021/11/28 14:38 

package dao

import (
	"fmt"
	"log"
)

//Register 将username,password插入数据库
func Register(username, password string) bool {
	stmt, err := DB.Prepare("insert into user (username,password) value (?,?)")
	if err != nil {
		log.Printf("Register:prepare failed: %v", err)
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, password)
	if err != nil {
		log.Printf("Register:insert failed: %v", err)
		return false
	}
	return true
}

// Login 根据username查询password并返回
func Login(username string) (password string, res bool) {
	stmt, err := DB.Query("select password from user where username = ?", username)
	if err != nil {
		fmt.Printf("Login:query failed:%v", err)
		return "",false
	}
	defer stmt.Close()

	for stmt.Next() {
		err = stmt.Scan(&password)
		if err != nil {
			fmt.Printf("Login:scan failed: %v", err)
			return "",false
		}
	}
	return password,true
}