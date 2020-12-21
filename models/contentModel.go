/*
 * @Descripttion:content数据库调用
 * @version:1.0
 * @Author: sunylin
 * @Date: 2020-12-15 17:25:48
 * @LastEditors: sunylin
 * @LastEditTime: 2020-12-21 14:25:41
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
func (m *ContentDB) AddContent(detail string, tag []string, ownID string, isPublic bool) error {
	var content Content
	content.ID = bson.NewObjectId()
	content.Detail = detail
	content.OwnID = bson.ObjectIdHex(ownID)
	content.PublishDate = time.Now().Unix() * 1000
	content.LikeNum = 0
	content.Public = isPublic
	content.Tag = tag
	err := m.DB.Insert(content)
	return err
}

//UpdateContent 增加内容
func (m *ContentDB) UpdateContent(contentID, detail string, tag []string, ownID string, isPublic bool) error {

	c, err := m.DB.Find(bson.M{"_id": bson.ObjectIdHex(contentID), "ownId": bson.ObjectIdHex(ownID)}).Count()
	if err != nil {
		return err
	}
	if c == 0 {
		err = errors.New(StatusUserContentNotMatching)
		return err
	}

	var content Content
	content.ID = bson.ObjectIdHex(contentID)
	content.Detail = detail
	content.OwnID = bson.ObjectIdHex(ownID)
	content.PublishDate = time.Now().Unix() * 1000
	content.LikeNum = 0
	content.Public = isPublic
	content.Tag = tag
	err = m.DB.UpdateId(bson.ObjectIdHex(contentID), content)
	if err != nil {
		return err
	}

	err = m.DBuser.UpdateId(bson.ObjectIdHex(ownID), bson.M{"$inc": bson.M{"contentCount": 1}})
	return err
}

// RemoveContent 删除内容
func (m *ContentDB) RemoveContent(id string) (err error) {

	if !bson.IsObjectIdHex(id) {
		return errors.New(StatusNoID)
	}
	var content Content
	err = m.DB.FindId(bson.ObjectIdHex(id)).One(&content)
	if err != nil {
		return err
	}
	err = m.DB.RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		return err
	}
	err = m.DBuser.UpdateId(content.OwnID, bson.M{"$inc": bson.M{"contentCount": -1}})
	return
}

//GetDetailByID 获取指定内容
func (m *ContentDB) GetDetailByID(id string) (res ContentDetailres, err error) {
	if !bson.IsObjectIdHex(id) {
		res.State = StatusNoID
		err = errors.New(StatusNoID)
		return
	}
	var c Content
	err = m.DB.FindId(bson.ObjectIdHex(id)).One(&c)
	if err != nil {
		return
	}
	res.Data.Detail = c.Detail
	res.Data.ID = c.ID.Hex()
	res.Data.OwnID = c.OwnID.Hex()
	res.Data.PublishDate = c.PublishDate
	res.Data.LikeNum = c.LikeNum
	res.Data.Public = c.Public
	res.Data.Tag = c.Tag
	var u User
	err = m.DBuser.FindId(c.OwnID).One(&u)
	res.User.Avatar = u.Info.Avatar
	res.User.Gender = u.Info.Gender
	res.User.Name = u.Info.Name
	return
}

//GetPublic 获取公共内容
func (m *ContentDB) GetPublic(page, eachpage int) (res ContentPublicList, err error) {
	type AllContentID struct {
		Allid bson.ObjectId `bson:"_id"`
	}
	var all []AllContentID
	err = m.DB.Find(bson.M{"public": true}).Sort("-publishDate").Select(bson.M{"_id": 1}).All(&all)
	if err != nil {
		return
	}
	for _, value := range all[(page-1)*eachpage : page*eachpage] {
		var data ContentDetailres
		data, err = m.GetDetailByID(value.Allid.Hex())
		if err != nil {
			return
		}
		res.Data = append(res.Data, data)
	}
	return
}

//GetContentSelf 根据自己的用户id获取文章列表
func (m *ContentDB) GetContentSelf(id string, page, eachpage int) (res ContentListByUser, err error) {
	var c []Content
	err = m.DB.Find(bson.M{"ownId": bson.ObjectIdHex(id)}).All(&c)
	for _, value := range c[(page-1)*eachpage : page*eachpage] {
		var resc Contentres
		resc.Detail = value.Detail
		resc.ID = value.ID.Hex()
		resc.OwnID = value.OwnID.Hex()
		resc.PublishDate = value.PublishDate
		resc.LikeNum = value.LikeNum
		resc.Public = value.Public
		resc.Tag = value.Tag
		res.Data = append(res.Data, resc)
	}

	return
}

//GetContentByUser 获取他人的文章列表
func (m *ContentDB) GetContentByUser(id string, page, eachpage int) (res ContentListByUser, err error) {
	var c []Content
	err = m.DB.Find(bson.M{"ownId": bson.ObjectIdHex(id), "public": true}).All(&c)
	for _, value := range c[(page-1)*eachpage : page*eachpage] {
		var resc Contentres
		resc.Detail = value.Detail
		resc.ID = value.ID.Hex()
		resc.OwnID = value.OwnID.Hex()
		resc.PublishDate = value.PublishDate
		resc.LikeNum = value.LikeNum
		resc.Public = value.Public
		resc.Tag = value.Tag
		res.Data = append(res.Data, resc)
	}
	return
}
