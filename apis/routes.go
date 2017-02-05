package apis

import (
	"github.com/gorilla/mux"
	"gocheese/middleware"
)

type ResponseBody struct {
	Code    int
	Msg     string
	Content map[string]interface{}
}

func Handlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/todos", middleware.Validate(getTodos)).Methods("GET")
	r.HandleFunc("/todos", createTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", deleteTodo).Methods("DELETE")

	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/sessions", createSession).Methods("POST")
	return r
}
