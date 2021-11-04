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

func NewPersonHandle() Handler {
	s := service.NewPersonService()
	return &personHandler{
		service: s,
	}
}

func (h *personHandler) GetPersonByDocument(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get person by document: /person/{document}")
	params := mux.Vars(r)
	person := h.service.GetPersonByDocument(params["document"])
	json.NewEncoder(w).Encode(&person)
}
func (h *personHandler) GetPersons(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get person: /person/{document}")
	persons := h.service.GetPersons()
	json.NewEncoder(w).Encode(&persons)
}

func (h *personHandler) CreatePerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create person: /person")
	message := h.service.CreatePerson(r)
	json.NewEncoder(w).Encode(&message)
}

func (h *personHandler) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update person: /person")
	message := h.service.UpdatePerson(r)
	json.NewEncoder(w).Encode(&message)
}

func (h *personHandler) DeletePerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete persons: /person")
	h.service.DeleteAllPersons()
}

func (h *personHandler) DeletePersonByDocument(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete person by document: /person")
	params := mux.Vars(r)
	h.service.DeletePersonById(params["document"])
}
