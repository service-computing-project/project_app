/*
 * @Descripttion:
 * @version:
 * @Author: sunylin
 * @Date: 2020-12-15 22:38:08
 * @LastEditors: sunylin
 * @LastEditTime: 2020-12-21 14:35:23
 */
package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"github.com/service-computing-project/project_app/controllers"
	"github.com/service-computing-project/project_app/models"
	"github.com/service-computing-project/project_app/service"
)

// JWT验证中间件
func ValidateJwtMiddleware(ctx iris.Context) {
	token, err := request.ParseFromRequest(ctx.Request(), request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (i interface{}, e error) {
			return []byte("My Secret"), nil
		})

	if err != nil || !token.Valid {
		var errres models.CommonRes
		errres.State = models.StatusBadReq
		errres.Data = err.Error()
		ctx.JSON(errres)
	} else {
		ctx.Next()
	}
}

func main() {
	//sessions,err :=service.DBservice()
	sesson, err := service.DBservice()
	if err != nil {
		fmt.Println("error")
	}
	//创建数据库
	var user models.UserDB
	user.DB = sesson.DB("project").C("user")
	var like models.LikeDB
	like.DBU = sesson.DB("project").C("user")
	like.DBL = sesson.DB("project").C("like")
	like.DBC = sesson.DB("project").C("content")
	like.DBN = sesson.DB("project").C("notification")
	var content models.ContentDB
	content.DB = sesson.DB("project").C("content")
	content.DBuser = sesson.DB("project").C("user")
	var notification models.NotifiationDB
	notification.DBN = sesson.DB("project").C("notification")
	notification.DBU = sesson.DB("project").C("user")

	app := iris.New()
	app.Use(myMiddleware)
	//app.Use(Cors)
	app.Handle("GET", "/api", func(ctx iris.Context) {
		ctx.JSON(models.RootRes{
			"http://47.103.210.109:8080/api/user/info/{userID}",
			"http://47.103.210.109:8080/api/user/login",
			"http://47.103.210.109:8080/api/user/register",
			"http://47.103.210.109:8080/api/user/logout",
			"http://47.103.210.109:8080/api/user/name",

			"http://47.103.210.109:8080/api/content/{contentID}",
			"http://47.103.210.109:8080/api/content/detail/{contentID}",
			"http://47.103.210.109:8080/api/content/public",
			"http://47.103.210.109:8080/api/content/texts/{userID}",
			"http://47.103.210.109:8080/api/content/update",

			"http://47.103.210.109:8080/api/like/{contentID}",

			"http://47.103.210.109:8080/api/notificaiton/{notificationID}",
			"http://47.103.210.109:8080/api/notification/all",
		})
	})

	sessionID := "mySession"
	//session的创建
	sess := sessions.New(sessions.Config{
		Cookie: sessionID,
		//DisableSubdomainPersistence: true,
	})
	app.Use(sess.Handler())
	crs := cors.New(cors.Options{
		//AllowedOrigins:   []string{"*"}, //允许通过的主机名称
		//AllowCredentials: true,
	})
	users := mvc.New(app.Party("/api/user", crs).AllowMethods())
	users.Register(sess.Start)
	users.Handle(&controllers.UsersController{Model: user})

	likes := mvc.New(app.Party("/api/like", crs).AllowMethods())
	likes.Register(sess.Start)
	likes.Handle(&controllers.LikeController{Model: like})

	contents := mvc.New(app.Party("/api/content", crs).AllowMethods())
	contents.Register(sess.Start)
	contents.Handle(&controllers.ContentController{Model: content})

	notifications := mvc.New(app.Party("/api/notification", crs).AllowMethods())
	notifications.Register(sess.Start)
	notifications.Handle(&controllers.NotificationController{Model: notification})

	// Listens and serves incoming http requests
	// on http://localhost:8080.
	app.Listen("0.0.0.0:8080")
}

func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	fmt.Println("test for middle")
	//ctx.Recorder().ResetHeaders()
	//ctx.Header("Access-Control-Allow-Origin", "*")
	//ctx.Header("Access-Control-Allow-Headers", "content-type")
	ctx.Header("Access-Control-Allow-Credentials","true");
	fmt.Println("Method", ctx.Request().Method)
	if ctx.Request().Method == "OPTIONS" {
		fmt.Println("test for core")
		ctx.StatusCode(200)	
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
		ctx.Header("Access-Control-Allow-Origin", "*")
		return
	}
	if ctx.Request().Method != "GET" && ctx.Request().Method != "POST"  {
		ctx.Header("Access-Control-Allow-Origin", "*")
	}
	ctx.Next()
}
