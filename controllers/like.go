package controllers

import (
	"github.com/globalsign/mgo/bson"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/service-computing-project/project_app/models"
)

type LikeController struct {
	Ctx     iris.Context
	Model   models.LikeDB
	Session *sessions.Session
}

// LikeRes 用户点赞数据返回值
type LikeRes struct {
	State string
	Data  []string
}

// GetBy Get /like/{contentID} 获取用户点赞列表
func (c *LikeController) GetBy(id string) (res LikeRes) {
	if !bson.IsObjectIdHex(id) {
		res.State = models.StatusBadReq
		return
	}
	var err error
	res.Data, err = c.Model.GetUserListByContentID(id)
	if err != nil {
		res.State = err.Error()
	}
	res.State = models.StatusSuccess
	return
}

//​ PostBy POST /like/{contentID} 对某个内容点赞
func (c *LikeController) PostBy(id string) (res models.CommonRes) {
	if c.Session.Get("id") == nil {
		res.State = models.StatusNotLogin
		return
	}
	if !bson.IsObjectIdHex(id) {
		res.State = models.StatusBadReq
		return
	}
	err := c.Model.LikeByID(id, c.Session.Get("id").(string))
	if err != nil {
		res.State = err.Error()
	}
	res.State = models.StatusSuccess
	return
}

//​ PatchBy PATCH /like/{contentID} 取消用户对某个内容的点赞
func (c *LikeController) PatchBy(id string) (res models.CommonRes) {
	if c.Session.Get("id") == nil {
		res.State = models.StatusNotLogin
		return
	}
	if !bson.IsObjectIdHex(id) {
		res.State = models.StatusBadReq
		return
	}
	err := c.Model.CancelLikeByID(id, c.Session.Get("id").(string))
	if err != nil {
		res.State = err.Error()
	}
	res.State = models.StatusSuccess
	return
}
