package apis

import (
	"fmt"
	"gocheese/db"
	"net/http"
	"os"
)

func getTodos(w http.ResponseWriter, r *http.Request) {
	db.GetAllTodos()
	fmt.Fprintf(
		w,
		"=>  started  from pid %d.\n",
		os.Getpid(),
	)
}
