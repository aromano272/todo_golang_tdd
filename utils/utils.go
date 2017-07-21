package utils

import (
	"github.com/aromano272/todo_golang_tdd/models"
	"testing"
)

func AreTodosEqual(a, b []*models.Todo, t *testing.T) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if *a[i] != *b[i] {
			return false
		}
	}

	return true
}

func RemoveTodoFromSlice(slice []*models.Todo, index int) []*models.Todo {
	return append(slice[:index], slice[index+1:]...)
}
