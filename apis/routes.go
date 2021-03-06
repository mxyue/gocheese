package apis

import (
	 log "github.com/Sirupsen/logrus"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gocheese/db"
	"gocheese/util"
	"net/http"
)

type ResponseBody struct {
	Code    int                    `json:"code"`
	Msg     string                 `json:"msg"`
	Content map[string]interface{} `json:"content"`
}

func Handlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/todos", userValid(getTodos)).Methods("GET")
	r.HandleFunc("/todos", userValid(createTodo)).Methods("POST")
	r.HandleFunc("/todos/{todo_id}/dones", userValid(createTodoDone)).Methods("POST")
	r.HandleFunc("/todos/{id}", userValid(deleteTodo)).Methods("DELETE")

	r.HandleFunc("/valid_email", sendCodeToEmail).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/sessions", createSession).Methods("POST")
	return r
}

func userValid(callHandle func(w http.ResponseWriter, r *http.Request, user db.User)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := util.ValidUser(w, r)
		if err == nil {
			callHandle(w, r, user)
		} else {
			json.NewEncoder(w).Encode(ResponseBody{401, "用户身份验证未通过", nil})
		}
	}
}

func simpleResponse(w http.ResponseWriter, err error, success string) {
	if err != nil {
		errStr := fmt.Sprintf("%s", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseBody{601, errStr, nil})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ResponseBody{200, success, nil})
	}
}

func fullResponse(w http.ResponseWriter, success string, err error, content map[string]interface{}){
	if err != nil {
		log.Debug("fullResponse err: ", err)
		errStr := fmt.Sprintf("%s", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseBody{601, errStr, nil})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ResponseBody{200, success, content})
	}
}

func errResponse(w http.ResponseWriter, err error, errMsg string){
	w.WriteHeader(http.StatusBadRequest)
	content := map[string]interface{}{"err": err}
	json.NewEncoder(w).Encode(ResponseBody{601, errMsg, content})
}