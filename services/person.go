package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/nahueldev23/playground/models"
)

func (p *person) Migrate(w http.ResponseWriter, r *http.Request) {
	err := p.storage.Migrate()
	if err != nil {
		log.Printf("no anda %v", err)
	}
	fmt.Printf("todo se migro")
}

func (p *person) Create(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		response := newResponse("error", "metodo no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	//esto deberia estar en un middleware
	// age := r.FormValue("age")

	data := models.Person{}
  //lo  convierto a un cadena de bytes
	body, err := ioutil.ReadAll(r.Body)
  

	if err != nil {
		log.Printf("error al decodificar el bory %v", err)
	}
  // la cadena de bytes body se la paso a la estructura de go 
  // ahora puedo acceder a data.name como si fuera un objeto
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("error al decodificar el bory %v", err)
	}

	// if !data.Details.Valid {
	// 	response := newResponse("error", "error la crear al usuario, details no puede estar vacia", nil)
	// 	responseJSON(w, http.StatusInternalServerError, response)
	// 	return
	// }

	err = p.storage.Create(&data)

	if err != nil {
		log.Printf("%v", err)
		response := newResponse("error", "error la crear al usuario", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse("ok", "Todo salio bien", &data)
	responseJSON(w, http.StatusOK, response)
}

func (p *person) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse("error", "metodo no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data, err := p.storage.GetAll()
	if err != nil {
		response := newResponse("error", "Error al obtener los registros de las personas", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return

	}

	response := newResponse("success", "Ok", data)
	responseJSON(w, http.StatusOK, response)

}
