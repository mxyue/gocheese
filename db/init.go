package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

var db_name = "gocheese"

const (
	TODOS = "todos"
	USERS = "users"
)

func SetDBName(new_name string) {
	db_name = new_name
}

var record = func() *mgo.Database {
	url := "mongodb://localhost"
	session, err := mgo.Dial(url)
	if err != nil {
		fmt.Println(err)
	}
	return session.DB(db_name)
}

var TodoColl = func() *mgo.Collection {
	return record().C(TODOS)
}

var UserColl = func() *mgo.Collection {
	return record().C(USERS)
}
