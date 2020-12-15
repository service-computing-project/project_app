/*
 * @Descripttion:
 * @version:
 * @Author: sunylin
 * @Date: 2020-12-15 02:53:22
 * @LastEditors: sunylin
 * @LastEditTime: 2020-12-15 17:24:16
 */
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
	var user models.UserDB
	user.DB = sesson.DB("project").C("user")
	// myid, err := user.AddUser("10@12.com", "123", "test", "", "test bio", 1)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// user.SetUserName(myid.Hex(), "change2")
	u, err := user.GetUserByID("5fd87ff1c93c7348cc21499a")
	fmt.Println(u)
}
