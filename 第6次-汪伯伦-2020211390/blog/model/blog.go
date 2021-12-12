// @Title : blog
// @Description :博客模型
// @Author : MX
// @Update : 2021/11/28 11:06 

package model

import "github.com/jinzhu/gorm"

// Blog 博客模型
type Blog struct {
	*gorm.Model
	Title   string `form:"title" json:"title" binding:"required"`     // 标题
	Content string `form:"content" json:"content" binding:"required"` // 内容
	Uid     uint   `json:"uid"`                                       // 作者id
	Author  string `json:"author"`                                    // 作者
	Likes   int    `json:"likes"`                                     // 获赞
}

type BlogResp struct {
	Title       string `form:"title" json:"title" binding:"required"`     // 标题
	Content     string `form:"content" json:"content" binding:"required"` // 内容
	Author      string `json:"author"`                                    // 作者
	Likes       int    `json:"likes"`                                     // 获赞
	CommentResp []CommentResp
}

type BlogList struct {
	Bid   uint   `json:"bid"`   // 博客id
	Title string `json:"title"` // 标题
}
