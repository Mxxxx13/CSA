// @Title : main
// @Description :用户注册登录
// @Author : MX
// @Update : 2021/11/26 20:19 

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 存储用户信息
var userInfo = make(map[string]string)

func main() {
	r := gin.Default()
	r.GET("/register", Register)
	r.GET("/login", Login)
	r.GET("/hello", Hello)
	r.Run(":8080")
}

func Register(c *gin.Context) {
	var message, username, password string

	username = c.PostForm("username")
	password = c.PostForm("password")
	// 用户注册错误返回错误信息
	if username == "" {
		message = "用户名不能为空"
	}
	if password == "" {
		message = "密码不能为空"
	}
	if _, ok := userInfo[username]; ok {
		message = "用户已存在"
	}
	if message != "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": message,
		})
		return
	}
	userInfo[username] = password
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "注册成功",
	})
}

func Login(c *gin.Context) {
	var message string
	username := c.PostForm("username")
	password := c.PostForm("password")
	// 用户登录错误返回错误信息
	if password == userInfo[username] {
		message = "登录成功"
		// 成功登录设置cookie
		cookie := &http.Cookie{
			Name: "username",
			Value: username,
			Path: "/",
			HttpOnly: true,
		}
		http.SetCookie(c.Writer,cookie)
	} else {
		message = "用户名或密码错误"
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": message,
	})
}

func Hello(c *gin.Context){
	var message string
	cookie,err :=  c.Request.Cookie("username")
	// 当cookie未设置或为空时标记为游客
	if err != nil || cookie.Value == ""{
		message = "游客你好!"
	} else {
		message = cookie.Value + "你好!"
	}
	c.JSON(http.StatusOK,gin.H{
		"code": 200,
		"message" : message,
	})
}