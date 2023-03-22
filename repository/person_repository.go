package repository

import "github.com/nahueldev23/playground/models"

type Person interface {
	Migrate() error
	Create(person *models.Person) error
	GetAll() (models.Persons,error)
}

