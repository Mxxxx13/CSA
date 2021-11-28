// @Title : user
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/11/28 14:38 

package service

import (
	"blog/dao"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context)  bool{
	username := c.PostForm("username")
	password := c.PostForm("password")
	res := dao.Register(username,password)
	return res
}

// Login 将输入的password和数据库中的password进行比较
func Login(c *gin.Context) bool {
	username := c.PostForm("username")
	password := c.PostForm("password")
	pass,res := dao.Login(username)
	if !res {
		return false
	}
	if pass == password {
		return true
	}
	return false
}
