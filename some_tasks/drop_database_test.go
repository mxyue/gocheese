package main

import (
	"fmt"
	"gocheese/db"
	"testing"
)

func TestDropDatabase(t *testing.T) {
	fmt.Println("清空数据库")
	db.TodoColl().RemoveAll(nil)
	t.Log("删除完成")
}
