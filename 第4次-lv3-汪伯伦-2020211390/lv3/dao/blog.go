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
func Upload(blog model.Blog) bool {
	stmt, err := DB.Prepare("insert into blog (title,content,author) value (?,?,?)")
	if err != nil {
		log.Printf("Upload:prepare failed: %v", err)
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(blog.Title, blog.Content, blog.Author)
	if err != nil {
		log.Printf("Upload:insert failed: %v", err)
		return false
	}
	return true
}

// Like 根据blog id查询like, like+1后更新数据库
func Like(id int) bool{
	stmt1, err := DB.Query("select likes from blog where id = ?", id)
	if err != nil {
		fmt.Printf("Like:query failed:%v", err)
		return false
	}
	defer stmt1.Close()
	var like int
	for stmt1.Next() {
		err = stmt1.Scan(&like)
		if err != nil {
			fmt.Printf("Like:scan failed: %v", err)
			return false
		}
	}
	like ++

	stmt2, err := DB.Prepare("update blog set likes = ? where id = ?")
	if err != nil {
		log.Printf("Like:prepare failed: %v", err)
		return false
	}
	defer stmt2.Close()
	_, err = stmt2.Exec(like,id)
	if err != nil {
		log.Printf("Like:insert failed: %v", err)
		return false
	}

	return true
}

// Show 根据id查询blog
func Show(id int) (blog model.Blog,res bool){
	stmt, err := DB.Query("select title,content,likes,author from blog where id = ?", id)
	if err != nil {
		fmt.Printf("Show:query failed:%v", err)
		return blog,false
	}
	defer stmt.Close()
	for stmt.Next() {
		err = stmt.Scan(&blog.Title, &blog.Content, &blog.Likes, &blog.Author)
		if err != nil {
			fmt.Printf("Show:scan failed: %v", err)
			return blog,false
		}
	}
	return blog,true
}
