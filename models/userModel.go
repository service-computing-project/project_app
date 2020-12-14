package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//UserModel user数据库
type UserModel struct {
	DB *mgo.Collection
}

//AddUser 添加用户
func (m *UserModel) AddUser(email, name, avatar, bio string, gender int) (newUser bson.ObjectId, err error) {
	newUser = bson.NewObjectId()
	err = m.DB.Insert(&User{
		ID:    newUser,
		Email: email,
		Info: UserInfo{
			Name:   name,
			Avatar: avatar,
			Bio:    bio,
			Gender: gender,
		},
	})
	return newUser, err
}
