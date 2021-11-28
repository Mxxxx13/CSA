// @Title : user
// @Description :user 相关接口
// @Author : MX
// @Update : 2021/11/28 14:36 

package controller

import (
	"net/http"

	"blog/service"
	"github.com/gin-gonic/gin"
)

// Register 返回注册接口
func Register(c *gin.Context) {
	res := service.Register(c)
	if res {
		c.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"message":"注册成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"message":"用户已存在,注册失败",
		})
	}
}

//Login 返回登录接口,成功登录则设置cookie
func Login(c *gin.Context){
	res := service.Login(c)
	username := c.PostForm("username")
	if res {
		c.SetCookie("username", username, 240, "/", "127.0.0.1", false, true)
		c.JSON(http.StatusOK, gin.H{
			"code": 1002,
			"message":"欢迎回来" + username,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 2002,
			"message":"用户名或密码错误",
		})
	}
}

