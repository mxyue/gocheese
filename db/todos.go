package db

import (
	"errors"
	log "github.com/Sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Todo struct {
	Id        bson.ObjectId `bson:"_id,omitempty" json:"id"`
	UserId    bson.ObjectId `bson:"user_id,omitempty" json:"user_id"`
	Content   string        `bson:"content" json:"content"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
	Dones     []Done
}

type Done struct {
	Id        bson.ObjectId `bson:"_id"`
	CreatedAt time.Time     `bson:"created_at"`
}

func (t *Todo) AddDone(done Done) int {
	t.Dones = append(t.Dones, done)
	return len(t.Dones)
}

func GetAllTodos() []Todo {
	var todos []Todo
	TodoColl().Find(nil).All(&todos)
	return todos
}

func GetUserTodos(user_id bson.ObjectId) []Todo {
	var todos []Todo
	TodoColl().Find(bson.M{"user_id": user_id}).All(&todos)
	return todos
}

func FindTodo(query bson.M) Todo {
	var todo Todo
	TodoColl().Find(query).One(&todo)
	return todo
}

func FindTodoById(id string) (todo Todo, err error) {
	if bson.IsObjectIdHex(id) {
		bsonObjectID := bson.ObjectIdHex(id)
		TodoColl().FindId(bsonObjectID).One(&todo)
	} else {
		err = errors.New("非法id")
	}
	return todo, err
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

func DeleteUserTodoById(user_id bson.ObjectId, id string) error {
	log.Info("DeleteUserTodoById id: ", id)
	objId := bson.ObjectIdHex(id)
	err := TodoColl().Remove(bson.M{"_id": objId, "user_id": user_id})
	if err != nil {
		log.Error(err)
	}
	return err
}
