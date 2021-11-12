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
func (s *personService) GetPersons() []model.Person {
	return s.repository.FindAll()
}

func (s *personService) GetPersonByDocument(document string) (model.Person, string) {

	person, findErr := s.repository.FindByDocument(document)
	if findErr != nil {
		return person, findErr.Error()
	}

	return person, ""
}

func (s *personService) CreatePerson(r *http.Request) string {
	person, err := toPerson(r)

	if err != nil {
		return err.Error()
	}

	_, findErr := s.repository.FindByDocument(person.Document)

	if findErr == nil {
		return "Pessoa já cadastrada!"
	}

	s.repository.Save(person)

	return "Pessoa cadastrada com sucesso!"
}

func (s *personService) UpdatePerson(r *http.Request) string {
	person, err := toPerson(r)
	if err != nil {
		return err.Error()
	}
	updateError := s.repository.Update(person)

	if updateError != nil {
		return updateError.Error()
	}

	return "Dados atualizados com sucesso"
}

func (s *personService) DeletePersonByDocument(document string) {
	s.repository.DeleteByDocument(document)
}

func (s *personService) DeleteAllPersons() {
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
