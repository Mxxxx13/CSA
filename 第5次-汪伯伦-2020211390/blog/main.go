// @Title : main
// @Description :main
// @Author : MX
// @Update : 2021/11/28 14:23 

package main

import (
	"blog/controller"
	"blog/dao"
	"github.com/gin-gonic/gin"
)

func main() {
	dao.MysqlInit()
	r := gin.Default()
	login := r.Group("")
	login.Use(controller.LoginRequired)

	r.POST("/register", controller.Register)
	r.GET("/login", controller.Login)

	r.GET("/blog/:id", controller.ShowBlog)
	login.POST("/blog", controller.UploadBlog)
	login.PUT("/blog/:id", controller.AlterBlog)
	login.DELETE("/blog/:id", controller.DeleteBlog)
	login.PUT("/blog/like/:id", controller.LikeBlog)

	_ = r.Run(":8080")
}
