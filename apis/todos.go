package apis

import (
	"encoding/json"
	"gocheese/db"
	"net/http"
	"time"
)

type Body struct {
	Todos []db.Todo
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	todos := db.GetAllTodos()
	body := Body{todos}
	json.NewEncoder(w).Encode(body)
}

type TodoContent struct {
	Content string `json:"content"`
}

func createTodos(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	content := r.Form.Get("content")
	if len(content) <= 1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseBody{600, "请填写内容"})
		return
	}
	todo := db.Todo{content, time.Now()}
	err := db.CreateTodo(todo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseBody{601, "保存失败"})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ResponseBody{200, "保存成功"})
	}
}
