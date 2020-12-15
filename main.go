package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/yilin0041/project_app/controllers"
	"github.com/yilin0041/project_app/service"
)

func main() {
	//sessions,err :=service.DBservice()
	
    app := iris.Default()
	app.Use(myMiddleware)
	
	users := mvc.New(app.Party("/user"))
	users.Handle(new(controllers.UsersController))
	
    // Listens and serves incoming http requests
    // on http://localhost:8080.
    app.Listen(":8080")
}


func myMiddleware(ctx iris.Context) {
    ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
    ctx.Next()
}