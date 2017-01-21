package main

import (
	"fmt"
	"gocheese/apis"
	"gocheese/db"
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
	firstUser = db.User{Email: "basic@126.com", Mobile: "18280196887", Password: []byte(password)}
	_, err = db.CreateUser(firstUser, password)
	if err != nil {
		fmt.Println("数据存储不成功:", err)
	}
	server = httptest.NewServer(apis.Handlers())
}
