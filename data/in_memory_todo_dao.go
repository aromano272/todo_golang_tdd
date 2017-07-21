package data

import (
	"errors"
	"github.com/aromano272/todo_golang_tdd/models"
	"gopkg.in/mgo.v2/bson"
	"github.com/aromano272/todo_golang_tdd/apierrors"
	"github.com/aromano272/todo_golang_tdd/utils"
)

type InMemoryTodoDAO struct {
	storage []*models.Todo
}

var inMemoryTodoDAO *InMemoryTodoDAO

func GetInMemoryTodoDAO() *InMemoryTodoDAO {
	if inMemoryTodoDAO == nil {
		inMemoryTodoDAO = &InMemoryTodoDAO{}
	}
	return inMemoryTodoDAO
}

func (dao *InMemoryTodoDAO) Read(key models.Key) (*models.Todo, error) {
	if !bson.IsObjectIdHex(key.GetKey()) {
		return nil, errors.New(apierrors.InvalidId)
	}
	var todo *models.Todo
	for _, val := range dao.storage {
		if val.GetKey() == key.GetKey() {
			todo = val
		}
	}
	if todo != nil {
		return todo, nil
	} else {
		return nil, errors.New(apierrors.IdNotFound)
	}
}

func (dao *InMemoryTodoDAO) ReadAll() ([]*models.Todo, error) {
	return dao.storage, nil
}

func (dao *InMemoryTodoDAO) Create(todo *models.Todo) (*models.Todo, error) {
	id := bson.NewObjectId().Hex()
	todo.SetKey(id)
	dao.storage = append(dao.storage, todo)

	return todo, nil
}

func (dao *InMemoryTodoDAO) Update(key models.Key, todo *models.Todo) error {
	if !bson.IsObjectIdHex(key.GetKey()) {
		return errors.New(apierrors.InvalidId)
	}
	var td *models.Todo
	for _, val := range dao.storage {
		if val.GetKey() == key.GetKey() {
			td = val
		}
	}
	if td != nil {
		*td = *todo
	} else {
		return errors.New(apierrors.IdNotFound)
	}

	return nil
}

func (dao *InMemoryTodoDAO) Delete(key models.Key) error {
	if !bson.IsObjectIdHex(key.GetKey()) {
		return errors.New(apierrors.InvalidId)
	}
	found := false
	for pos, val := range dao.storage {
		if val.GetKey() == key.GetKey() {
			dao.storage = utils.RemoveTodoFromSlice(dao.storage, pos)
			found = true
		}
	}
	if !found {
		return errors.New(apierrors.IdNotFound)
	}

	return nil
}

func (dao *InMemoryTodoDAO) Clear() {
	dao.storage = nil
}

func (dao *InMemoryTodoDAO) Add(todos ...*models.Todo) []*models.Todo {
	for _, val := range todos {
		val.SetKey(dao.NewKey())
		dao.storage = append(dao.storage, val)
	}

	return dao.storage
}

func (dao *InMemoryTodoDAO) NewKey() string {
	return bson.NewObjectId().Hex()
}
