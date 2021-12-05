// @Title : blog
// @Description :blog 相关接口
// @Author : MX
// @Update : 2021/11/28 16:32 

package controller

import (
	"net/http"

	"blog/service"
	"github.com/gin-gonic/gin"
)

const (
	UploadBlogSuccess = 1003
	UploadBlogFailure = 2003
	LikeBlogSuccess = 1004
	LikeBlogFailure = 2004
	ShowBlogSuccess = 1005
	ShowBlogFailure = 2005
	AlterBlogSuccess = 1006
	AlterBlogFailure = 2006
	DeleteBlogSuccess = 1007
	DeleteBlogFailure = 2007
)

func UploadBlog(c *gin.Context){
	err := service.UploadBlog(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": UploadBlogFailure,
			"message":"上传失败",
		})

	} else {

		c.JSON(http.StatusOK, gin.H{
			"code": UploadBlogSuccess,
			"message":"上传成功",
		})
	}
}

func LikeBlog(c *gin.Context) {
	err := service.LikeBlog(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": LikeBlogFailure,
			"message":"点赞失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": LikeBlogSuccess,
			"message":"点赞成功",
		})
	}
}

func ShowBlog(c *gin.Context){
	blog, err := service.ShowBlog(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": ShowBlogFailure,
			"message":"失败",
			"data":"",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": ShowBlogSuccess,
			"message":"成功",
			"data":blog,
		})
	}
}

func AlterBlog(c *gin.Context){
	err := service.AlterBlog(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": AlterBlogFailure,
			"message":"修改失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": AlterBlogSuccess,
			"message":"修改成功",
		})
	}
}

func DeleteBlog(c *gin.Context){
	err := service.DeleteBlog(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": DeleteBlogFailure,
			"message":"删除失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": DeleteBlogSuccess,
			"message":"删除成功",
		})
	}
}