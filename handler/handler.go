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
	person, errMessage := h.service.GetPersonByDocument(params["document"])

	if errMessage != nil {
		http.Error(w, errMessage.Error(), 500)
	}

	if err := json.NewEncoder(w).Encode(&person); err != nil {
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

	errMsg := h.service.CreatePerson(r)
	if errMsg != nil {
		fmt.Println("Erro ao processar dados service.CreatePerson", errMsg.Error())
		err := json.NewEncoder(w).Encode(&errMsg)
		if err != nil {
			http.Error(w, err.Error(), 500)
		} else {
			http.Error(w, errMsg.Error(), 500)
		}
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"cadastro realizado com sucesso"}`))
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
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"operação realizada com sucesso"}`))
}

func (h personHandler) DeletePersonByDocument(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete person by document: /person")
	params := mux.Vars(r)
	h.service.DeletePersonByDocument(params["document"])
}
