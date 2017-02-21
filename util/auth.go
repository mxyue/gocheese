package util

import (
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"gocheese/db"
	"net/http"
)

func ValidUser(w http.ResponseWriter, r *http.Request) (db.User, error) {
	token := r.Header.Get("token")
	log.Debug("ValidUser token: ", token)
	claims, err := Decrypt(token)
	log.Debug("ValidUser user_id: ", claims["user_id"])
	if err != nil || claims["user_id"] == nil {
		log.Info("====验证不过===")
		w.WriteHeader(http.StatusUnauthorized)
		return db.User{}, errors.New("not auth")
	} else {
		userId := fmt.Sprintf("%s", claims["user_id"])
		user, err := db.FindUserById(userId)
		if err != nil {
			log.Info("====验证不过2===")
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			log.Info("====验证通过===")
		}
		return user, err
	}
}
