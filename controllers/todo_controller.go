package controllers

import (
	"github.com/aromano272/todo_golang_tdd/data"
	"github.com/aromano272/todo_golang_tdd/models"
	"net/http"
	"github.com/aromano272/todo_golang_tdd/apierrors"
)

type TodoController struct {
	source data.TodoSource
}

func NewTodoController(source data.TodoSource) *TodoController {
	return &TodoController{source}
}

func (tc TodoController) ReadAllTodos(req models.ReadAllTodosRequest) ([]*models.Todo, apierrors.ApiError) {
	todos, err := tc.source.ReadAll()

	if err != nil {
		return nil, apierrors.NewApiError(err.Error(), http.StatusBadRequest)
	}

	return todos, nil
}

func (tc TodoController) ReadTodo(req models.ReadTodoRequest) (*models.Todo, apierrors.ApiError) {
	id := req.Id

	if id == "" {
		return nil, apierrors.NewApiError("id field is required", http.StatusBadRequest)
	}

	key := models.NewKey(id)

	todo, err := tc.source.Read(key)
	if err != nil {
		return nil, apierrors.NewApiError(err.Error(), http.StatusBadRequest)
	}

	return todo, nil
}

func (tc TodoController) CreateTodo(req models.CreateTodoRequest) (*models.Todo, apierrors.ApiError) {
	todo := &models.Todo{
		Title: req.Title,
		Desc:  req.Desc,
	}

	newtodo, err := tc.source.Create(todo)
	if err != nil {
		return nil, apierrors.NewApiError(err.Error(), http.StatusBadRequest)
	}

	return newtodo, nil
}

func (tc TodoController) UpdateTodo(req models.UpdateTodoRequest) apierrors.ApiError {
	todo := &models.Todo{
		Title: req.Title,
		Desc:  req.Desc,
	}

	todo.SetKey(req.Id)

	if err := tc.source.Update(todo); err != nil {
		return apierrors.NewApiError(err.Error(), http.StatusBadRequest)
	}

	return nil
}

func (tc TodoController) DeleteTodo(req models.DeleteTodoRequest) apierrors.ApiError {
	if err := tc.source.Delete(models.NewKey(req.Id)); err != nil {
		return apierrors.NewApiError(err.Error(), http.StatusBadRequest)
	}

	return nil
}
