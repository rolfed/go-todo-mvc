package controller

import (
	"fmt"
	"net/http"
)

func GetTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET TODO")
}
