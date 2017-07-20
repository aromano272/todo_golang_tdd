package data

import (
	"errors"
	"github.com/aromano272/todo_golang_tdd/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TodoSource interface {

	Read(models.Key) (*models.Todo, error)

	ReadAll() ([]*models.Todo, error)

	Create(*models.Todo) (*models.Todo, error)

	Update(*models.Todo) error

	Delete(*models.Todo) error

}

type TodoDAO struct {
	session *mgo.Session
}

func NewTodoDAO(session *mgo.Session) *TodoDAO {
	return &TodoDAO{session}
}

func (dao *TodoDAO) Read(key models.Key) (*models.Todo, error) {
	id := key.GetKey()

	if !bson.IsObjectIdHex(id) {
		return nil, errors.New("id field is invalid")
	}

	oid := bson.ObjectIdHex(id)

	todo := &models.Todo{}

	if err := coll(dao).FindId(oid).One(todo); err != nil {
		return nil, errors.New("todo with this id was not found")
	}

	return todo, nil
}

func (dao *TodoDAO) ReadAll() ([]*models.Todo, error) {
	var todos []*models.Todo

	if err := coll(dao).Find(nil).All(&todos); err != nil {
		errors.New("TODO: to be implemented") // TODO: implement
	}

	return todos, nil
}

func (dao *TodoDAO) Create(todo *models.Todo) (*models.Todo, error) {

	if todo.Title == "" {
		return nil, errors.New("The title field is mandatory")
	}

	todo.Id = bson.NewObjectId()

	if err := coll(dao).Insert(todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (dao *TodoDAO) Update(todo *models.Todo) error {

	if todo.Title == "" {
		return errors.New("The title field is mandatory")
	}

	if !bson.IsObjectIdHex(todo.GetKey()) {
		return errors.New("id field is invalid")
	}

	change := mgo.Change{
		Update:    bson.M{"$set": todo},
		ReturnNew: false,
	}

	oid := bson.ObjectIdHex(todo.GetKey())

	_, err := coll(dao).Find(bson.M{"_id": oid}).Apply(change, nil)

	return err
}

func (dao *TodoDAO) Delete(todo *models.Todo) error {
	id := todo.GetKey()

	if !bson.IsObjectIdHex(id) {
		return errors.New("id field is invalid")
	}

	oid := bson.ObjectIdHex(id)

	return coll(dao).Remove(bson.M{"_id": oid})
}

func coll(dao *TodoDAO) *mgo.Collection {
	return dao.session.DB("cooldb").C("todos")
}
