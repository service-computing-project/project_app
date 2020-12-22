/*
 * @Descripttion:
 * @version:
 * @Author: sunylin
 * @Date: 2020-12-20 22:54:43
 * @LastEditors: sunylin
 * @LastEditTime: 2020-12-22 13:07:04
 */
package controllers

import (
	"github.com/service-computing-project/project_app/models"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

type NotificationController struct {
	Ctx     iris.Context
	Model   models.NotifiationDB
	Session *sessions.Session
}

// type PageParams struct {
// 	Page   int   `url:"page"`
// 	PerPage int  `url:"per_page"`
// }
//OptionsAll OPTIONS /api/notification/all  获取用户所有通知
func (c *NotificationController) OptionsAll() (res models.UserNotificationres) {
	return
}

//GetAll GET /api/notification/all  获取用户所有通知
func (c *NotificationController) GetAll() (res models.UserNotificationres) {
	if c.Session.Get("id") == nil {
		res.State = models.StatusNotLogin
		return
	}
	id := c.Session.GetString("id")
	notificationres, err := c.Model.GetNotificationByUserID(id)
	res = notificationres
	if err != nil {
		res.State = err.Error()
	} else {
		res.State = models.StatusSuccess
	}
	return
}

//OptionsBy OPTIONS /api/notificaiton/{NotificationID:string}  删除指定通知
func (c *NotificationController) OptionsBy(notificationID string) {
	return
}

//DeleteBy DELETE /api/notificaiton/{NotificationID:string}  删除指定通知
func (c *NotificationController) DeleteBy(notificationID string) (res models.CommonRes) {
	if c.Session.Get("id") == nil {
		res.State = models.StatusNotLogin
		return
	}
	if notificationID == "" {
		res.State = models.StatusBadReq
	}
	err := c.Model.DeleteNotificationByID(notificationID)
	if err != nil {
		res.State = err.Error()
	} else {
		res.State = models.StatusSuccess
	}
	return

}
