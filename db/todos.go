package db

import "time"

type Todo struct {
	content    string
	created_at time.Time
	finish_at  time.Time
	do_num     int
}

func GetAllTodos() Todo {
	var todo Todo
	TodoColl().Find(nil).All(&todo)
	return todo
}
