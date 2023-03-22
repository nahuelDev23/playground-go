package services

import (
	"github.com/nahueldev23/playground/repository"
)

type person struct {
	storage repository.Person
}

func NewPerson(storage repository.Person) person {
	return person{storage}
}
