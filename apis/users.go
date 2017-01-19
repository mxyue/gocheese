package apis

import (
	"encoding/json"
	"fmt"
	"gocheese/db"
	"net/http"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	user := db.User{Email: email}
	_, err := db.CreateUser(user, password)
	if err != nil {
		errStr := fmt.Sprintf("%s", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseBody{601, errStr, nil})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ResponseBody{200, "保存成功", nil})
	}
}
