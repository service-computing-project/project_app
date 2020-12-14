package main

import (
	"fmt"

	"github.com/yilin0041/project_app/models"
	"github.com/yilin0041/project_app/service"
)

func main() {
	sesson, err := service.DBservice()
	if err != nil {
		fmt.Println("error")
	}
	var user models.UserModel
	user.DB = sesson.DB("project").C("user")
	user.AddUser("10@12.com", "test", "", "test bio", 1)
}
