package app

import (
	"log"
	"net/http"

	"../controller"
)

func StartApp() {

	http.HandleFunc("/todo", controller.GetTodo)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
