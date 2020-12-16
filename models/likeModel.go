/*
 * @Descripttion:
 * @version:
 * @Author: sunylin
 * @Date: 2020-12-16 15:03:45
 * @LastEditors: sunylin
 * @LastEditTime: 2020-12-16 16:02:23
 */
package models

import (
	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//LikeDB like数据库
type LikeDB struct {
	DBU *mgo.Collection //用户数据库
	DBL *mgo.Collection //点赞数据库
	DBC *mgo.Collection //文章数据库
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
	if c != 0 {
		err = errors.New(StatusLikeExist)
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
	var userID []bson.ObjectId
	c, err := m.DBC.FindId(bson.ObjectIdHex(Contentid)).Count()
	if c == 0 {
		err = errors.New(StatusNoContent)
		return
	}
	err = m.DBL.Find(bson.M{"contentId": bson.ObjectIdHex(Contentid)}).Select(bson.M{"contentId": 1}).All(&userID)
	if err != nil {
		return
	}
	for _, value := range userID {
		var likeUser User
		err = m.DBU.FindId(value).One(&likeUser)
		if err != nil {
			return
		}
		user = append(user, likeUser.Info.Name)
	}
	return
}
