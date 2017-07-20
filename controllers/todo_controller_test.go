package controllers

import (
	"testing"
	"github.com/aromano272/todo_golang_tdd/data"
)

var tc = NewTodoController(data.GetInMemoryTodoDAO())

func TestTodoController(t *testing.T) {

	tc.CreateTodo()

}
