package models

import "gopkg.in/mgo.v2/bson"

//Status
const (
	StatusSuccess          = "success"
	StatusBadReq           = "bad_req"
	StatusNotLogin         = "not_login"
	StatusUserNameExist    = "username_exist"
	StatusUserNameNotExist = "username_notexist"
	StatusEmailExist       = "email_exist"
	StatusPasswordError    = "password_error"
)

//User

//User 用户信息
type User struct {
	ID           bson.ObjectId `bson:"_id"`          // 用户ID
	Email        string        `bson:"email"`        // 用户唯一邮箱
	Info         UserInfo      `bson:"info"`         // 用户个性信息
	LikeCount    int64         `bson:"likeCount"`    // 被点赞数
	ContentCount int64         `bson:"contentCount"` // 内容数量
}

// UserInfo 用户个性信息
type UserInfo struct {
	Name   string `bson:"name"`   // 用户昵称
	Avatar string `bson:"avatar"` // 头像URL
	Bio    string `bson:"bio"`    // 个人简介
	Gender int    `bson:"gender"` // 性别
}

//UserInfoRes 获取用户信息的回应
type UserInfoRes struct {
	ID    string
	State string
	Email string
	Name  string
	Class int
	Info  UserInfo
}

//content

// Content 正文信息
type Content struct {
	ID          bson.ObjectId `bson:"_id"`
	Detail      string        `bson:"detail"`      // 详情介绍
	OwnID       bson.ObjectId `bson:"ownId"`       // 作者ID [索引]
	PublishDate int64         `bson:"publishDate"` // 发布日期
	LikeNum     int64         `bson:"likeNum"`     // 点赞数
	Public      bool          `bson:"public"`      // 是否公开
	Tag         []string      `bson:"tag"`         // 标签
}

//like

//Like 点赞信息
type Like struct {
	ID        bson.ObjectId `bson:"_id"`
	UserID    bson.ObjectId `bson:"userId"`        // 用户ID
	ContentID bson.ObjectId `bson:"notifications"` // 内容ID
}

//notification

//Notification 用户通知集合
type Notification struct {
	ID            bson.ObjectId        `bson:"_id"`
	UserID        bson.ObjectId        `bson:"userId"`        // 用户ID 【索引】
	Notifications []NotificationDetail `bson:"notifications"` // 通知集合
}

//NotificationDetail 通知详情
type NotificationDetail struct {
	ID         bson.ObjectId `bson:"_id"`
	CreateTime int64         `bson:"time"`
	Content    string        `bson:"content"`  // 通知内容
	SourceID   string        `bson:"sourceId"` // 源ID （点赞人）
	TargetID   string        `bson:"targetId"` // 目标ID （点赞文章）
	Type       string        `bson:"type"`     // 类型：暂时只有like
}
