// @Title : user
// @Description :用户模型
// @Author : MX
// @Update : 2021/11/28 11:06 

package model

import "github.com/jinzhu/gorm"

type User struct {
	*gorm.Model
	Username string `json:"username"` // 用户名
	Password string // 密码
}

type UserResp struct {
	Uid      uint       `json:"uid"`       // 用户id
	Username string     `json:"username"`  // 用户名
	BlogList []BlogList `json:"blog_list"` // 博客列表
}
