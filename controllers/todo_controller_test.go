package controllers_test

import (
	"testing"
	"github.com/aromano272/todo_golang_tdd/data"
	"github.com/aromano272/todo_golang_tdd/models"
	"github.com/aromano272/todo_golang_tdd/controllers"
	"github.com/aromano272/todo_golang_tdd/apierrors"
	"net/http"
	"github.com/aromano272/todo_golang_tdd/utils"
)

var dao = data.GetInMemoryTodoDAO()
var tc = controllers.NewTodoController(dao)

func TestTodoController_CreateTodo(t *testing.T) {

	tests := []struct {
		input          models.CreateTodoRequest
		expectedTodo   *models.Todo
		expectedApierr apierrors.ApiError
	}{
		{
			models.CreateTodoRequest{Title: "Title", Desc: "Desc"},
			&models.Todo{Title: "Title", Desc: "Desc"},
			nil,
		},
		{
			models.CreateTodoRequest{Desc: "Desc"},
			nil,
			apierrors.NewApiError(apierrors.TodoTitleFieldMissing, http.StatusBadRequest),
		},
		{
			models.CreateTodoRequest{Title: "Title"},
			&models.Todo{Title: "Title"},
			nil,
		},
		{
			models.CreateTodoRequest{},
			nil,
			apierrors.NewApiError(apierrors.TodoTitleFieldMissing, http.StatusBadRequest),
		},
	}

	for _, test := range tests {
		actualTodo, actualApierr := tc.CreateTodo(test.input)
		// we have to set the Key of the expected equal to the actual, because in this request we dont
		// care about the key, and the comparison will fail if the expected key is nil and the actual is not
		if test.expectedTodo != nil && actualTodo != nil {
			test.expectedTodo.SetKey(actualTodo.GetKey())
		}

		if actualTodo == nil && test.expectedTodo != nil ||
			actualTodo != nil && test.expectedTodo == nil ||
			(actualTodo != nil && test.expectedTodo != nil && (*actualTodo != *test.expectedTodo)) {

			t.Errorf("CreateTodo(%v): expected %v, actual %v", test.input, test.expectedTodo, actualTodo)
		}

		if actualApierr == nil && test.expectedApierr != nil ||
			actualApierr != nil && test.expectedApierr == nil ||
			(actualApierr != nil && test.expectedApierr != nil &&
				(actualApierr.GetError() != test.expectedApierr.GetError() ||
					actualApierr.GetStatusCode() != test.expectedApierr.GetStatusCode())) {
			t.Errorf("CreateTodo(%v): expected %v, actual %v", test.input, test.expectedApierr, actualApierr)
		}
	}

}

func TestTodoController_ReadTodo(t *testing.T) {

	dao.Clear()
	var todos = dao.Add(
		&models.Todo{Title: "Title0", Desc: "Desc0" },
		&models.Todo{Title: "Title1", Desc: "Desc1" },
		&models.Todo{Title: "Title2", Desc: "Desc2" },
		&models.Todo{Title: "Title3", Desc: "Desc3" },
	)

	tests := []struct {
		input          models.ReadTodoRequest
		expectedTodo   *models.Todo
		expectedApierr apierrors.ApiError
	}{
		{
			models.ReadTodoRequest{},
			nil,
			apierrors.NewApiError(apierrors.IdFieldMissing, http.StatusBadRequest),
		},
		{
			models.ReadTodoRequest{Id: dao.NewKey()},
			nil,
			apierrors.NewApiError(apierrors.IdNotFound, http.StatusBadRequest),
		},
		{
			models.ReadTodoRequest{Id: todos[0].GetKey()},
			todos[0],
			nil,
		},
	}

	for _, test := range tests {
		actualTodo, actualApierr := tc.ReadTodo(test.input)

		if actualTodo == nil && test.expectedTodo != nil ||
			actualTodo != nil && test.expectedTodo == nil ||
			(actualTodo != nil && test.expectedTodo != nil && (*actualTodo != *test.expectedTodo)) {

			t.Errorf("ReadTodo(%v): expected %v, actual %v", test.input, test.expectedTodo, actualTodo)
		}

		if actualApierr == nil && test.expectedApierr != nil ||
			actualApierr != nil && test.expectedApierr == nil ||
			(actualApierr != nil && test.expectedApierr != nil &&
				(actualApierr.GetError() != test.expectedApierr.GetError() ||
					actualApierr.GetStatusCode() != test.expectedApierr.GetStatusCode())) {
			t.Errorf("ReadTodo(%v): expected %v, actual %v", test.input, test.expectedApierr, actualApierr)
		}

	}

}

func TestTodoController_ReadAllTodos(t *testing.T) {

	dao.Clear()
	var todos = dao.Add(
		&models.Todo{Title: "Title0", Desc: "Desc0" },
		&models.Todo{Title: "Title1", Desc: "Desc1" },
		&models.Todo{Title: "Title2", Desc: "Desc2" },
		&models.Todo{Title: "Title3", Desc: "Desc3" },
	)

	tests := []struct {
		input          models.ReadAllTodosRequest
		expectedTodos  []*models.Todo
		expectedApierr apierrors.ApiError
	}{
		{
			models.ReadAllTodosRequest{},
			todos,
			nil,
		},
	}

	for _, test := range tests {
		actualTodos, actualApierr := tc.ReadAllTodos(test.input)

		if !utils.AreTodosEqual(actualTodos, test.expectedTodos, t) {
			t.Errorf("ReadAllTodos(%v): expected %v, actual %v", test.input, test.expectedTodos, actualTodos)
		}

		if actualApierr == nil && test.expectedApierr != nil ||
			actualApierr != nil && test.expectedApierr == nil ||
			(actualApierr != nil && test.expectedApierr != nil &&
				(actualApierr.GetError() != test.expectedApierr.GetError() ||
					actualApierr.GetStatusCode() != test.expectedApierr.GetStatusCode())) {
			t.Errorf("ReadAllTodos(%v): expected %v, actual %v", test.input, test.expectedApierr, actualApierr)
		}

	}

}

func TestTodoController_UpdateTodo(t *testing.T) {

	dao.Clear()
	var todos = dao.Add(
		&models.Todo{Title: "Title0", Desc: "Desc0" },
		&models.Todo{Title: "Title1", Desc: "Desc1" },
		&models.Todo{Title: "Title2", Desc: "Desc2" },
		&models.Todo{Title: "Title3", Desc: "Desc3" },
	)

	tests := []struct {
		input          models.UpdateTodoRequest
		expectedApierr apierrors.ApiError
	}{
		{
			models.UpdateTodoRequest{Title: "updated title", Desc: "updated desc"},
			apierrors.NewApiError(apierrors.IdFieldMissing, http.StatusBadRequest),
		},
		{
			models.UpdateTodoRequest{Id: "daslkdasd", Title: "updated title", Desc: "updated desc"},
			apierrors.NewApiError(apierrors.InvalidId, http.StatusBadRequest),
		},
		{
			models.UpdateTodoRequest{Id: dao.NewKey(), Title: "updated title", Desc: "updated desc"},
			apierrors.NewApiError(apierrors.IdNotFound, http.StatusBadRequest),
		},
		{
			models.UpdateTodoRequest{Id: todos[0].GetKey(), Desc: "updated desc"},
			apierrors.NewApiError(apierrors.TodoTitleFieldMissing, http.StatusBadRequest),
		},
		{
			models.UpdateTodoRequest{Id: todos[0].GetKey(), Title: "updated title", Desc: "updated desc"},
			nil,
		},
	}

	for _, test := range tests {
		actualApierr := tc.UpdateTodo(test.input)

		if actualApierr == nil && test.expectedApierr != nil ||
			actualApierr != nil && test.expectedApierr == nil ||
			(actualApierr != nil && test.expectedApierr != nil &&
				(actualApierr.GetError() != test.expectedApierr.GetError() ||
					actualApierr.GetStatusCode() != test.expectedApierr.GetStatusCode())) {
			t.Errorf("UpdateTodo(%v): expected %v, actual %v", test.input, test.expectedApierr, actualApierr)
		}

	}

}

func TestTodoController_DeleteTodo(t *testing.T) {

	dao.Clear()
	var todos = dao.Add(
		&models.Todo{Title: "Title0", Desc: "Desc0" },
		&models.Todo{Title: "Title1", Desc: "Desc1" },
		&models.Todo{Title: "Title2", Desc: "Desc2" },
		&models.Todo{Title: "Title3", Desc: "Desc3" },
	)

	tests := []struct {
		input          models.DeleteTodoRequest
		expectedApierr apierrors.ApiError
	}{
		{
			models.DeleteTodoRequest{},
			apierrors.NewApiError(apierrors.IdFieldMissing, http.StatusBadRequest),
		},
		{
			models.DeleteTodoRequest{Id: "daslkdasd"},
			apierrors.NewApiError(apierrors.InvalidId, http.StatusBadRequest),
		},
		{
			models.DeleteTodoRequest{Id: dao.NewKey()},
			apierrors.NewApiError(apierrors.IdNotFound, http.StatusBadRequest),
		},
		{
			models.DeleteTodoRequest{Id: todos[0].GetKey()},
			nil,
		},
	}

	for _, test := range tests {
		actualApierr := tc.DeleteTodo(test.input)

		if actualApierr == nil && test.expectedApierr != nil ||
			actualApierr != nil && test.expectedApierr == nil ||
			(actualApierr != nil && test.expectedApierr != nil &&
				(actualApierr.GetError() != test.expectedApierr.GetError() ||
					actualApierr.GetStatusCode() != test.expectedApierr.GetStatusCode())) {
			t.Errorf("DeleteTodo(%v): expected %v, actual %v", test.input, test.expectedApierr, actualApierr)
		}

	}

}


