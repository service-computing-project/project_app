/*
 * @Descripttion:content数据库调用
 * @version:1.0
 * @Author: sunylin
 * @Date: 2020-12-15 17:25:48
 * @LastEditors: sunylin
 * @LastEditTime: 2020-12-15 22:22:22
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
	DB     *mgo.Collection
	DBuser *mgo.Collection
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
		return errors.New(StatusNoID)
	}
	err = m.DB.RemoveId(bson.ObjectIdHex(id))
	return
}

//GetDetailByID 获取指定内容
func (m *ContentDB) GetDetailByID(id string) (res ContentDetailres, err error) {
	if !bson.IsObjectIdHex(id) {
		res.State = StatusNoID
		err = errors.New(StatusNoID)
		return
	}
	err = m.DB.FindId(bson.ObjectIdHex(id)).One(&res.Data)
	if err != nil {
		return
	}
	err = m.DBuser.FindId(res.Data.OwnID).Select(bson.M{"info.name": 1, "info.avatar": 1, "info.gender": 1}).One(&res.User)
	return
}

//GetPublic 获取公共内容
func (m *ContentDB) GetPublic() (res ContentPublicList, err error) {
	var Allid []bson.ObjectId
	err = m.DB.Find(bson.M{"public": true}).Select(bson.M{"_id": 1}).All(&Allid)
	if err != nil {
		return
	}
	for _, value := range Allid {
		var data ContentDetailres
		data, err = m.GetDetailByID(value.Hex())
		if err != nil {
			return
		}
		res.Data = append(res.Data, data)
	}
	return
}

//GetContentSelf 根据自己的用户id获取文章列表
func (m *ContentDB) GetContentSelf(id string) (res ContentListByUser, err error) {
	err = m.DB.Find(bson.M{"ownId": bson.ObjectIdHex(id)}).All(&res.Data)
	return
}

//GetContentByUser 获取他人的文章列表
func (m *ContentDB) GetContentByUser(id string) (res ContentListByUser, err error) {
	err = m.DB.Find(bson.M{"ownId": bson.ObjectIdHex(id), "public": true}).All(&res.Data)
	return
}
