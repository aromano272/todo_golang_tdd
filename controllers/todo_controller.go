package controllers

import (
	"encoding/json"
	"github.com/aromano272/todo_golang_tdd/data"
	"github.com/aromano272/todo_golang_tdd/models"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

type TodoController struct {
	source data.Source
}

func NewTodoController(source data.Source) *TodoController {
	return &TodoController{source}
}

func (tc TodoController) GetTodos(res http.ResponseWriter, req *http.Request) {
	todos, err := tc.source.ReadAll()

	if err != nil {
		ApiError{err.Error()}.Serve(res, http.StatusBadRequest)
		return
	}

	applyDefaults(res)
	json.NewEncoder(res).Encode(todos)
}

func (tc TodoController) GetTodo(res http.ResponseWriter, req *http.Request) {
	id, ok := mux.Vars(req)["id"]

	if !ok {
		ApiError{"id field is required"}.Serve(res, http.StatusBadRequest)
		return
	}

	key := models.NewKey(id)

	model, err := tc.source.Read(key)
	if err != nil {
		ApiError{err.Error()}.Serve(res, http.StatusBadRequest)
		return
	}

	applyDefaults(res)
	json.NewEncoder(res).Encode(model)
}

func (tc TodoController) CreateTodo(res http.ResponseWriter, req *http.Request) {
	todo := &models.Todo{}

	err := json.NewDecoder(req.Body).Decode(todo)
	if err != nil {
		ApiError{"Error reading request"}.Serve(res, http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	newtodo, err := tc.source.Create(todo)

	if err != nil {
		ApiError{err.Error()}.Serve(res, http.StatusBadRequest)
		return
	}

	applyDefaults(res)
	json.NewEncoder(res).Encode(newtodo)
}

func (tc TodoController) UpdateTodo(res http.ResponseWriter, req *http.Request) {
	todo := &models.Todo{}

	if err := json.NewDecoder(req.Body).Decode(todo); err != nil {
		ApiError{"Error reading request"}.Serve(res, http.StatusBadRequest).Log(err)
		return
	}


	if err := tc.source.Update(todo); err != nil {
		ApiError{err.Error()}.Serve(res, http.StatusBadRequest).Log(err)
		return
	}

	res.WriteHeader(http.StatusOK)
}

func (tc TodoController) DeleteTodo(res http.ResponseWriter, req *http.Request) {

}

func applyDefaults(res http.ResponseWriter) {
	res.Header().Set("Content-Type", "application/json")
}
