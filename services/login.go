package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/nahueldev23/playground/models"
)

func (l *login) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse("error", "metodo no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	//esto deberia estar en un middleware
	// age := r.FormValue("age")

	data := models.Login{}
	//lo  convierto a un cadena de bytes
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("error al decodificar el bory en cadena de bytes %v", err)
	}
	// la cadena de bytes body se la paso a la estructura de go
	// ahora puedo acceder a data.name como si fuera un objeto
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("error al pasar la cadena de bytes a estructura  %v", err)
		return
	}

	isLogged, err := l.storage.SignIn(&data)

	if err != nil {
		log.Printf("error al logear %v", err)
		return
	}
	var token string
	if isLogged {
		log.Printf("todo liso")
		token, err = GenerateToken(&data)
		if err != nil {

			log.Printf(" erro al obtener el token %v", err)
		}
	}

	log.Printf(" %v", token)
	response := newResponse("ok", "se logeo", &token)
	responseJSON(w, http.StatusOK, response)

}
