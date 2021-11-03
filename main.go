package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mercadolibre/api/business/service"
	"github.com/mercadolibre/api/handler"
)

var _handler = handler.NewPersonHandle()
var _service = service.NewPersonService()

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", GetPersons).Methods("GET")
	router.HandleFunc("/person", GetPersons).Methods("GET")
	router.HandleFunc("/person/{document}", GetPersonByDocument).Methods("GET")
	router.HandleFunc("/person", CreatePerson).Methods("POST")
	router.HandleFunc("/person", UpdatePerson).Methods("PUT")
	router.HandleFunc("/person/{document}", DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":10000", router))
}

func GetPersonByDocument(w http.ResponseWriter, r *http.Request) {
	_handler.GetPersonByDocument(w, r)
}
func GetPersons(w http.ResponseWriter, r *http.Request) {
	_handler.GetPersons(w, r)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	_handler.CreatePerson(w, r)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	_handler.UpdatePerson(w, r)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	_handler.DeletePerson(w, r)
}

func DeletePersonById(w http.ResponseWriter, r *http.Request) {
	_handler.DeletePersonById(w, r)
}
