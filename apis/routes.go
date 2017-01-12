package apis

import (
	"github.com/gorilla/mux"
)

func Handlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/todos", getTodos).Methods("GET")
	r.HandleFunc("/todos", createTodos).Methods("POST")
	return r
}

type ResponseBody struct {
	Code int
	Msg  string
}
