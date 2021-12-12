// @Title : main
// @Description :main
// @Author : MX
// @Update : 2021/11/28 14:23 

package main

import (
	"blog/controller"
	"blog/dao"
	"blog/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	dao.ConnDB()
	r := gin.Default()
	login := r.Group("")
	login.Use(middleware.LoginRequired)

	r.POST("/user", controller.Register)
	r.GET("/user/login", controller.Login)
	r.GET("/user/:id",controller.ShowUser)
	r.PUT("/user/:id",controller.AlterUser)

	r.GET("/blog/:id", controller.ShowBlog)
	login.POST("/blog", controller.UploadBlog)
	login.PUT("/blog/:id", controller.AlterBlog)
	login.DELETE("/blog/:id", controller.DeleteBlog)
	login.PUT("/blog/like/:id", controller.LikeBlog)

	login.POST("/comment",controller.UploadComment)
	//login.DELETE("/comment/:id",controller.DeleteComment)


	_ = r.Run(":8080")
}
