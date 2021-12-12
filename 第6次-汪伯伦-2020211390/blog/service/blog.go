// @Title : blog
// @Description :blog 相关逻辑
// @Author : MX
// @Update : 2021/11/28 16:32 

package service

import (
	"errors"
	"log"
	"strconv"

	"blog/dao"
	"blog/model"
	"github.com/gin-gonic/gin"
)

func UploadBlog(c *gin.Context) (err error) {
	var blog model.Blog
	blog.Title = c.PostForm("title")
	blog.Content = c.PostForm("content")
	user, exists := c.Get("username")
	if !exists {
		return errors.New("user is not exist")
	}
	blog.Author = user.(string) // 接口断言
	err = dao.UploadBlog(blog)
	return err
}

func LikeBlog(c *gin.Context) (err error) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		log.Printf("atoi err:%v\n", err)
		return
	}
	err = dao.LikeBlog(id)
	return err
}

func ShowBlog(c *gin.Context) (BlogResp model.BlogResp, err error) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		log.Printf("atoi err:%v\n", err)
		return
	}
	blog, err := dao.ShowBlog(id)
	if err != nil{
		return
	}
	BlogResp = model.BlogResp{
		Title:       blog.Title,
		Content:     blog.Content,
		Author:      blog.Author,
		Likes:       blog.Likes,
		CommentResp: SplicingComment(int(blog.ID)),
	}
	return
}

func AlterBlog(c *gin.Context) (err error) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		log.Printf("atoi err:%v\n", err)
		return
	}
	var blog model.Blog
	blog.Title = c.PostForm("title")
	blog.Content = c.PostForm("content")
	err = dao.AlterBlog(id,blog)
	return err
}

func DeleteBlog(c *gin.Context) (err error) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		log.Printf("atoi err:%v\n", err)
		return
	}
	err = dao.DeleteBlog(id)
	return err
}
