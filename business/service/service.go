package service

import (
	"net/http"

	"github.com/mercadolibre/api/business/model"
)

type Service interface {
	GetPersons() []model.Person
	GetPersonByDocument(document string) model.Person
	CreatePerson(r *http.Request) string
	UpdatePerson(r *http.Request) string
	DeletePersonByDocument(documento string)
	DeleteAllPersons()
}
