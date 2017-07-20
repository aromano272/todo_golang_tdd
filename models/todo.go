package models

type Todo struct {
	Id    string        `json:"id" bson:"_id"`
	Title string        `json:"title" bson:"title"`
	Desc  string        `json:"desc" bson:"desc"`
}

func (todo *Todo) SetKey(key string) {
	todo.Id = key
}

func (todo *Todo) GetKey() string {
	return todo.Id
}

type ReadTodoRequest struct {
	Id string `json:"id"`
}

type ReadAllTodosRequest struct {
}

type CreateTodoRequest struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type UpdateTodoRequest struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type DeleteTodoRequest struct {
	Id string `json:"id"`
}
