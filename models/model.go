package models


// Model is the generic interface to represent data that's stored in db.DB implementations
type Model interface {
	Key
}

type Key interface {
	SetKey(key string)
	GetKey() string
}

type Id struct {
	Key string
}

func (id Id) GetKey() string {
	return id.Key
}

func (id Id) SetKey(key string) {
	id.Key = key
}

func NewKey(key string) Id {
	return Id{key}
}
