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
		CONSTRAINT products_id_pk PRIMARY KEY (id) 
	)`
	psqlCreateProduct = `INSERT INTO users(name, age) VALUES($1, $2) RETURNING id`
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
	stmt, err := p.db.Prepare(psqlCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		person.Name,
		person.Age,
	).Scan(&person.ID)

	if err != nil {
		return err
	}

	fmt.Printf("se creo el producto correctamente")

	return nil
}


