// @Title : user
// @Description :用户模型
// @Author : MX
// @Update : 2021/11/28 11:06 

package model

type User struct {
	Username string `json:"username"`	// 用户名
	Password string						// 密码
}