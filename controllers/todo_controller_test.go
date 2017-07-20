package controllers

import (
	"testing"
	"github.com/aromano272/todo_golang_tdd/data"
	"github.com/aromano272/todo_golang_tdd/models"
)

var tc = NewTodoController(data.GetInMemoryTodoDAO())

func TestTodoController(t *testing.T) {

	tc.CreateTodo(models.CreateTodoRequest{Title: "Title", Desc: "Desc"})


}
