package services

import (
	"encoding/json"
	"fmt"
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
	data := models.Person{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("error al decodificar el bory %v", err)
	}
	err = p.storage.Create(&data)

	if err != nil {
		fmt.Printf("%v", err)
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
