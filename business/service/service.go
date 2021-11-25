package service

import (
	"net/http"

	"github.com/mercadolibre/api/business/model"
)

type Service interface {
	GetPersons() []model.Person
	GetPersonByDocument(document string) (model.Person, error)
	CreatePerson(r *http.Request) error
	UpdatePerson(r *http.Request) error
	DeletePersonByDocument(document string)
	DeleteAllPersons()
}
