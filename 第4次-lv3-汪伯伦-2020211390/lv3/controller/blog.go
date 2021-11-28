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

func Upload(c *gin.Context){
	res := service.Upload(c)
	if res {
		c.JSON(http.StatusOK, gin.H{
			"code": 1003,
			"message":"上传成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 2003,
			"message":"上传失败",
		})
	}
}

func Like(c *gin.Context) {
	res := service.Like(c)
	if res {
		c.JSON(http.StatusOK, gin.H{
			"code": 1004,
			"message":"点赞成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 2004,
			"message":"点赞失败",
		})
	}
}

func Show(c *gin.Context){
	blog, res := service.Show(c)
	if res {
		c.JSON(http.StatusOK, gin.H{
			"code": 1005,
			"message":"成功",
			"data":blog,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 2005,
			"message":"失败",
			"data":"",
		})
	}
}