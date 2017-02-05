package main

import (
	"encoding/json"
	"fmt"
	"gocheese/apis"
	"gocheese/db"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

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
		t.Log("pass")
	} else if err != nil {
		t.Error(err)
	} else {
		t.Error("帖子数量为", len(data.Todos))
		fmt.Println("code: ", res.StatusCode)
		fmt.Println("todos len: ", len(data.Todos))
	}
}

func TestCreateTodo(t *testing.T) {
	res, err := http.PostForm(fmt.Sprintf("%s/todos", server.URL), url.Values{"content": {"新任务"}})
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var data apis.ResponseBody
	err = json.Unmarshal(body, &data)
	fmt.Println("body:", data)
	id := data.Content["id"]
	todo := db.FindTodoById(fmt.Sprint(id))
	if res.StatusCode == 200 && err == nil && todo.Content == "新任务" {
		t.Log("通过")
	} else {
		t.Log(res.StatusCode)
		t.Error(err)
	}
}

func TestDeleteTodo(t *testing.T) {
	var todo db.Todo
	db.TodoColl().Find(bson.M{}).One(&todo)
	var client = &http.Client{}
	id := todo.Id.Hex()
	fmt.Println("get id==>", id)
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/todos/%s", server.URL, id), nil)
	res, err := client.Do(req)
	if res.StatusCode == 200 && err == nil {
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
