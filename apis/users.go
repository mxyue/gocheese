package apis

import (
	"encoding/json"
	"errors"
	log "github.com/Sirupsen/logrus"
	"github.com/dgrijalva/jwt-go"
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
	log.Info("接受的email: ", email)
	log.Info("缓存的验证码：", cacheCode)
	log.Info("接收的验证码：", code)
	if found && code == cacheCode {
		user := db.User{Email: email}
		util.CacheDelete(config.RegistSPACE, email)
		user_id, err := db.CreateUser(user, password)
		if err == nil {
			jwtToken := jwt.MapClaims{"user_id": user_id}
			content := map[string]interface{}{"token": util.Encrypt(jwtToken)}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(ResponseBody{200, "登陆成功", content})
			return
		}
	} else {
		log.Info("验证码不对")
		err = errors.New("验证码不正确")
	}
	simpleResponse(w, err, "保存成功")
}

func sendCodeToEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	var err error
	log.Debug("api sendCodeToEmail email:", email)
	if !util.EmailRegex(email) {
		err = errors.New("不符合规则的邮箱")
	} else if _, found := db.FindUser(bson.M{"email": email}); found {
		err = errors.New("该邮箱已经注册")
	} else {
		err = util.SendCode(email, config.RegistSPACE)
	}
	simpleResponse(w, err, "发送成功")
}
