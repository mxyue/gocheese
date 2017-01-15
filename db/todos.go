package db

import (
	log "github.com/Sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Todo struct {
	Id        bson.ObjectId `bson:"_id,omitempty"`
	Content   string        `bson:"content"`
	CreatedAt time.Time     `bson:"created_at"`
}

func GetAllTodos() []Todo {
	var todos []Todo
	TodoColl().Find(nil).All(&todos)
	return todos
}

func FindTodo(query bson.M) Todo {
	var todo Todo
	TodoColl().Find(query).One(&todo)
	return todo
}

func FindTodoById(id string) Todo {
	var todo Todo
	bsonObjectID := bson.ObjectIdHex(id)
	TodoColl().FindId(bsonObjectID).One(&todo)
	return todo
}

func CreateTodo(todo Todo) (bson.ObjectId, interface{}) {
	id := bson.NewObjectId()
	log.Info("CreateTodo  id: ", id.Hex())
	todo.Id = id
	todo.CreatedAt = time.Now()
	doc := TodoColl().Insert(todo)
	return id, doc
}

func DeleteTodoById(id string) error {
	log.Info("DeleteTodoById id: ", id)
	objId := bson.ObjectIdHex(id)
	err := TodoColl().RemoveId(objId)
	if err != nil {
		log.Error(err)
	}
	return err
}
