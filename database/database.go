package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/nahueldev23/playground/repository"
)

var (
	once sync.Once
	db   *sql.DB
)

type Driver string

const (
	Postgres Driver = "POSTGRES"
)

func New(driver Driver) {
	switch driver {
	case Postgres:
		newPostgresDB()
	}
}

func newPostgresDB() {
	once.Do(func() {
		var err error
		//este db apunta al db global, linea 39
		db, err = sql.Open("postgres", "postgres://edteam:edteam@localhost:5432/godb?sslmode=disable")
		if err != nil {
			log.Fatalf("no se pudo abrir la db de mysql, error: %v", err)
		}
		if err = db.Ping(); err != nil {
			log.Fatalf("can't do ping: %v", err)
		}

		fmt.Println("conectado a mySQL")
	})
}

// DAOProduct factory of product.Storage
func DAOPerson(driver Driver) (repository.Person, error) {
	switch driver {
	case Postgres:
		return repository.NewPsqlPerson(db), nil
	default:
		return nil, fmt.Errorf("Driver not implemented")
	}
}
