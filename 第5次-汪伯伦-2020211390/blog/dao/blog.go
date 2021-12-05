// @Title : blog
// @Description :blog 数据库相关操作
// @Author : MX
// @Update : 2021/11/28 16:32 

package dao

import (
	"fmt"
	"log"

	"blog/model"
)

// Upload 将blog存入数据库
func UploadBlog(blog model.Blog) (err error) {
	stmt, err := DB.Prepare("insert into blog (title,content,author) value (?,?,?)")
	if err != nil {
		log.Printf("Upload:prepare failed: %v", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(blog.Title, blog.Content, blog.Author)
	if err != nil {
		log.Printf("Upload:insert failed: %v", err)
		return
	}
	return nil
}

// Like 根据blog id查询like, like+1后更新数据库
func LikeBlog(id int) (err error){
	stmt1, err := DB.Query("select likes from blog where id = ?", id)
	if err != nil {
		log.Printf("Like:query failed:%v", err)
		return
	}
	defer stmt1.Close()
	var like int
	for stmt1.Next() {
		err = stmt1.Scan(&like)
		if err != nil {
			log.Printf("Like:scan failed: %v", err)
			return
		}
	}
	like ++

	stmt2, err := DB.Prepare("update blog set likes = ? where id = ?")
	if err != nil {
		log.Printf("Like:prepare failed: %v", err)
		return
	}
	defer stmt2.Close()
	_, err = stmt2.Exec(like,id)
	if err != nil {
		log.Printf("Like:insert failed: %v", err)
		return
	}

	return nil
}

// Show 根据id查询blog
func ShowBlog(id int) (blog model.Blog,err error){
	stmt, err := DB.Query("select title,content,likes,author from blog where id = ?", id)
	if err != nil {
		fmt.Printf("Show:query failed:%v", err)
		return
	}
	defer stmt.Close()
	for stmt.Next() {
		err = stmt.Scan(&blog.Title, &blog.Content, &blog.Likes, &blog.Author)
		if err != nil {
			fmt.Printf("Show:scan failed: %v", err)
			return
		}
	}
	if blog.Title == "" && blog.Content == ""{
		return
	}
	return blog,nil
}


// AlterBlog 根据id修改blog
func AlterBlog(id int,blog model.Blog) (err error){
	stmt, err := DB.Prepare("update blog set title = ?, content = ? where id = ?")
	if err != nil {
		log.Printf("AlterBlog:prepare failed: %v", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(blog.Title,blog.Content,id)
	if err != nil {
		log.Printf("AlterBlog:insert failed: %v", err)
		return
	}
	return nil
}

// DeleteBlog 根据id删除blog
func DeleteBlog(id int) (err error){
	stmt, err := DB.Prepare("delete from blog where id = ?")
	if err != nil {
		log.Printf("DeleteBlog:prepare failed: %v", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		log.Printf("DeleteBlog:insert failed: %v", err)
		return
	}
	return nil
}