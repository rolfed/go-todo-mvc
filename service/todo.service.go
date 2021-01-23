package service

import (
	"../model"
	"../util"
)

// Service is responsible for constructing multiple models to
// accomplish task

// GetTodo  
func GetTodo(todoId int64) (*model.Todo, *util.AppError) {
	return model.GetTodo(todoId)
}
