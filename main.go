package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mercadolibre/api/business/service"
	"github.com/mercadolibre/api/handler"
	"github.com/mercadolibre/api/repositories"
)

func main() {
	repository := repositories.NewPersonRepository()
	handler := handler.NewPersonHandle(bootstrapService(repository))
	router := mux.NewRouter()

	router.HandleFunc("/", handler.GetPersons).Methods("GET")
	router.HandleFunc("/persons", handler.GetPersons).Methods("GET")
	router.HandleFunc("/person/{document}", handler.GetPersonByDocument).Methods("GET")
	router.HandleFunc("/person", handler.CreatePerson).Methods("POST")
	router.HandleFunc("/person/{document}", handler.UpdatePerson).Methods("PUT")
	router.HandleFunc("/person/{document}", handler.DeletePersonByDocument).Methods("DELETE")
	router.HandleFunc("/persons", handler.DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":10000", router))
}

func bootstrapService(r repositories.Repository) service.Service {
	return service.NewPersonService(r)
}
