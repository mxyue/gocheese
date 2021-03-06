package apis

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/dgrijalva/jwt-go"
	"gocheese/db"
	"gocheese/util"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func createSession(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	log.Debug(fmt.Sprintf("email: %s, password: %s", email, password))
	user, found := db.FindUser(bson.M{"email": email})
	if found {
		if user.ValidPassword(password) {
			log.Debug("db user id: ", user.Id.Hex())
			jwtToken := jwt.MapClaims{"user_id": user.Id.Hex()}
			content := map[string]interface{}{"token": util.Encrypt(jwtToken), email: email}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(ResponseBody{200, "登陆成功", content})
		} else {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ResponseBody{601, "密码错误", nil})
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseBody{601, "账号不存在", nil})
	}
}
