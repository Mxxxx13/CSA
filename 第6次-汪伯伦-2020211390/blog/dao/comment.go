// @Title : comment
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/9 20:00 

package dao

import "blog/model"

func UploadComment(comment model.Comment) (err error) {
	if err = DB.Create(&comment).Error;err != nil {
		return
	}
	return
}

func GetComments(bid int) (comments []model.Comment, err error) {
	if err = DB.Where("bid = ?",bid).Find(&comments).Error;err != nil {
		return
	}
	return
}