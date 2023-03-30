package main

import (
	"log"
	"net/http"

	"github.com/nahueldev23/playground/api"
	"github.com/nahueldev23/playground/database"
)

func main() {
	driver := database.Postgres
	database.New(driver)
	store, err := database.DAOPerson(driver)

	if err != nil {
		log.Printf("error en daoperson, %v \n", err)
	}

	mux := http.NewServeMux()
	api.RoutePerson(mux, store)

	err = http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Printf("error en el servidor, %v \n", err)
	}

}
