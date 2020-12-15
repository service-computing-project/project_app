package controllers

import(
	"errors"
	"regexp"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"

	"github.com/yilin0041/project_app/models"
	"github.com/yilin0041/project_app/service"
)

type UsersController struct {
	Ctx     iris.Context
	Model *models.UserDB
	Session *sessions.Session
}

// RegisterReq POST /user/register 注册请求
type RegisterReq struct {
	Name     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//PostLogin POST /user/register 用户注册
func (c *UsersController) PostRegister() (res CommonRes){
	req := RegisterReq{}
	if err := c.Ctx.ReadJSON(&req); err != nil {
		res.State = StatusBadReq
	}
	if err := c.Model.Register(req.Name, req.Password, req.Email); err != nil {
		res.State = err.Error()
	} else {
		//c.Session.Set("name", req.Name)
		res.State = StatusSuccess
	}
	return
}

// LoginReq POST /user/login 登陆请求
type LoginReq struct {
	Name     string `json:"username"`
	Password string `json:"password"`
}

//PostLogin POST /user/login 用户登陆
func (c *UsersController) PostLogin() (res CommonRes){
	req := LoginReq{}
	err := c.Ctx.ReadJSON(&req)
	if err != nil || req.Name == "" || req.Password == "" {
		res.State = StatusBadReq
		return
	}
	err := c.Model.Login(req.Name, req.Password)
	if err != nil {
		res.State = err.Error()
	} else {
		c.Session.Set("id", userID)
		res.State = StatusSuccess
	}
	return	
}

//PostLogout POST /user/logout 退出登陆 
func (c *UsersController) PostLogout() (res CommonRes){
	c.Session.Delete("id")
	res.State = StatusSuccess
	return
}

//NameReq POST /user/name 更新用户名
type NameReq struct {
	Name string `json:"name"`
}

//PostName POST /user/name 更新用户名
func (c *UsersController) PostName() (res CommonRes){
	if c.Session.Get("id") == nil {
		res.State = StatusNotLogin
		return
	}
	req := NameReq{}
	err := c.Ctx.ReadJSON(&req)
	if err != nil || req.Name == "" {
		res.State = StatusBadReq
		return
	}
	err := c.Model.SetUserName(c.Session.GetString("id"), req.Name)
	if err != nil {
		res.State = err.Error()
	} else {
		res.State = StatusSuccess
	}
	return	
}

//GetInfo GET /user/info/{userID:string}
func (c *UsersController) GetInfoBy(id string) (res UserInfoRes){
	if id=="self"{
		if c.Session.Get("id") == nil {
			res.State = StatusNotLogin
			return
		}
		id = c.Session.GetString("id")
		
	}
	userinfores, err := c.Model.GetUserInfo(c.Session.GetString("id"))
	res = userinfores
	if err != nil {
		res.State = err.Error()
	} else {
		res.State = StatusSuccess
	}
	return	
}

      