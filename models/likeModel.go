/*
 * @Descripttion:
 * @version:
 * @Author: sunylin
 * @Date: 2020-12-16 15:03:45
 * @LastEditors: sunylin
 * @LastEditTime: 2020-12-18 01:58:45
 */
package models

import (
	"errors"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//LikeDB like数据库
type LikeDB struct {
	DBU *mgo.Collection //用户数据库
	DBL *mgo.Collection //点赞数据库
	DBC *mgo.Collection //文章数据库
	DBN *mgo.Collection //通知数据库
}

//LikeByID 通过用户id和文章id点赞
func (m *LikeDB) LikeByID(Contentid, Userid string) (err error) {
	c, err := m.DBC.FindId(bson.ObjectIdHex(Contentid)).Count()
	if c == 0 {
		err = errors.New(StatusNoContent)
		return
	}
	c, err = m.DBU.FindId(bson.ObjectIdHex(Userid)).Count()
	if c == 0 {
		err = errors.New(StatusNoUser)
		return
	}
	c, err = m.DBL.Find(bson.M{"contentId": bson.ObjectIdHex(Contentid), "userId": bson.ObjectIdHex(Userid)}).Count()
	if c != 0 {
		err = errors.New(StatusLikeExist)
		return
	}

	newLike := bson.NewObjectId()
	err = m.DBL.Insert(&Like{
		ID:        newLike,
		UserID:    bson.ObjectIdHex(Userid),
		ContentID: bson.ObjectIdHex(Contentid),
	})
	if err != nil {
		return
	}
	err = m.DBC.UpdateId(bson.ObjectIdHex(Contentid), bson.M{"$inc": bson.M{"likeNum": 1}})
	if err != nil {
		return
	}
	type notificationTarget struct {
		ContentOwner bson.ObjectId `bson:"ownId"`
		Content      string        `bson:"detail"`
	}
	var n notificationTarget
	err = m.DBC.FindId(bson.ObjectIdHex(Contentid)).Select(bson.M{"ownId": 1, "detail": 1}).One(&n)
	if err != nil {
		return
	}
	c, err = m.DBN.Find(bson.M{"sourceId": bson.ObjectIdHex(Userid), "contentId": bson.ObjectIdHex(Contentid)}).Count()
	if c != 0 {
		err = errors.New(StatusNotificationExist)
		return
	}
	newNotification := bson.NewObjectId()
	err = m.DBN.Insert(&NotificationDetail{
		ID:         newNotification,
		CreateTime: time.Now().Unix() * 1000,
		Content:    n.Content,
		SourceID:   bson.ObjectIdHex(Userid),
		ContentID:  bson.ObjectIdHex(Contentid),
		TargetID:   n.ContentOwner,
		Type:       "like",
	})

	return
}

//CancelLikeByID 通过用户id和文章id取消点赞
func (m *LikeDB) CancelLikeByID(Contentid, Userid string) (err error) {
	c, err := m.DBC.FindId(bson.ObjectIdHex(Contentid)).Count()
	if c == 0 {
		err = errors.New(StatusNoContent)
		return
	}
	c, err = m.DBU.FindId(bson.ObjectIdHex(Userid)).Count()
	if c == 0 {
		err = errors.New(StatusNoUser)
		return
	}
	c, err = m.DBL.Find(bson.M{"contentId": bson.ObjectIdHex(Contentid), "userId": bson.ObjectIdHex(Userid)}).Count()
	if c == 0 {
		err = errors.New(StatusLikeNotExist)
		return
	}

	err = m.DBL.Remove(bson.M{"contentId": bson.ObjectIdHex(Contentid), "userId": bson.ObjectIdHex(Userid)})
	if err != nil {
		return
	}

	err = m.DBC.UpdateId(bson.ObjectIdHex(Contentid), bson.M{"$inc": bson.M{"likeNum": -1}})
	return
}

//GetUserListByContentID 通过文章id获取点赞的用户列表
func (m *LikeDB) GetUserListByContentID(Contentid string) (user []string, err error) {
	type TempUser struct {
		UserID bson.ObjectId `bson:"userId"`
	}
	var userid []TempUser
	c, err := m.DBC.FindId(bson.ObjectIdHex(Contentid)).Count()
	if c == 0 {
		err = errors.New(StatusNoContent)
		return
	}
	err = m.DBL.Find(bson.M{"contentId": bson.ObjectIdHex(Contentid)}).Select(bson.M{"userId": 1}).All(&userid)
	if err != nil {
		return
	}
	for _, value := range userid {
		var likeUser User
		err = m.DBU.FindId(value.UserID).One(&likeUser)
		if err != nil {
			return
		}
		user = append(user, likeUser.Info.Name)
	}
	return
}
