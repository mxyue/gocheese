package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

const (
	DB_NAME = "gocheese"
	TODOS   = "todos"
)

var record = func() *mgo.Database {
	url := "mongodb://localhost"
	session, err := mgo.Dial(url)
	if err != nil {
		fmt.Println(err)
	}
	return session.DB(DB_NAME)
}

var TodoColl = func() *mgo.Collection {
	return record().C(TODOS)
}
