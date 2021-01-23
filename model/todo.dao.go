package model

var (
	todoCollection = map[int64]*Todo{
		001: {Title: "My First Todo task", Desc: "Finish todo app", isComplete: false},
		002: {Title: "My Second Todo task", Desc: "Finish todo app", isComplete: false},
	}
)

func GetTodo(todoId) (Todo, error) {

}
