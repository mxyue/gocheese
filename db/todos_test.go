package db

import (
	"fmt"
	"testing"
)

var TestMode bool

func init() {
	db_name = "gocheese_test"
	TodoColl().RemoveAll(nil)
	todo := Todo{Content: "第一个任务"}
	err := TodoColl().Insert(todo)
	if err != nil {
		fmt.Println("数据存储不成功")
	}
}

func TestCreateTodo(t *testing.T) {
	todo := Todo{Content: "第一个任务"}
	_, err := CreateTodo(todo)
	if err == nil {
		t.Log("成功")
	} else {
		t.Error("失败")
	}
}

func TestGetAllTodos(t *testing.T) {
	todos := GetAllTodos()
	if len(todos) != 2 {
		t.Error("数据查询不对，todos长度为：", len(todos))
	} else {
		t.Log("成功")
	}

}
