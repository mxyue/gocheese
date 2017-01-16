package db

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"testing"
	"time"
)

var todo Todo

func init() {
	db_name = "gocheese_test"
	log.SetLevel(log.DebugLevel)
	TodoColl().RemoveAll(nil)
	todo = Todo{Content: "第一个任务"}
	err := TodoColl().Insert(todo)
	if err != nil {
		fmt.Println("数据存储不成功")
	}
}

func TestBcrypt(t *testing.T) {
	password := []byte("MyDarkSecret")
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	log.Debug(string(hashedPassword))
	err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		t.Error(err)
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

func TestAddDone(t *testing.T) {
	done := Done{CreatedAt: time.Now()}
	count := todo.AddDone(done)
	if count == 1 {
		t.Log("成功")
	} else {
		t.Error("失败done长度：", len(todo.Dones))
	}
}
