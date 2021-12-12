// @Title : blog
// @Description :博客模型
// @Author : MX
// @Update : 2021/11/28 11:06 

package model

import "github.com/jinzhu/gorm"

// Blog 博客模型
type Blog struct {
	*gorm.Model
	Title   string `form:"title" json:"title" binding:"required"`   // 标题
	Content string `form:"content" json:"content" binding:"required"` // 内容
	Author  string `json:"author"`  // 作者
	Likes   int `json:"likes"`   // 获赞
}

type BlogResp struct {
	//Blog
	Title   string `form:"title" json:"title" binding:"required"`   // 标题
	Content string `form:"content" json:"content" binding:"required"` // 内容
	Author  string `json:"author"`  // 作者
	Likes   int `json:"likes"`   // 获赞
	CommentResp []CommentResp
}