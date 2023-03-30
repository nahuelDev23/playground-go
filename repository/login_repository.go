package repository

import "github.com/nahueldev23/playground/models"

type Login interface {
	SignIn(login *models.Login) (bool, error)
}
