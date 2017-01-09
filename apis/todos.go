package apis

import (
	"encoding/json"
	"gocheese/db"
	"net/http"
)

type Body struct {
	Todos []db.Todo
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	todos := db.GetAllTodos()
	body := Body{todos}
	json.NewEncoder(w).Encode(body)
}
