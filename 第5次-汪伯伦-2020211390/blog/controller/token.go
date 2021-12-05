// @Title : token
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/4 22:24 

package controller

import (
	"net/http"

	"blog/util"
	"github.com/gin-gonic/gin"
)

func LoginRequired(c *gin.Context) {
	//token := c.GetHeader("Authorization")
	//// 获取token
	//cookie, err := c.Request.Cookie("token")
	//if err != nil {
	//	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
	//		"code":    2000,
	//		"message": "需要登录",
	//	})
	//}
	//fmt.Println(cookie.Value)

	token := c.PostForm("token")
	//根据token解析出jwt
	jwt, err := util.CheckJWT(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    2000,
			"message": "需要登录",
		})
		return
	}

	// 设置id和username方便后续操作
	c.Set("id", jwt.Payload.Sub.Id)
	c.Set("username", jwt.Payload.Sub.Username)

	c.Next()
}
