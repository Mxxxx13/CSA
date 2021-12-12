// @Title : blog
// @Description :blog 数据库相关操作
// @Author : MX
// @Update : 2021/11/28 16:32 

package dao

import (
	"blog/model"
)

// Upload 将blog存入数据库
func UploadBlog(blog model.Blog) (err error) {
	if err = DB.Create(&blog).Error; err != nil {
		return
	}
	return
}

// Like 根据blog id查询like, like+1后更新数据库
func LikeBlog(id int) (err error) {
	var blog model.Blog
	if err = DB.First(&blog, id).Error; err != nil {
		return
	}

	// 点赞数+1
	blog.Likes++

	if err = DB.Model(&model.Blog{}).Where("id", id).Update("likes", blog.Likes).Error; err != nil {
		return
	}
	return
}

// Show 根据id查询blog
func ShowBlog(id int) (blog model.Blog, err error) {
	if err = DB.First(&blog, id).Error; err != nil {
		return
	}
	return blog, nil
}

// AlterBlog 根据id修改blog
func AlterBlog(id int, blog model.Blog) (err error) {
	if err = DB.Model(&blog).Where("id = ?", id).Updates(model.Blog{
		Title:   blog.Title,
		Content: blog.Content,
	}).Error; err != nil {
		return
	}
	return
}

// DeleteBlog 根据id删除blog
func DeleteBlog(id int) (err error) {
	if err = DB.Delete(&model.Blog{}, id).Error; err != nil {
		return
	}
	return
}
