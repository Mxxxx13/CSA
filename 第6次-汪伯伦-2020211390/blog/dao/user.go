// @Title : user
// @Description :user 数据库相关操作
// @Author : MX
// @Update : 2021/11/28 14:38 

package dao

import (
	"blog/model"
)

//Register 将username,password插入数据库
func Register(user model.User) (err error) {
	if err = DB.Create(&user).Error; err != nil {
		return
	}
	return
}

// Login 根据username查询password并返回
func Login(username string) (password string, err error) {
	var user model.User
	if err = DB.Where("username = ?", username).First(&user).Error; err != nil {
		return
	}
	return user.Password, nil
}

func GetUid(username string) (id int, err error) {
	var user model.User
	if err = DB.Where("username = ?", username).First(&user).Error; err != nil {
		return
	}
	return int(user.ID), nil
}

func GetUsername(uid uint) (username string, err error) {
	var user model.User
	if err = DB.Select("username").Where("id = ?", uid).First(&user).Error; err != nil {
		return
	}
	return user.Username, err
}
