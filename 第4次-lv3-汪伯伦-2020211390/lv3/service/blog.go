// @Title : blog
// @Description :blog 相关逻辑
// @Author : MX
// @Update : 2021/11/28 16:32 

package service

import (
	"log"
	"strconv"

	"blog/dao"
	"blog/model"
	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) bool{
	var blog model.Blog
	if err := c.ShouldBind(&blog); err!=nil {
		log.Printf("Upload:bind err:%v", err)
		return false
	}
	cookie,_ := c.Request.Cookie("username")
    blog.Author = cookie.Value
	res := dao.Upload(blog)
	return res
}

func Like(c *gin.Context) bool {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		log.Printf("atoi err:%v\n",err)
		return false
	}
	res := dao.Like(id)
	return res
}

func Show(c *gin.Context) (blog model.Blog,res bool){
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		log.Printf("atoi err:%v\n",err)
		return blog,false
	}
	blog, res = dao.Show(id)
	return blog, res
}