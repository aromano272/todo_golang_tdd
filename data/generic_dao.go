package data

import (
	"errors"
	"github.com/aromano272/todo_golang_tdd/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type GenericDAO struct {
	coll *mgo.Collection
}

func NewGenericDAO(coll *mgo.Collection) *GenericDAO {
	return &GenericDAO{coll}
}

func (dao *GenericDAO) Read(key models.Key) (models.Model, error) {
	id := key.GetKey()

	if !bson.IsObjectIdHex(id) {
		return nil, errors.New("id field is invalid")
	}

	oid := bson.ObjectIdHex(id)

	todo := &models.Todo{}

	if err := dao.coll.FindId(oid).One(todo); err != nil {
		return nil, errors.New("todo with this id was not found")
	}

	return todo, nil
}

func (dao *GenericDAO) ReadAll() ([]models.Model, error) {
	var todos []models.Model

	if err := dao.coll.Find(nil).All(&todos); err != nil {
		errors.New("TODO: to be implemented") // TODO: implement
	}

	return todos, nil
}

func (dao *GenericDAO) Create(model models.Model) (models.Model, error) {
	id := bson.NewObjectId()
	model.SetKey(id.Hex())

	if err := dao.coll.Insert(model); err != nil {
		return nil, err
	}

	return model, nil
}

func (dao *GenericDAO) Update(model models.Model) error {
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"n": 1}},
		ReturnNew: false,
	}

	_, err := dao.coll.Find(bson.M{"_id": model.GetKey()}).Apply(change, nil)

	return err
}

func (dao *GenericDAO) Delete(model models.Model) error {
	return dao.coll.Remove(bson.M{"_id": model.GetKey()})
}
