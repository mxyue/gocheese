package apis

import (
	"encoding/json"
	 log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"gocheese/db"
	"net/http"
	"time"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"errors"
)

type Body struct {
	Todos []db.Todo `json:"todos"`
}

func getTodos(w http.ResponseWriter, r *http.Request, user db.User) {
	todos := db.GetUserTodos(user.Id)
	body := Body{todos}
	json.NewEncoder(w).Encode(body)
}

type TodoContent struct {
	Content string `json:"content"`
}

func createTodo(w http.ResponseWriter, r *http.Request, user db.User) {
	r.ParseForm()
	content := r.Form.Get("content")
	if len(content) <= 1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseBody{600, "内容不能为空", nil})
	} else {
		todo := db.Todo{Content: content, UserId: user.Id}
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

}

func deleteTodo(w http.ResponseWriter, r *http.Request, user db.User) {
	vars := mux.Vars(r)
	todoId := vars["id"]
	log.Info("delete todo id: ", todoId)
	err := db.DeleteUserTodoById(user.Id, todoId)
	simpleResponse(w,err, "删除成功")
}

func createTodoDone(w http.ResponseWriter, r *http.Request, user db.User) {
	r.ParseForm()
	did_at_string := r.Form.Get("did_at")
	vars := mux.Vars(r)
	todoId := vars["todo_id"]
	log.Info("createTodoDone todo_id:", todoId)
	var todo db.Todo
	if bson.IsObjectIdHex(todoId){
		todo = db.FindTodo(bson.M{"_id": bson.ObjectIdHex(todoId), "user_id": user.Id})
		log.Info("find todo ,content=====",todo.Content)
	}else{
		log.Info("createTodoDone 不合格的 todo id ")
		errResponse(w, errors.New("不合格的id"), "todo id 不正确")
	}
	if len(todo.Id.Hex()) < 10  {
		log.Info("createTodoDone todo不存在 ")
		errResponse(w, errors.New("todo不存在"), "添加失败")
	}else{
		var did_at time.Time
		if did_at_string == "" {
			log.Debug("did_at time now ")
			did_at = time.Now()
		}else{
			i, err := strconv.ParseInt(did_at_string, 10, 64)
			did_at = time.Unix(i, 0)
			errResponse(w, err, "时间戳不标准")
		}
		done_id, err := db.CreateDone(todo,did_at)
		content := map[string]interface{}{"id": done_id.Hex()}
		fullResponse(w, "保存成功", err, content)
	}
}