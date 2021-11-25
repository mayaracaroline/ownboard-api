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
	}

	return person, nil
}

func (s personService) CreatePerson(r *http.Request) error {
	person, err := toPerson(r)

	if err != nil {
		return err
	}

	_, findErr := s.repository.FindByDocument(person.Document)

	if findErr == nil {
		return errors.New("pessoa já cadastrada")
	}

	s.repository.Save(person)

	return nil
}

func (s personService) UpdatePerson(r *http.Request) error {
	person, err := toPerson(r)
	if err != nil {
		return errors.New(err.Error())
	}
	updateError := s.repository.Update(person)

	if updateError != nil {
		return updateError
	}

	return nil
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
