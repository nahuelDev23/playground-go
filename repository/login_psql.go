package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/nahueldev23/playground/models"
)

var (
	psqlGetPerson = `SELECT id,name,age FROM users WHERE name = $1  AND password = $2`
)

type psqlLogin struct {
	db *sql.DB
}

func NewPsqlLogin(db *sql.DB) *psqlLogin {
	return &psqlLogin{db}
}

func (l *psqlLogin) SignIn(login *models.Login) (bool, error) {

	stmt, err := l.db.Prepare(psqlGetPerson)
	if err != nil {
		log.Printf("%v", err)
		return false, err
	}
	defer stmt.Close()

	resPerson := models.Person{}

	err = stmt.QueryRow(
		login.Name,
		login.Password,
	).Scan(&resPerson.ID, &resPerson.Name, &resPerson.Age)

	if err != nil {
		log.Printf("%v", err)
		return false, err
	}

	fmt.Print("se logeo correctamente")

	return true, nil
}
