/*
 * @Descripttion:user的数据库调用
 * @version:1.0
 * @Author: sunylin
 * @Date: 2020-12-15 02:41:11
 * @LastEditors: sunylin
 * @LastEditTime: 2020-12-22 10:18:36
 */
package models

import (
	"errors"
	"regexp"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//UserDB user数据库
type UserDB struct {
	DB *mgo.Collection
}

//AddUser 添加用户,返回新用户ID和错误信息
func (m *UserDB) AddUser(email, pwd, name, avatar, bio string, gender int) (newUser bson.ObjectId, err error) {
	if len(name) == 0 {
		return "", errors.New(StatusEmptyName)
	}
	if len(email) == 0 {
		return "", errors.New(StatusEmptyEmail)
	}
	if !verifyEmailFormat(email) {
		return "", errors.New(StatusEmailFormatError)
	}
	Validname, err := m.validName(name)
	if err != nil {
		return "", err
	}
	if !Validname {
		return "", errors.New(StatusUserNameExist)
	}
	Validemail, err := m.validEmail(email)
	if err != nil {
		return "", err
	}
	if !Validemail {
		return "", errors.New(StatusEmailExist)
	}
	newUser = bson.NewObjectId()
	err = m.DB.Insert(&User{
		ID:    newUser,
		Pwd:   pwd,
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

//SetUserInfo 设置用户信息
func (m *UserDB) SetUserInfo(id string, info UserInfo) error {
	if !bson.IsObjectIdHex(id) {
		return errors.New(StatusNoID)
	}
	return m.DB.UpdateId(bson.ObjectIdHex(id), bson.M{"$set": bson.M{"info": info}})
}

//validName 用于检测用户名是否重复
func (m *UserDB) validName(username string) (bool, error) {
	count, err := m.DB.Find(bson.M{"info.name": username}).Count()
	if count == 0 {
		return true, err
	}
	return false, err
}

//validEmail 用于检测邮箱是否重复
func (m *UserDB) validEmail(email string) (bool, error) {
	count, err := m.DB.Find(bson.M{"email": email}).Count()
	if count == 0 {
		return true, err
	}
	return false, err
}

//verifyEmailFormat 判断邮箱合法性
func verifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// SetUserName 设置用户名
func (m *UserDB) SetUserName(id, name string) error {
	if !bson.IsObjectIdHex(id) {
		return errors.New(StatusNoID)
	}
	isValid, err := m.validName(name)
	if err != nil {
		return err
	}
	if isValid {
		return m.DB.UpdateId(bson.ObjectIdHex(id), bson.M{"$set": bson.M{"info.name": name}})
	}
	return errors.New(StatusUserNameExist)
}

// GetUserByID 根据ID查询用户
func (m *UserDB) GetUserByID(id string) (User, error) {
	var user User
	if !bson.IsObjectIdHex(id) {
		err := errors.New("not_id")
		return user, err
	}
	err := m.DB.FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

//Login 登陆
func (m *UserDB) Login(username, pwd string) (string, error) {
	count, err := m.DB.Find(bson.M{"info.name": username}).Count()
	if err != nil {
		return "", err
	}
	if count == 0 {
		return "", errors.New(StatusUserNameNotExist)
	}
	type password struct {
		Pwd string        `bson:"password"`
		ID  bson.ObjectId `bson:"_id"`
	}
	var p password
	err = m.DB.Find(bson.M{"info.name": username}).Select(bson.M{"password": 1, "_id": 1}).One(&p)
	if err != nil {
		return "", err
	}
	if p.Pwd == pwd {
		return p.ID.Hex(), nil
	}
	err = errors.New(StatusPasswordError)
	return "", err
}

//Register 注册
func (m *UserDB) Register(username, pwd, email string) (err error) {
	_, err = m.AddUser(email, pwd, username, "https://ss0.bdstatic.com/70cFvHSh_Q1YnxGkpoWK1HF6hhy/it/u=1763861859,2626066999&fm=26&gp=0.jpg", "个性签名", 0)
	if err != nil {
		return err
	}
	return nil
}

//GetUserInfo 获取用户信息
func (m *UserDB) GetUserInfo(id string) (res UserInfoRes, err error) {
	user, err := m.GetUserByID(id)
	if err != nil {
		return
	}
	res.ID = id
	res.Email = user.Email
	res.Info = user.Info
	return
}
