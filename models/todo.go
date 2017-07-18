package models

import (
	"encoding/json"
	"errors"
	"gopkg.in/mgo.v2/bson"
)

type Todo struct {

	Id bson.ObjectId `json:"id" bson:"_id"`
	Title string `json:"title" bson:"title"`
	Desc  string `json:"desc" bson:"desc"`

}

func (todo *Todo) MarshalBinary() (data []byte, err error) {
	return json.Marshal(todo)
}

func (todo *Todo) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, todo)
}

func (todo *Todo) Set(model Model) error {
	t, ok := model.(*Todo)
	if !ok {
		return errors.New("Incompatible type")
	}
	*todo = *t
	return nil
}