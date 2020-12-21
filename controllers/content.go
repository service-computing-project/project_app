package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/service-computing-project/project_app/models"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

//ContentController content 控制器
type ContentController struct {
	Ctx     iris.Context
	Model   models.ContentDB
	Session *sessions.Session
}

//GetDetailBy GET /api/content/detail/{contentID:string} 获取指定内容
func (c *ContentController) GetDetailBy(contentID string) (res models.ContentDetailres) {
	if c.Session.Get("id") == nil {
		res.State = models.StatusNotLogin
	}
	if contentID == "" {
		res.State = models.StatusBadReq
		return
	}
	contentdetailres, err := c.Model.GetDetailByID(contentID)
	res = contentdetailres
	if err != nil {
		res.State = err.Error()
	} else {
		res.State = models.StatusSuccess
	}
	return
}

//Options options
func (c *ContentController) Options() {
	return
}

//OptionsBy options
func (c *ContentController) OptionsBy(contentID string) {
	return
}

//DeleteBy DELETE /api/content/{contentID:string}  删除指定内容
func (c *ContentController) DeleteBy(contentID string) (res models.CommonRes) {
	if c.Session.Get("id") == nil {
		res.State = models.StatusNotLogin
	}
	if contentID == "" {
		res.State = models.StatusBadReq
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
	err = c.Model.RemoveContent(contentID)
	if err != nil {
		res.State = err.Error()
	} else {
		res.State = models.StatusSuccess
	}
	return
}

type PageParams struct {
	Page    int `url:"page"`
	PerPage int `url:"per_page"`
}

//GetPublic GET /api/content/public  获取公共内容
func (c *ContentController) GetPublic() (res models.ContentPublicList) {
	if c.Session.Get("id") == nil {
		res.State = models.StatusNotLogin
	}
	params := PageParams{}
	err := c.Ctx.ReadQuery(&params)
	if err != nil && !iris.IsErrPath(err) {
		res.State = models.StatusBadReq
		return
	}
	if params.Page < 1 || params.PerPage < 1 {
		res.State = models.StatusBadReq
		return
	}
	contentpublicres, err := c.Model.GetPublic(params.Page, params.PerPage)
	res = contentpublicres
	if err != nil {
		res.State = err.Error()
	} else {
		res.State = models.StatusSuccess
	}
	return
}

//GetUsercontentBy GET /api/content/usercontent/{userID:string} 获取指定用户的所有内容
func (c *ContentController) GetUsercontentBy(userID string) (res models.ContentListByUser) {
	var contentlistbyuserres models.ContentListByUser
	var err error
	params := PageParams{}
	err = c.Ctx.ReadQuery(&params)
	if err != nil && !iris.IsErrPath(err) {
		res.State = models.StatusBadReq
		return
	}
	if params.Page < 1 || params.PerPage < 1 {
		res.State = models.StatusBadReq
		return
	}
	if userID == "" {
		res.State = models.StatusBadReq
		return
	} else if userID == "self" {
		if c.Session.Get("id") == nil {
			res.State = models.StatusNotLogin
			return
		}
		userID = c.Session.GetString("id")
		contentlistbyuserres, err = c.Model.GetContentSelf(userID, params.Page, params.PerPage)
	} else {
		contentlistbyuserres, err = c.Model.GetContentByUser(userID, params.Page, params.PerPage)
	}

	res = contentlistbyuserres
	if err != nil {
		res.State = err.Error()
	} else {
		res.State = models.StatusSuccess
	}
	return
}

//TextReq POST /api/content/text 增加文本内容
type TextReq struct {
	Detail   string   `json:"detail"`
	Tags     []string `json:"tags"`
	IsPublic bool     `json:"isPublic"`
}

//OptionsText  OPTIONS /api/content/text 增加文本内容
func (c *ContentController) OptionsText() {
	return
}

//PostText  POST /api/content/text 增加文本内容
func (c *ContentController) PostText() (res models.CommonRes) {
	id := c.Session.Get("id")
	if id == nil {
		res.State = models.StatusNotLogin
		return
	}
	// //token check
	// token, err := request.ParseFromRequest(c.Ctx.Request(), request.AuthorizationHeaderExtractor,
	// 	func(token *jwt.Token) (i interface{}, e error) {
	// 		return []byte("My Secret"), nil
	// 	})

	// if err != nil || !token.Valid {
	// 	res.Data = err.Error()
	// 	res.State = models.StatusBadReq
	// 	return
	// }
	req := TextReq{}
	err := c.Ctx.ReadJSON(&req)
	if err != nil || req.Detail == "" {
		res.State = models.StatusBadReq
		return
	}
	err1 := c.Model.AddContent(req.Detail, req.Tags, id.(string), req.IsPublic)
	if err1 != nil {
		res.State = err1.Error()
	} else {
		res.State = models.StatusSuccess
	}
	return
}

//TextUpdateReq 文本内容
type TextUpdateReq struct {
	ID       string   `json:"contentID"`
	Detail   string   `json:"detail"`
	Tags     []string `json:"tags"`
	IsPublic bool     `json:"isPublic"`
}

//OptionsUpdate OPTIONS /api/content/update 增加文本内容
func (c *ContentController) OptionsUpdate() {
	return
}

//PostUpdate  POST /api/content/update 增加文本内容
func (c *ContentController) PostUpdate() (res models.CommonRes) {
	id := c.Session.Get("id")
	if id == nil {
		res.State = models.StatusNotLogin
		return
	}
	// //token check
	// token, err := request.ParseFromRequest(c.Ctx.Request(), request.AuthorizationHeaderExtractor,
	// 	func(token *jwt.Token) (i interface{}, e error) {
	// 		return []byte("My Secret"), nil
	// 	})

	// if err != nil || !token.Valid {
	// 	res.Data = err.Error()
	// 	res.State = models.StatusBadReq
	// 	return
	// }
	req := TextUpdateReq{}
	err := c.Ctx.ReadJSON(&req)
	if err != nil || req.Detail == "" {
		res.State = models.StatusBadReq
		return
	}
	err1 := c.Model.UpdateContent(req.ID, req.Detail, req.Tags, id.(string), req.IsPublic)
	if err1 != nil {
		res.State = err1.Error()
	} else {
		res.State = models.StatusSuccess
	}
	return
}
