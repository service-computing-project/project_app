/*
 * @Descripttion:content数据库调用
 * @version:1.0
 * @Author: sunylin
 * @Date: 2020-12-15 17:25:48
 * @LastEditors: sunylin
 * @LastEditTime: 2020-12-15 17:48:53
 */
package models

import "gopkg.in/mgo.v2"

//ContentDB content数据库
type ContentDB struct {
	DB *mgo.Collection
}

// AddContent 增加内容
// func (m *ContentDB) AddContent(content Content) (bson.ObjectId, error) {
// 	content.ID = bson.NewObjectId()
// 	content.PublishDate = time.Now().Unix() * 1000
// 	content.EditDate = time.Now().Unix() * 1000
// 	content.LikeNum = 0
// 	content.CommentNum = 0
// 	err := m.DB.Insert(content)
// 	return content.ID, err
// }
