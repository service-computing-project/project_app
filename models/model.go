/*
 * @Descripttion:
 * @version:
 * @Author: sunylin
 * @Date: 2020-12-14 23:13:17
 * @LastEditors: sunylin
 * @LastEditTime: 2020-12-16 00:47:07
 */
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
	StatusEmailFormatError = "email_format_error"
	StatusPasswordError    = "password_error"
	StatusNoID             = "no_this_id"
	StatusEmptyName        = "name_nil"
	StatusEmptyEmail       = "email_nil"
)

//User

//User 用户信息
type User struct {
	ID           bson.ObjectId `bson:"_id"`          // 用户ID
	Pwd          string        `bson:"password"`     //用户密码
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
	State string
	ID    string
	Email string
	Info  UserInfo
}

//Commonres 通用回应
type CommonRes struct {
	State string
	Data  string
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

//ContentUserInfo 用于返回详情中的的用户信息
type ContentUserInfo struct {
	Name   string `bson:"name"`   // 用户昵称
	Avatar string `bson:"avatar"` // 头像URL
	Gender int    `bson:"gender"` // 性别
}

//ContentDetailres 返回详情
type ContentDetailres struct {
	State string          //状态
	Data  Content         //文档数据
	User  ContentUserInfo //用户数据（这里少一个bio字段）
}

//ContentPublicList 返回公共文章
type ContentPublicList struct {
	State string
	Data  []ContentDetailres
}

//ContentListByUser 根据用户返回文章
type ContentListByUser struct {
	State string
	Data  []Content
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
