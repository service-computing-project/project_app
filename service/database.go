package service

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

//DBservice :connect DB
func DBservice() (*mgo.Session, error) {
	session, err := mgo.Dial("mongodb://47.103.210.109:27017")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)

	return session, nil
}
