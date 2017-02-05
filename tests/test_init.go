package main

import (
	"fmt"
	"gocheese/apis"
	"gocheese/db"
	"gopkg.in/mgo.v2/bson"
	"net/http/httptest"
)

var server *httptest.Server
var firstUser db.User

func init() {
	db.SetDBName("gocheese_test")
	db.UserColl().RemoveAll(nil)
	db.TodoColl().RemoveAll(nil)
	todo := db.Todo{Content: "第一个任务"}
	err := db.TodoColl().Insert(todo)
	password := "123456"
	userData := db.User{Email: "basic@126.com", Mobile: "18280196887", Password: []byte(password)}
	_, err = db.CreateUser(userData, password)
	if err != nil {
		fmt.Println("数据存储不成功:", err)
	}
	firstUser = db.FindUser(bson.M{"email": userData.Email})
	server = httptest.NewServer(apis.Handlers())
}
