// @Title : user
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/11/28 14:38 

package service

import (
	"errors"
	"log"
	"strconv"

	"blog/dao"
	"blog/model"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) (err error){
	var user model.User
	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")
	err = dao.Register(user)
	return
}

// Login 将输入的password和数据库中的password进行比较
func Login(c *gin.Context) (uid uint,err error) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	pass,err := dao.Login(username)	// 获取password
	uid,err = dao.GetUid(username)	// 获取id
	if err != nil {
		return
	}
	// 对password进行验证
	if pass != password {
		return uid,errors.New("用户名或密码错误")
	}
	return
}

// ShowUser
func ShowUser(c *gin.Context) (UserResp model.UserResp,err error){
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		log.Printf("atoi err:%v\n", err)
		return
	}
	username, err := dao.GetUsername(uint(id))
	if err != nil {
		return
	}
	blogs, err := dao.GetBlog(uint(id))
	if err != nil {
		return
	}
	var blogList []model.BlogList
	for _, blog := range blogs {
		temp := model.BlogList{
			Bid:   blog.ID,
			Title: blog.Title,
		}
		blogList = append(blogList,temp)
	}
	UserResp = model.UserResp{
		Uid:      uint(id),
		Username: username,
		BlogList: blogList,
	}
	return
}

// AlterUser
func AlterUser(c *gin.Context) (err error){
	username := c.PostForm("username")
	id, exists := c.Get("uid")
	if !exists  {
		return errors.New("id not exist")
	}
	err = dao.AlterUser(username,id.(uint))
	return
}