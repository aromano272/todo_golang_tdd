package models

import "gopkg.in/mgo.v2/bson"

type Todo struct {
	Id    bson.ObjectId `json:"id" bson:"_id"`
	Title string        `json:"title" bson:"title"`
	Desc  string        `json:"desc" bson:"desc"`
}

func (todo *Todo) SetKey(key string) {
	todo.Id = bson.ObjectIdHex(key)
}

func (todo *Todo) GetKey() string {
	return todo.Id.Hex()
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
