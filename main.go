/*
 * @Descripttion:
 * @version:
 * @Author: sunylin
 * @Date: 2020-12-15 22:38:08
 * @LastEditors: sunylin
 * @LastEditTime: 2020-12-16 01:54:11
 */
package main

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"github.com/yilin0041/project_app/controllers"
	"github.com/yilin0041/project_app/models"
	"github.com/yilin0041/project_app/service"
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

	// Listens and serves incoming http requests
	// on http://localhost:8080.
	app.Listen(":8080")
}

func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}
