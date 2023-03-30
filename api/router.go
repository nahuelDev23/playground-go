package api

import (
	"net/http"

	"github.com/nahueldev23/playground/repository"
	"github.com/nahueldev23/playground/services"
)

func RoutePerson(mux *http.ServeMux, storage repository.Person) {

	p := services.NewPerson(storage)

	mux.HandleFunc("/v1/user/create", p.Create)
	mux.HandleFunc("/v1/user/migrate", p.Migrate)
	mux.HandleFunc("/v1/user/get-all", p.GetAll)
}

func RouteLogin(mux *http.ServeMux, storage repository.Login) {

	l := services.NewLogin(storage)

	mux.HandleFunc("/v1/login", l.SignIn)
}
