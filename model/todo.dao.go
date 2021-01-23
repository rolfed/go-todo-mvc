package model

import (
	"fmt"
	"net/http"

	"../util"
)

// Mock database
var (
	todoCollection = map[int64]*Todo{
		001: &Todo{Title: "My first todo task", Desc: "Finish todo app", isComplete: false},
		002: &Todo{Title: "My second todo task", Desc: "Finish todo app", isComplete: false},
	}
)

// GetTodo querys database todo column by id
func GetTodo(todoId int64) (*Todo, *util.AppError) {
	if todo := todoCollection[todoId]; todo != nil {
		return todo, nil
	}

	return nil, &util.AppError{
		Message:    fmt.Sprintf("todo %v was not found", todoId),
		StatusCode: http.StatusNotFound,
		Code:       "not found",
	}

}
