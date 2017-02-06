package apis

import (
	"github.com/gorilla/mux"
	// "gocheese/db"
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
	// r.HandleFunc("/todos", util.Validate(getTodos)).Methods("GET")
	// r.HandleFunc("/todos", UserValid(getTodos)).Methods("GET")
	r.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		user, err := util.ValidUser(w, r)
		if err == nil {
			getTodos(w, r, user)
		}
	}).Methods("GET")
	r.HandleFunc("/todos", createTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", deleteTodo).Methods("DELETE")

	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/sessions", createSession).Methods("POST")
	return r
}
