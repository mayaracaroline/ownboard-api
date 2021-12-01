package service

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/mercadolibre/api/business/model"
	"github.com/mercadolibre/api/repositories"
)

type personService struct {
	repository repositories.Repository
}

func NewPersonService(r repositories.Repository) Service {
	return &personService{
		repository: r,
	}
}
func (s personService) GetPersons() []model.Person {
	return s.repository.FindAll()
}

func (s personService) GetPersonByDocument(document string) (model.Person, error) {

	person, findErr := s.repository.FindByDocument(document)
	if findErr != nil {
		return person, findErr
	} else if person.Document == "" {
		return person, errors.New("pessoa não encontrada para este documento")
	}

	return person, nil
}

func (s personService) CreatePerson(r *http.Request) error {
	person, err := toPerson(r)

	if err != nil {
		return err
	}

	p, findErr := s.repository.FindByDocument(person.Document)

	if findErr != nil {
		return errors.New("erro ao criar pessoa " + findErr.Error())
	} else if p.Document != "" {
		return errors.New("pessoa já cadastrada para este documento")
	}

	return s.repository.Save(person)

}

func (s personService) UpdatePerson(r *http.Request) error {
	person, err := toPerson(r)
	if err != nil {
		return errors.New(err.Error())
	}

	return s.repository.Update(person)

}

func (s personService) DeletePersonByDocument(document string) {
	s.repository.DeleteByDocument(document)
}

func (s personService) DeleteAllPersons() {
	s.repository.DeleteAll()
}

func toPerson(r *http.Request) (model.Person, error) {
	var person model.Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		return person, errors.New("erro ao processar dados, revise os campos inseridos")
	}
	return person, nil
}
