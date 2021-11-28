// @Title : main
// @Description :main
// @Author : MX
// @Update : 2021/11/28 14:23 

package main

import (
	"net/http"

	"blog/controller"
	"blog/dao"
	"github.com/gin-gonic/gin"
)

func main() {
	dao.MysqlInit()
	r := gin.Default()
	r.POST("/register",controller.Register)
	r.GET("/login",controller.Login)
	login := r.Group("/")
	login.Use(LoginRequired())
	{
		login.POST("/blog/upload",controller.Upload)
		login.PUT("/blog/like/:id",controller.Like)
	}
	r.GET("/blog/show/:id",controller.Show)
	_ = r.Run(":8080")
}

// LoginRequired 需要登录中间件
func LoginRequired() gin.HandlerFunc{
	return func(c *gin.Context) {
		cookie, err := c.Cookie("username")
		if err == nil && cookie !="" {
			c.Next()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "需要登录",
		})
		c.Abort()
	}
}
