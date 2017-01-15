package apis

import (
	"github.com/gorilla/mux"
)

type ResponseBody struct {
	Code    int
	Msg     string
	Content map[string]interface{}
}

func Handlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/todos", getTodos).Methods("GET")
	r.HandleFunc("/todos", createTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", deleteTodo).Methods("DELETE")
	return r
}
