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
		DeletePersonById(w http.ResponseWriter, r *http.Request)
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
	params := mux.Vars(r)
	person := h.service.GetPersonByDocument(params["document"])
	json.NewEncoder(w).Encode(&person)
	fmt.Println("Get person by document: /person/{document}")
}
func (h *personHandler) GetPersons(w http.ResponseWriter, r *http.Request) {
	persons := h.service.GetPersons()
	json.NewEncoder(w).Encode(&persons)
	fmt.Println("Get person: /person/{document}")
}

func (h *personHandler) CreatePerson(w http.ResponseWriter, r *http.Request) {
	message := struct{ Message string }{h.service.CreatePerson(r)}
	json.NewEncoder(w).Encode(&message)
	fmt.Println("Create person: /person")
}

func (h *personHandler) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	message := h.service.UpdatePerson(r)
	json.NewEncoder(w).Encode(&message)
	fmt.Println("Update person: /person")
}

func (h *personHandler) DeletePerson(w http.ResponseWriter, r *http.Request) {
	h.service.DeletePersonById(r)
	fmt.Println("Delete persons: /person")
}

func (h *personHandler) DeletePersonById(w http.ResponseWriter, r *http.Request) {
	h.service.DeletePersonById(r)
	fmt.Println("Delete person by document: /person")
}
