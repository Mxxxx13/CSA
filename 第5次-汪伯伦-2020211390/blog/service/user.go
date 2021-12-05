// @Title : user
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/11/28 14:38 

package service

import (
	"errors"

	"blog/dao"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) (err error){
	username := c.PostForm("username")
	password := c.PostForm("password")
	err = dao.Register(username,password)
	return
}

// Login 将输入的password和数据库中的password进行比较
func Login(c *gin.Context) (id int,err error) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	pass,err := dao.Login(username)	// 获取password
	id,err = dao.GetId(username)	// 获取id
	if err != nil {
		return
	}
	// 对password进行验证
	if pass != password {
		return id,errors.New("用户名或密码错误")
	}
	return id,nil
}
