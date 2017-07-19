package data

import "github.com/aromano272/todo_golang_tdd/models"

type Source interface {

	Read(models.Key) (models.Model, error)

	ReadAll() ([]models.Model, error)

	Create(models.Model) (models.Model, error)

	Update(models.Model) error

	Delete(models.Model) error

}
