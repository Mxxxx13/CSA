// @Title : blog
// @Description :blog 相关接口
// @Author : MX
// @Update : 2021/11/28 16:32 

package controller

import (
	"net/http"

	"blog/resp"
	"blog/service"
	"github.com/gin-gonic/gin"
)

func UploadBlog(c *gin.Context){
	err := service.UploadBlog(c)
	if err != nil {
		resp.ErrorResp(c,http.StatusBadRequest,"上传失败")
	} else {
		resp.SuccessResp(c,"上传成功")
	}
}

func LikeBlog(c *gin.Context) {
	err := service.LikeBlog(c)
	if err != nil {
		resp.ErrorResp(c,http.StatusBadRequest,"点赞失败")
	} else {
		resp.SuccessResp(c,"点赞成功")
	}
}

func ShowBlog(c *gin.Context){
	blog, err := service.ShowBlog(c)
	if err != nil {
		resp.ErrorResp(c,http.StatusBadRequest,"操作失败")
	} else {
		resp.SuccessResp(c,"操作成功",blog)
	}
}

func AlterBlog(c *gin.Context){
	err := service.AlterBlog(c)
	if err != nil {
		resp.ErrorResp(c,http.StatusBadRequest,"修改失败")
	} else {
		resp.SuccessResp(c,"修改成功")
	}
}

func DeleteBlog(c *gin.Context){
	err := service.DeleteBlog(c)
	if err != nil {
		resp.ErrorResp(c,http.StatusBadRequest,"删除失败")
	} else {
		resp.SuccessResp(c,"删除成功")
	}
}