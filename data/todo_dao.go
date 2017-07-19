package data

import (
	"errors"
	"github.com/aromano272/todo_golang_tdd/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TodoDAO struct {
	session *mgo.Session
}

func NewTodoDAO(session *mgo.Session) *TodoDAO {
	return &TodoDAO{session}
}

func (dao *TodoDAO) Read(key models.Key) (models.Model, error) {
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

func (dao *TodoDAO) ReadAll() ([]models.Model, error) {
	var todos []models.Todo

	if err := coll(dao).Find(nil).All(&todos); err != nil {
		errors.New("TODO: to be implemented") // TODO: implement
	}

	mdls := make([]models.Model, len(todos))
	for i := range todos {
		mdls[i] = &todos[i]
	}

	return mdls, nil
}

func (dao *TodoDAO) Create(model models.Model) (models.Model, error) {
	todo := model.(*models.Todo)

	if todo.Title == "" {
		return nil, errors.New("The title field is mandatory")
	}

	todo.Id = bson.NewObjectId()

	if err := coll(dao).Insert(todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (dao *TodoDAO) Update(model models.Model) error {
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"n": 1}},
		ReturnNew: false,
	}

	_, err := coll(dao).Find(bson.M{"_id": model.GetKey()}).Apply(change, nil)

	return err
}

func (dao *TodoDAO) Delete(model models.Model) error {
	return coll(dao).Remove(bson.M{"_id": model.GetKey()})
}

func coll(dao *TodoDAO) *mgo.Collection {
	return dao.session.DB("cooldb").C("todos")
}
