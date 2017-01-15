package apis

import (
	"encoding/json"
	// "fmt"
	"github.com/gorilla/mux"
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

type TodoContent struct {
	Content string `json:"content"`
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	content := r.Form.Get("content")
	if len(content) <= 1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseBody{600, "请填写内容", nil})
		return
	}
	todo := db.Todo{Content: content}
	id, err := db.CreateTodo(todo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseBody{601, "保存失败", nil})
	} else {
		w.WriteHeader(http.StatusOK)
		content := map[string]interface{}{"id": id.Hex()}
		json.NewEncoder(w).Encode(ResponseBody{200, "保存成功", content})
	}
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["id"]
	err := db.DeleteTodoById(todoId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		content := map[string]interface{}{"err": err}
		json.NewEncoder(w).Encode(ResponseBody{601, "删除失败", content})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ResponseBody{200, "删除成功", nil})
	}
}
