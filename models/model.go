/*
 * @Descripttion:
 * @version:
 * @Author: sunylin
 * @Date: 2020-12-14 23:13:17
 * @LastEditors: sunylin
 * @LastEditTime: 2020-12-21 22:06:08
 */
package models

import "gopkg.in/mgo.v2/bson"

//Status
const (
	StatusSuccess                = "success"
	StatusBadReq                 = "bad_req"
	StatusNotLogin               = "not_login"
	StatusUserNameExist          = "username_exist"
	StatusUserNameNotExist       = "username_notexist"
	StatusEmailExist             = "email_exist"
	StatusEmailFormatError       = "email_format_error"
	StatusPasswordError          = "password_error"
	StatusNoID                   = "no_this_id"
	StatusEmptyName              = "name_nil"
	StatusEmptyEmail             = "email_nil"
	StatusLikeExist              = "like_exist"
	StatusLikeNotExist           = "like_not_exist"
	StatusNoContent              = "no_this_content"
	StatusNoUser                 = "no_this_user"
	StatusNotificationExist      = "notification_exist"
	StatusUserContentNotMatching = "user_content_id_not_matching"
	StatusContentOutofRange      = "content_out_of_range"
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

//CommonRes 通用回应
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

// Contentres 正文信息
type Contentres struct {
	ID          string   `bson:"_id"`
	Detail      string   `bson:"detail"`      // 详情介绍
	OwnID       string   `bson:"ownId"`       // 作者ID [索引]
	PublishDate int64    `bson:"publishDate"` // 发布日期
	LikeNum     int64    `bson:"likeNum"`     // 点赞数
	Public      bool     `bson:"public"`      // 是否公开
	Tag         []string `bson:"tag"`         // 标签
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
	Data  Contentres      //文档数据
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
	Data  []Contentres
}

//like

//Like 点赞信息
type Like struct {
	ID        bson.ObjectId `bson:"_id"`
	UserID    bson.ObjectId `bson:"userId"`    // 用户ID
	ContentID bson.ObjectId `bson:"contentId"` // 内容ID
}

//notification

//Notificationres 通知集合
type Notificationres struct {
	Notification NotificationDetail
	SourceInfo  ContentUserInfo
}

//UserNotificationres 响应用户通知集合
type UserNotificationres struct {
	State         string
	Notifications []Notificationres
}

//NotificationDetail 通知详情
type NotificationDetail struct {
	ID         bson.ObjectId `bson:"_id"`
	CreateTime int64         `bson:"time"`
	Content    string        `bson:"content"`   // 通知内容
	SourceID   bson.ObjectId `bson:"sourceId"`  // 源ID （点赞用户）
	TargetID   bson.ObjectId `bson:"targetId"`  // 目标ID （被点赞用户）
	ContentID  bson.ObjectId `bson:"contentId"` //点赞文章ID
	Type       string        `bson:"type"`      // 类型：暂时只有like
}

//RootRes 简单 API 服务列表
type RootRes struct {
	UserGetInfo      string `json:"user_information_url"`
	UserPostLogin    string `json:"user_login_url"`
	UserPostRegister string `json:"user_register_url"`
	UserPostLogout   string `json:"user_logout_url"`
	UserPostName     string `json:"user_rename_url"`

	ContentDeleteBy    string `json:"content_url"`
	ContentGetDetailBy string `json:"content_detail_url"`
	ContentGetPublic   string `json:"content_public_url"`
	ContentGetTexts    string `json:"content_text_url"`
	ContentPostUpdate  string `json:"content_upadate_url"`

	Like string `json:"like_url"`

	Notification       string `json:"notification_url"`
	NotificationGetAll string `json:"notification_all_url"`
}
