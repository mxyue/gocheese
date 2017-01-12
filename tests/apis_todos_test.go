package main

import (
	"encoding/json"
	"fmt"
	"gocheese/apis"
	"gocheese/db"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

var server *httptest.Server

func init() {
	db.SetDBName("gocheese_test")
	db.TodoColl().RemoveAll(nil)
	todo := db.Todo{"第一个任务", time.Now()}
	err := db.TodoColl().Insert(todo)
	if err != nil {
		fmt.Println("数据存储不成功")
	}
	server = httptest.NewServer(apis.Handlers())
}

func TestGetTodos(t *testing.T) {
	var client = &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/todos", server.URL), nil)
	checkErr(err)
	res, err := client.Do(req)
	checkErr(err)
	fmt.Println("response code: ", res.StatusCode)
	defer res.Body.Close()
	type Data struct {
		Todos []db.Todo
	}
	var data Data
	err = json.NewDecoder(res.Body).Decode(&data)
	checkErr(err)

	if len(data.Todos) == 1 && res.StatusCode == 200 && err == nil {
		fmt.Printf("first Content: %v\n", data.Todos[0].Content)
	} else if err != nil {
		t.Error(err)
	} else {
		t.Error("帖子数量为", len(data.Todos))
		fmt.Println("code: ", res.StatusCode)
		fmt.Println("todos len: ", len(data.Todos))
	}
}

func TestCreateTodo(t *testing.T) {
	// var client = &http.Client{}
	// req, err := http.NewRequest("POST", fmt.Sprintf("%s/todos", server.URL), body)
	res, err := http.PostForm(fmt.Sprintf("%s/todos", server.URL), url.Values{"content": {"新任务"}})
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var data interface{}
	err = json.Unmarshal(body, &data)
	fmt.Println(data)

	if res.StatusCode == 200 && err == nil {
		todos := db.GetAllTodos()
		fmt.Println("Content:", todos[0].Content)
		fmt.Println("Content:", todos[1].Content)
		t.Log("通过")
	} else {
		t.Log(res.StatusCode)
		t.Error(err)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("错误：", err)
	}
}
