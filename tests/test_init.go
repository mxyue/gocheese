package main

import (
	"fmt"
	"gocheese/apis"
	"gocheese/db"
	"net/http/httptest"
)

var server *httptest.Server

func init() {
	db.SetDBName("gocheese_test")
	db.UserColl().RemoveAll(nil)
	db.TodoColl().RemoveAll(nil)
	todo := db.Todo{Content: "第一个任务"}
	err := db.TodoColl().Insert(todo)
	if err != nil {
		fmt.Println("数据存储不成功:", err)
	}
	server = httptest.NewServer(apis.Handlers())
}
