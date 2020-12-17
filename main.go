/*
 * @Descripttion:
 * @version:
 * @Author: sunylin
 * @Date: 2020-12-15 22:38:08
 * @LastEditors: sunylin
 * @LastEditTime: 2020-12-16 22:47:35
 */
package main

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"github.com/service-computing-project/project_app/controllers"
	"github.com/service-computing-project/project_app/models"
	"github.com/service-computing-project/project_app/service"
)

func main() {
	//sessions,err :=service.DBservice()
	sesson, err := service.DBservice()
	if err != nil {
		fmt.Println("error")
	}
	//创建数据库
	var user models.UserDB
	user.DB = sesson.DB("project").C("user")
	app := iris.Default()
	app.Use(myMiddleware)

	sessionID := "mySession"
	//session的创建
	sess := sessions.New(sessions.Config{
		Cookie: sessionID,
	})
	users := mvc.New(app.Party("/user"))
	users.Register(sess.Start)
	users.Handle(&controllers.UsersController{Model: user})

	// create Content model
	var content models.ContentDB
	content.DB = sesson.DB("project").C("content")
	content.DBuser = sesson.DB("project").C("user")
	contents := mvc.New(app.Party("/content"))
	contents.Register(sess.Start)
	contents.Handle(&controllers.ContentController{Model: content})
	//create Notification model
	var notification models.NotifiationDB
	notification.DBN = sesson.DB("project").C("notification")
	notification.DBU = sesson.DB("project").C("user")
	notifications := mvc.New(app.Party("/notification"))
	notifications.Register(sess.Start)
	notifications.Handle(&controllers.NotificationController{Model: notification})
	// Listens and serves incoming http requests
	// on http://localhost:8080.
	app.Listen("0.0.0.0:8080")
}

func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}
