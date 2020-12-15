package main

import (
	"github.com/kataras/iris/v12"
	"github.com/yilin0041/project_app/controllers"
)

func main() {
    app := iris.Default()
	app.Use(myMiddleware)
	
	users := app.New(app.Party("/user"))
	users.Handle(new(controllers.UsersController))
	
    // Listens and serves incoming http requests
    // on http://localhost:8080.
    app.Listen(":8080")
}

func myMiddleware(ctx iris.Context) {
    ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
    ctx.Next()
}