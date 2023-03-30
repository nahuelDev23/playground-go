package services

import (
	"github.com/nahueldev23/playground/repository"
)

type login struct {
	storage repository.Login
}

func NewLogin(storage repository.Login) login {
	return login{storage}
}
