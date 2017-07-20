package models


// Model is the generic interface to represent data that's stored in db.DB implementations
type Model interface {
	Key
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

func NewKey(key string) Id {
	return Id{key}
}
