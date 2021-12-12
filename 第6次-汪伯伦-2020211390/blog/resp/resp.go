// @Title : resp
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/7 21:54 

package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Reps 响应
type Reps struct {
	Code    int         `json:"code"` // 响应码
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SuccessResp 正确返回值
func SuccessResp(c *gin.Context, message string, data ...interface{}) {
	resp(c, 0, message, data...)
}

// ErrorResp 错误返回值
func ErrorResp(c *gin.Context, code int, message string, data ...interface{}) {
	resp(c, code, message, data...)
}

func resp(c *gin.Context, code int, message string, data ...interface{}) {
	resp := Reps{
		Code:    code,
		Message: message,
		Data:    data,
	}

	if len(data) == 1 {
		resp.Data = data[0]
	}

	c.JSON(http.StatusOK,resp)
}
