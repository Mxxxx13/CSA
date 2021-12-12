// @Title : comment
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/9 20:00 

package controller

import (
	"net/http"

	"blog/resp"
	"blog/service"
	"github.com/gin-gonic/gin"
)


func UploadComment(c *gin.Context)  {
	err := service.UploadComment(c)
	if err != nil {
		resp.ErrorResp(c,http.StatusBadRequest,"评论失败")
	} else {
		resp.SuccessResp(c,"评论成功")
	}
}
