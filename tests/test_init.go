package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/dgrijalva/jwt-go"
	"gocheese/apis"
	"gocheese/db"
	"gocheese/util"
	"gopkg.in/mgo.v2/bson"
	"net/http/httptest"
	"os"
)

var server *httptest.Server
var firstUser db.User
var firstUserToken string

func init() {
	db.SetDBName("gocheese_test")

	db.UserColl().RemoveAll(nil)
	db.TodoColl().RemoveAll(nil)

	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)

	password := "123456"
	userData := db.User{Email: "basic@126.com", Mobile: "18280196887", Password: []byte(password)}
	_, err := db.CreateUser(userData, password)

	firstUser, _ = db.FindUser(bson.M{"email": userData.Email})
	mapClaims := jwt.MapClaims{"user_id": firstUser.Id.Hex()}
	firstUserToken = util.Encrypt(mapClaims)

	todo := db.Todo{Content: "第一个任务", UserId: firstUser.Id}
	err = db.TodoColl().Insert(todo)
	if err != nil {
		fmt.Println("数据存储不成功:", err)
	}
	server = httptest.NewServer(apis.Handlers())
}
