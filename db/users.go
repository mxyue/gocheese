package db

import (
	"errors"
	log "github.com/Sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	Id        bson.ObjectId `bson:"_id,omitempty"`
	Email     string        `bson:"email"`
	Mobile    string        `bson:"mobile"`
	Password  []byte        `bson:"password"`
	CreatedAt time.Time     `bson:"created_at"`
}

func CreateUser(user User, password string) (interface{}, error) {
	dbEmailUser := FindUser(bson.M{"email": user.Email})
	if dbEmailUser.Email != "" {
		return nil, errors.New("该邮箱已经注册")
	}
	dbMobileUser := FindUser(bson.M{"mobile": user.Mobile})
	if dbMobileUser.Mobile != "" {
		return nil, errors.New("该手机已经注册")
	}
	id := bson.NewObjectId()
	user.Id = id
	user.CreatedAt = time.Now()
	bt_password := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bt_password, bcrypt.DefaultCost)
	if err != nil {
		log.Error("create user error", err)
	}
	user.Password = hashedPassword
	doc := UserColl().Insert(user)
	return doc, nil
}

func GetAllUsers() []User {
	var users []User
	UserColl().Find(nil).All(&users)
	return users
}

func FindUser(query bson.M) User {
	var user User
	UserColl().Find(query).One(&user)
	return user
}

func (user *User) ValidPassword(password string) bool {
	log.Debug(string(user.Password))
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		log.Info("db ValidPassword error:", err)
		return false
	} else {
		return true
	}
}
