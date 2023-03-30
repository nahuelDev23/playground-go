package main

import (
	"log"
	"net/http"

	//"github.com/nahueldev23/playground/auth"
	"github.com/nahueldev23/playground/api"
	"github.com/nahueldev23/playground/database"
)

func main() {

	//	err := auth.LoadFiles("app.rsa", "app.rsa.pub")
	//if err != nil {
	//	log.Printf("error el cargar los RSA , %v \n", err)
	//}

	driver := database.Postgres
	database.New(driver)
	store, err := database.DAOPerson(driver)
	if err != nil {
		log.Printf("error en DAO Person, %v \n", err)
	}
	storeLogin, err := database.DAOLogin(driver)

	if err != nil {
		log.Printf("error en DAO Login, %v \n", err)
	}

	mux := http.NewServeMux()
	api.RoutePerson(mux, store)
	api.RouteLogin(mux, storeLogin)

	err = http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Printf("error en el servidor, %v \n", err)
	}

}
