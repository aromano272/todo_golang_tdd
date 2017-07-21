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
	if req.Id == "" {
		return nil, apierrors.NewApiError(apierrors.IdFieldMissing, http.StatusBadRequest)
	}

	key := models.NewKey(req.Id)

	todo, err := tc.source.Read(key)
	if err != nil {
		return nil, apierrors.NewApiError(err.Error(), http.StatusBadRequest)
	}

	return todo, nil
}

func (tc TodoController) CreateTodo(req models.CreateTodoRequest) (*models.Todo, apierrors.ApiError) {
	if req.Title == "" {
		return nil, apierrors.NewApiError(apierrors.TodoTitleFieldMissing, http.StatusBadRequest)
	}

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
	if req.Id == "" {
		return apierrors.NewApiError(apierrors.IdFieldMissing, http.StatusBadRequest)
	}

	if req.Title == "" {
		return apierrors.NewApiError(apierrors.TodoTitleFieldMissing, http.StatusBadRequest)
	}

	todo := &models.Todo{
		Title: req.Title,
		Desc:  req.Desc,
	}

	if err := tc.source.Update(models.NewKey(req.Id), todo); err != nil {
		return apierrors.NewApiError(err.Error(), http.StatusBadRequest)
	}

	return nil
}

func (tc TodoController) DeleteTodo(req models.DeleteTodoRequest) apierrors.ApiError {
	if req.Id == "" {
		return apierrors.NewApiError(apierrors.IdFieldMissing, http.StatusBadRequest)
	}

	if err := tc.source.Delete(models.NewKey(req.Id)); err != nil {
		return apierrors.NewApiError(err.Error(), http.StatusBadRequest)
	}

	return nil
}
