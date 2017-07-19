package models

import (
	"encoding"
)

// Model is the generic interface to represent data that's stored in db.DB implementations
type Model interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler

	Key
	// Set replaces the contents of the model with the given model
	Set(Model) error
}

type Key interface {
	GetKey() string
}

type Id struct {
	Key string
}

func (id Id) GetKey() string {
	return id.Key
}
