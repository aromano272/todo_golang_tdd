package data

import (
	"errors"
	"github.com/aromano272/todo_golang_tdd/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/aromano272/todo_golang_tdd/apierrors"
	"fmt"
)

type TodoSource interface {
	Read(models.Key) (*models.Todo, error)

	ReadAll() ([]*models.Todo, error)

	Create(*models.Todo) (*models.Todo, error)

	Update(models.Key, *models.Todo) error

	Delete(models.Key) error
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
		return nil, errors.New(apierrors.InvalidId)
	}

	fmt.Println(id)

	oid := bson.ObjectIdHex(id)

	todo := &models.Todo{}

	if err := dao.coll().FindId(oid).One(todo); err != nil {
		return nil, errors.New(apierrors.IdNotFound)
	}

	return todo, nil
}

func (dao *TodoDAO) ReadAll() ([]*models.Todo, error) {
	var todos []*models.Todo

	if err := dao.coll().Find(nil).All(&todos); err != nil {
		errors.New("TODO: to be implemented") // TODO: implement
	}

	return todos, nil
}

func (dao *TodoDAO) Create(todo *models.Todo) (*models.Todo, error) {
	if todo.Title == "" {
		return nil, errors.New(apierrors.TodoTitleFieldMissing)
	}

	todo.Id = bson.NewObjectId()

	if err := dao.coll().Insert(todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (dao *TodoDAO) Update(key models.Key, todo *models.Todo) error {
	if todo.Title == "" {
		return errors.New(apierrors.TodoTitleFieldMissing)
	}

	if !bson.IsObjectIdHex(key.GetKey()) {
		return errors.New(apierrors.InvalidId)
	}

	oid := bson.ObjectIdHex(key.GetKey())

	todo.SetKey(oid.Hex())

	change := mgo.Change{
		Update:    bson.M{"$set": todo},
		ReturnNew: false,
	}

	_, err := dao.coll().Find(bson.M{"_id": oid}).Apply(change, nil)

	return err
}

func (dao *TodoDAO) Delete(key models.Key) error {
	id := key.GetKey()

	if !bson.IsObjectIdHex(id) {
		return errors.New(apierrors.InvalidId)
	}

	oid := bson.ObjectIdHex(id)

	return dao.coll().Remove(bson.M{"_id": oid})
}

func (dao *TodoDAO) coll() *mgo.Collection {
	return dao.session.DB("cooldb").C("todos")
}
