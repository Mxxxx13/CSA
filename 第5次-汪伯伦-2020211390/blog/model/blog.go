// @Title : blog
// @Description :博客模型
// @Author : MX
// @Update : 2021/11/28 11:06 

package model

// Blog 博客模型
type Blog struct {
	Title   string `form:"title" json:"title" binding:"required"`   // 标题
	Content string `form:"content" json:"content" binding:"required"` // 内容
	Author  string `json:"author"`  // 作者
	Likes   string `json:"likes"`   // 获赞
}
