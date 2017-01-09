package db

import "time"

type Todo struct {
	Content   string
	CreatedAt time.Time
}

func GetAllTodos() []Todo {
	var todos []Todo
	TodoColl().Find(nil).All(&todos)
	return todos
}

func CreateTodo(todo Todo) error {
	err := TodoColl().Insert(todo)
	return err
}
