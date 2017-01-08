package apis

import (
	"github.com/gorilla/mux"
)

func Handlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/todos", getTodos).Methods("GET")
	return r
}
