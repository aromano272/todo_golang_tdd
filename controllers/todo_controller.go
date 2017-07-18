package controllers

import (
	"gopkg.in/mgo.v2"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
	"github.com/aromano272/todo_golang_tdd/models"
	"log"
	"encoding/json"
)

type TodoController struct {
	session *mgo.Session
}

func NewTodoController(session *mgo.Session) *TodoController {
	return &TodoController{session}
}

func (tc TodoController) GetTodos(res http.ResponseWriter, req *http.Request) {

}

func (tc TodoController) GetTodo(res http.ResponseWriter, req *http.Request) {
	id, ok := mux.Vars(req)["id"]

	if !ok {
		ApiError{"id field is required"}.Serve(res, http.StatusBadRequest)
		return
	}

	if !bson.IsObjectIdHex(id) {
		ApiError{"id field is invalid"}.Serve(res, http.StatusBadRequest)
		return
	}

	oid := bson.ObjectIdHex(id)

	todo := models.Todo{}

	if err := Coll(tc).FindId(oid).One(&todo); err != nil {
		ApiError{"todo with this id was not found"}.Serve(res, http.StatusNotFound)
		return
	}

	jsontodo, err := todo.MarshalBinary()
	if err != nil {
		log.Fatalln(err)
	}

	json.NewEncoder(res).Encode(jsontodo)
}

func (tc TodoController) CreateTodo(res http.ResponseWriter, req *http.Request) {

}

func (tc TodoController) UpdateTodo(res http.ResponseWriter, req *http.Request) {

}

func (tc TodoController) DeleteTodo(res http.ResponseWriter, req *http.Request) {

}

func Coll(tc TodoController) *mgo.Collection {
	return tc.session.DB("cooldb").C("todos")
}

