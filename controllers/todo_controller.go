package controllers

import (
	"github.com/aromano272/todo_golang_tdd/data"
	"github.com/aromano272/todo_golang_tdd/models"
	"net/http"
	"github.com/aromano272/todo_golang_tdd/handlers"
)

type TodoController struct {
	source data.TodoSource
}

func NewTodoController(source data.TodoSource) *TodoController {
	return &TodoController{source}
}

func (tc TodoController) ReadAllTodos(req models.ReadAllTodosRequest) ([]*models.Todo, handlers.ApiError) {
	todos, err := tc.source.ReadAll()

	if err != nil {
		return nil, handlers.NewApiError(err.Error(), http.StatusBadRequest)
	}

	return todos, nil
}

func (tc TodoController) ReadTodo(req models.ReadTodoRequest) (*models.Todo, handlers.ApiError) {
	id := req.Id

	if id == "" {
		return nil, handlers.NewApiError("id field is required", http.StatusBadRequest)
	}

	key := models.NewKey(id)

	todo, err := tc.source.Read(key)
	if err != nil {
		return nil, handlers.NewApiError(err.Error(), http.StatusBadRequest)
	}

	return todo, nil
}

func (tc TodoController) CreateTodo(req models.CreateTodoRequest) (*models.Todo, handlers.ApiError) {
	todo := &models.Todo{
		Title: req.Title,
		Desc:  req.Desc,
	}

	newtodo, err := tc.source.Create(todo)
	if err != nil {
		return nil, handlers.NewApiError(err.Error(), http.StatusBadRequest)
	}

	return newtodo, nil
}

func (tc TodoController) UpdateTodo(req models.UpdateTodoRequest) handlers.ApiError {
	todo := &models.Todo{
		Id:    req.Id,
		Title: req.Title,
		Desc:  req.Desc,
	}

	if err := tc.source.Update(todo); err != nil {
		return handlers.NewApiError(err.Error(), http.StatusBadRequest)
	}

	return nil
}

func (tc TodoController) DeleteTodo(req models.DeleteTodoRequest) handlers.ApiError {
	todo := &models.Todo{
		Id: req.Id,
	}

	if err := tc.source.Delete(todo); err != nil {
		return handlers.NewApiError(err.Error(), http.StatusBadRequest)
	}

	return nil
}
