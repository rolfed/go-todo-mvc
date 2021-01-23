package model

type Todo struct {
	Title string `json:"title"`
	Desc string `json:"desc"`
	isComplete bool `json:"isComplete`
}

type TodoCollection struct {
	collection []Todo `json:"todoCollection"`
}