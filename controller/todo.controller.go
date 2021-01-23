package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../service"
	"../util"
)

func GetTodo(w http.ResponseWriter, r *http.Request) {
	todoID, err := strconv.ParseInt(r.URL.Query().Get("todo_id"), 10, 64)
	if err != nil {

		getTodoError := &util.AppError{
			Message:    "todo_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad request",
		}

		getTodoErrorJSON, _ := json.Marshal(getTodoError)

		w.WriteHeader(getTodoError.StatusCode)
		w.Write([]byte(getTodoErrorJSON))

		return
	}

	todo, getTodoServiceError := service.GetTodo(todoID)
	if getTodoServiceError != nil {
		// Handle the error and sent to client
		getTodoServiceErrorJSON, _ := json.Marshal(getTodoServiceError)
		w.WriteHeader(getTodoServiceError.StatusCode)
		w.Write([]byte(getTodoServiceErrorJSON))
	}

	// return todo to client
	todoJSON, _ := json.Marshal(todo)
	w.Write(todoJSON)

}
