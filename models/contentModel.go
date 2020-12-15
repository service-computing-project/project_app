/*
 * @Descripttion:content数据库调用
 * @version:1.0
 * @Author: sunylin
 * @Date: 2020-12-15 17:25:48
 * @LastEditors: sunylin
 * @LastEditTime: 2020-12-15 18:04:57
 */
package models

import (
	"errors"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//ContentDB content数据库
type ContentDB struct {
	DB *mgo.Collection
}

// AddContent 增加内容
func (m *ContentDB) AddContent(detail string, tag []string, ownID string, isPublic bool) (bson.ObjectId, error) {
	var content Content
	content.ID = bson.NewObjectId()
	content.Detail = detail
	content.OwnID = bson.ObjectIdHex(ownID)
	content.PublishDate = time.Now().Unix() * 1000
	content.LikeNum = 0
	content.Public = isPublic
	content.Tag = tag
	err := m.DB.Insert(content)
	return content.ID, err
}

// RemoveContent 删除内容
func (m *ContentDB) RemoveContent(id string) (err error) {
	if !bson.IsObjectIdHex(id) {
		return errors.New("not_id")
	}
	err = m.DB.RemoveId(bson.ObjectIdHex(id))
	return
}
