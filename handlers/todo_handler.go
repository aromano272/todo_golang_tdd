package handlers

import (
	"encoding/json"
	"github.com/aromano272/todo_golang_tdd/models"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"github.com/aromano272/todo_golang_tdd/controllers"
)

type TodoHandler struct {
	controller controllers.TodoController
}

func NewTodoHandler(controller controllers.TodoController) *TodoHandler {
	return &TodoHandler{controller}
}

func (handler TodoHandler) ReadAllTodos(res http.ResponseWriter, req *http.Request) {
	var request models.ReadAllTodosRequest
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		NewApiError("Error reading request", http.StatusBadRequest).Serve(res)
		fmt.Println(err)
		return
	}

	todos, apierr := handler.controller.ReadAllTodos(request)

	if apierr != nil {
		apierr.Serve(res)
		return
	}

	json.NewEncoder(res).Encode(todos)
}

func (handler TodoHandler) ReadTodo(res http.ResponseWriter, req *http.Request) {
	var request models.ReadTodoRequest
	// FIXME: Get request might not have a body and this below might not work
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		NewApiError("Error reading request", http.StatusBadRequest).Serve(res)
		fmt.Println(err)
		return
	}
	id, ok := mux.Vars(req)["id"]

	todo, apierr := handler.controller.ReadTodo(request)
	if apierr != nil {
		apierr.Serve(res)
		return
	}

	json.NewEncoder(res).Encode(todo)
}

func (handler TodoHandler) CreateTodo(res http.ResponseWriter, req *http.Request) {
	var request models.CreateTodoRequest

	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		NewApiError("Error reading request", http.StatusBadRequest).Serve(res)
		fmt.Println(err)
		return
	}

	newtodo, apierr := handler.controller.CreateTodo(request)

	if apierr != nil {
		apierr.Serve(res)
		return
	}

	json.NewEncoder(res).Encode(newtodo)
}

func (handler TodoHandler) UpdateTodo(res http.ResponseWriter, req *http.Request) {
	var request models.UpdateTodoRequest

	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		NewApiError("Error reading request", http.StatusBadRequest).ServeAndLog(res, err)
		return
	}


	if apierr := handler.controller.UpdateTodo(request); apierr != nil {
		apierr.Serve(res)
		return
	}

	res.WriteHeader(http.StatusOK)
}

func (handler TodoHandler) DeleteTodo(res http.ResponseWriter, req *http.Request) {
	var request models.DeleteTodoRequest

	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		NewApiError("Error reading request", http.StatusBadRequest).ServeAndLog(res, err)
		return
	}

	if apierr := handler.controller.DeleteTodo(request); apierr != nil {
		apierr.Serve(res)
		return
	}

	res.WriteHeader(http.StatusOK)
}
