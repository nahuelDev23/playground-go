package repository

import (
	"database/sql"
	"fmt"

	"github.com/nahueldev23/playground/models"
)

var (
	psqlMigrateUser = `CREATE TABLE IF NOT EXISTS users(
		id SERIAL NOT NULL,
		name VARCHAR(25) NOT NULL,
    age VARCHAR(2) NOT NULL,
    password VARCHAR(100) NOT NULL,
    details VARCHAR(100),
		CONSTRAINT products_id_pk PRIMARY KEY (id) 
	)`
	psqlCreatePerson  = `INSERT INTO users(name, age,password,details) VALUES($1, $2, $3,$4) RETURNING id`
	psqlGetAllPersons = `SELECT id, name,age FROM users`
)

//messageType implements interface Storage
type psqlProduct struct {
	db *sql.DB
}

func NewPsqlPerson(db *sql.DB) *psqlProduct {
	return &psqlProduct{db}
}

// Migrate implement the interface product.Storage
func (p *psqlProduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateUser)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("migración de producto ejecutada correctamente")
	return nil
}

func (p *psqlProduct) Create(person *models.Person) error {
	stmt, err := p.db.Prepare(psqlCreatePerson)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		person.Name,
		person.Age,
		person.Password,
		person.Details,
	).Scan(&person.ID)

	if err != nil {
		return err
	}

	fmt.Printf("se creo el producto correctamente")

	return nil
}

//return slice of person
func (p *psqlProduct) GetAll() (models.Persons, error) {
	stmt, err := p.db.Prepare(psqlGetAllPersons)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	//make a empty slice of persons
	modelPersons := make(models.Persons, 0)

	for rows.Next() {
		person, err := scanRowPerson(rows)
		if err != nil {
			return nil, err
		}
		modelPersons = append(modelPersons, person)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return modelPersons, nil

}

type scanner interface {
	Scan(dest ...interface{}) error
}

func scanRowPerson(s scanner) (*models.Person, error) {
	mp := &models.Person{}
	err := s.Scan(
		&mp.ID,
		&mp.Name,
		&mp.Age,
		&mp.Password,
		&mp.Details,
	)
	if err != nil {
		return &models.Person{}, err
	}

	return mp, nil
}
