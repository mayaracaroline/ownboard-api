package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mercadolibre/api/business/service"
)

type (
	Handler interface {
		GetPersonByDocument(w http.ResponseWriter, r *http.Request)
		GetPersons(w http.ResponseWriter, r *http.Request)
		CreatePerson(w http.ResponseWriter, r *http.Request)
		UpdatePerson(w http.ResponseWriter, r *http.Request)
		DeletePerson(w http.ResponseWriter, r *http.Request)
		DeletePersonByDocument(w http.ResponseWriter, r *http.Request)
	}
	personHandler struct {
		service service.Service
	}
)

func NewPersonHandle(s service.Service) Handler {
	return &personHandler{
		service: s,
	}
}

func (h personHandler) GetPersonByDocument(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get person by document: /person/{document}")
	params := mux.Vars(r)
	person, message := h.service.GetPersonByDocument(params["document"])
	var err error

	if message != nil {
		err = json.NewEncoder(w).Encode(&message)
	} else {
		err = json.NewEncoder(w).Encode(&person)
	}

	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func (h personHandler) GetPersons(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get person: /person/{document}")
	persons := h.service.GetPersons()
	err := json.NewEncoder(w).Encode(&persons)

	if err != nil {
		fmt.Println("Erro ao buscar dados em service.GetPersons ", err.Error())
		http.Error(w, "erro ao buscar pessoa: "+err.Error(), 500)

	}
}

func (h personHandler) CreatePerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create person: /person")

	message := h.service.CreatePerson(r)
	err := json.NewEncoder(w).Encode(&message)

	if err != nil {
		fmt.Println("Erro ao processar dados service.CreatePerson", err.Error())
		http.Error(w, err.Error(), 500)
	}

}

func (h personHandler) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update person: /person")
	message := h.service.UpdatePerson(r)
	err := json.NewEncoder(w).Encode(&message)

	if err != nil {
		fmt.Println("Erro ao processar dados service.UpdatePerson", err.Error())
		http.Error(w, err.Error(), 500)
	}
}

func (h personHandler) DeletePerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete persons: /person")
	h.service.DeleteAllPersons()
}

func (h personHandler) DeletePersonByDocument(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete person by document: /person")
	params := mux.Vars(r)
	h.service.DeletePersonByDocument(params["document"])
}
