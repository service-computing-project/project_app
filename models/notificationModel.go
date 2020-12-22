/*
 * @Descripttion:
 * @version:
 * @Author: sunylin
 * @Date: 2020-12-16 16:06:27
 * @LastEditors: sunylin
 * @LastEditTime: 2020-12-22 19:19:49
 */
package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//NotifiationDB 通知数据库
type NotifiationDB struct {
	DBN *mgo.Collection //通知数据库
	DBU *mgo.Collection //用户数据库
}

//GetNotificationByUserID 获取指定用户通知
func (m *NotifiationDB) GetNotificationByUserID(id string) (res UserNotificationres, err error) {
	var AllNotification []NotificationDetail
	err = m.DBN.Find(bson.M{"targetId": bson.ObjectIdHex(id)}).All(&AllNotification)
	if err != nil {
		return
	}
	for _, value := range AllNotification {
		var n Notificationres
		var u User
		err = m.DBU.FindId(value.SourceID).One(&u)
		n.User.Avatar = u.Info.Avatar
		n.User.Gender = u.Info.Gender
		n.User.Name = u.Info.Name
		n.Notification = value
		res.Notifications = append(res.Notifications, n)
	}
	return
}

//DeleteNotificationByID 删除指定的内容
func (m *NotifiationDB) DeleteNotificationByID(id string) (err error) {
	err = m.DBN.RemoveId(bson.ObjectIdHex(id))
	return
}
