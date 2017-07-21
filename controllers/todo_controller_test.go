package controllers_test

import (
	"testing"
	"github.com/aromano272/todo_golang_tdd/data"
	"github.com/aromano272/todo_golang_tdd/models"
	"github.com/aromano272/todo_golang_tdd/controllers"
	"github.com/aromano272/todo_golang_tdd/apierrors"
	"net/http"
)

var tc = controllers.NewTodoController(data.GetInMemoryTodoDAO())

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
	}

	for _, test := range tests {
		actualTodo, actualApierr := tc.CreateTodo(test.input)

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

		if actualApierr != test.expectedApierr {

		}
	}

}
