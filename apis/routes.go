package apis

import (
	// log "github.com/Sirupsen/logrus"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gocheese/db"
	"gocheese/util"
	"net/http"
)

type ResponseBody struct {
	Code    int
	Msg     string
	Content map[string]interface{}
}

func Handlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/todos", userValid(getTodos)).Methods("GET")
	r.HandleFunc("/todos", userValid(createTodo)).Methods("POST")
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
