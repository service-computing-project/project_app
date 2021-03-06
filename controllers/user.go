package controllers

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/globalsign/mgo/bson"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/service-computing-project/project_app/models"
)

type UsersController struct {
	Ctx     iris.Context
	Model   models.UserDB
	Session *sessions.Session
}

// RegisterReq POST /user/register 注册请求
type RegisterReq struct {
	Name     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//OptionsRegister OPTIONS /user/register 用户注册
func (c *UsersController) OptionsRegister() {
	return
}

//PostRegister POST /user/register 用户注册
func (c *UsersController) PostRegister() (res models.CommonRes) {
	req := RegisterReq{}
	if err := c.Ctx.ReadJSON(&req); err != nil {
		res.State = models.StatusBadReq
	}
	if err := c.Model.Register(req.Name, req.Password, req.Email); err != nil {
		res.State = err.Error()
	} else {
		//c.Session.Set("name", req.Name)
		res.State = models.StatusSuccess
	}
	return
}

// LoginReq POST /user/login 登陆请求
type LoginReq struct {
	Name     string `json:"username"`
	Password string `json:"password"`
}

//OptionsLogin Options /user/login 用户登陆
func (c *UsersController) OptionsLogin() {
	return
}

//PostLogin POST /user/login 用户登陆
func (c *UsersController) PostLogin() (res models.CommonRes) {
	req := LoginReq{}
	err := c.Ctx.ReadJSON(&req)
	if err != nil || req.Name == "" || req.Password == "" {
		res.State = models.StatusBadReq
		return
	}
	userID, err := c.Model.Login(req.Name, req.Password)
	if err != nil {
		res.State = err.Error()
	} else {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"name":     req.Name,
			"password": req.Password,
			"exp":      time.Now().Add(time.Hour * 1).Unix(),
		})

		// 这里的密钥和前面的必须一样
		tokenString, _ := token.SignedString([]byte("My Secret"))
		res.Data = tokenString
		c.Session.Set("id", userID)
		fmt.Println(c.Session)
		res.State = models.StatusSuccess
	}
	return
}

//OptionsLogout Options user/logout 退出登陆
func (c *UsersController) OptionsLogout() {
	return
}

//PostLogout POST /user/logout 退出登陆
func (c *UsersController) PostLogout() (res models.CommonRes) {
	fmt.Println(c.Session)
	if c.Session.Get("id") == nil {
		res.State = models.StatusNotLogin
		return
	}
	c.Session.Delete("id")
	res.State = models.StatusSuccess
	return
}

//NameReq POST /user/name 更新用户名
type NameReq struct {
	Name string `json:"name"`
}

//OptionsName OPTIONS /user/name
func (c *UsersController) OptionsName() {
	return
}

//PostName POST /user/name 更新用户名 (Token required)
func (c *UsersController) PostName() (res models.CommonRes) {
	if c.Session.Get("id") == nil {
		res.State = models.StatusNotLogin
		return
	}
	token, err := request.ParseFromRequest(c.Ctx.Request(), request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (i interface{}, e error) {
			return []byte("My Secret"), nil
		})

	if err != nil || !token.Valid {
		res.Data = err.Error()
		res.State = models.StatusBadReq
		return
	}
	req := NameReq{}
	err = c.Ctx.ReadJSON(&req)
	if err != nil || req.Name == "" {
		res.State = models.StatusBadReq
		return
	}
	err = c.Model.SetUserName(c.Session.GetString("id"), req.Name)
	if err != nil {
		res.State = err.Error()
	} else {
		res.State = models.StatusSuccess
	}
	return
}

//OptionsInfoBy OPTIONS /user/info/{userID:string}
func (c *UsersController) OptionsInfoBy(id string) {
	return
}

//GetInfo GET /user/info/{userID:string}
func (c *UsersController) GetInfoBy(id string) (res models.UserInfoRes) {
	if id == "self" {
		if c.Session.Get("id") == nil {
			res.State = models.StatusNotLogin
			return
		}
		id = c.Session.GetString("id")
	} else if !bson.IsObjectIdHex(id) {
		res.State = models.StatusBadReq
		return
	}
	userinfores, err := c.Model.GetUserInfo(id)
	res = userinfores
	if err != nil {
		res.State = err.Error()
	} else {
		res.State = models.StatusSuccess
	}
	return
}
