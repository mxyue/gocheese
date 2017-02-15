package apis

import (
	"errors"
	log "github.com/Sirupsen/logrus"
	"gocheese/config"
	"gocheese/db"
	"gocheese/util"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	var err error
	r.ParseForm()
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	code := r.Form.Get("code")
	cacheCode, found := util.CacheGet(config.RegistSPACE, email)
	log.Info("缓存的验证码：", cacheCode)
	log.Info("接收的验证码：", code)
	if found && code == cacheCode {
		user := db.User{Email: email}
		util.CacheDelete(config.RegistSPACE, email)
		_, err = db.CreateUser(user, password)
	} else {
		log.Info("验证码不对")
		err = errors.New("验证码不正确")
	}
	simpleResponse(w, err, "保存成功")
}

func sendCodeToEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	var err error
	if _, found := db.FindUser(bson.M{"email": email}); found {
		err = errors.New("该邮箱已经注册")
	} else {
		err = util.SendCode(email, config.RegistSPACE)
	}
	simpleResponse(w, err, "发送成功")
}
