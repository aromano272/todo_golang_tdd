package data

import (
	"errors"
	"github.com/aromano272/todo_golang_tdd/models"
	"gopkg.in/mgo.v2/bson"
	"github.com/aromano272/todo_golang_tdd/apierrors"
)

type InMemoryTodoDAO struct {
	storage map[string]*models.Todo
}

var inMemoryTodoDAO *InMemoryTodoDAO

func GetInMemoryTodoDAO() *InMemoryTodoDAO {
	if inMemoryTodoDAO == nil {
		inMemoryTodoDAO = &InMemoryTodoDAO{make(map[string]*models.Todo)}
	}
	return inMemoryTodoDAO
}

func (dao *InMemoryTodoDAO) Read(key models.Key) (*models.Todo, error) {
	if todo, ok := dao.storage[key.GetKey()]; ok {
		return todo, nil
	} else {
		return nil, errors.New(apierrors.IdNotFound)
	}
}

func (dao *InMemoryTodoDAO) ReadAll() ([]*models.Todo, error) {
	var todos []*models.Todo
	for _, val := range dao.storage {
		todos = append(todos, val)
	}
	return todos, nil
}

func (dao *InMemoryTodoDAO) Create(todo *models.Todo) (*models.Todo, error) {
	id := bson.NewObjectId().Hex()
	todo.SetKey(id)
	dao.storage[id] = todo

	return todo, nil
}

func (dao *InMemoryTodoDAO) Update(todo *models.Todo) error {
	if _, ok := dao.storage[todo.GetKey()]; ok {
		dao.storage[todo.GetKey()] = todo
	} else {
		return errors.New(apierrors.IdNotFound)
	}

	return nil
}

func (dao *InMemoryTodoDAO) Delete(key models.Key) error {
	if _, ok := dao.storage[key.GetKey()]; ok {
		delete(dao.storage, key.GetKey())
	} else {
		return errors.New(apierrors.IdNotFound)
	}

	return nil
}
