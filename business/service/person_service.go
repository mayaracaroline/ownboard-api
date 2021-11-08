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

func (s *personService) GetPersonByDocument(document string) model.Person {
	return s.repository.FindById(document)
}

func (s *personService) CreatePerson(r *http.Request) string {
	person, err := toPerson(r)

	if err != nil {
		return err.Error()
	}

	if s.repository.FindById(person.Document) != (model.Person{}) {
		return "Usuário já cadastrado!"
	}

	s.repository.Save(person)

	return "Usuário cadastrado com sucesso!"
}

func (s *personService) UpdatePerson(r *http.Request) string {
	person, err := toPerson(r)
	existsPerson := s.repository.CheckForExistingPerson(person.Document)
	if err != nil {
		return err.Error()
	}
	if !existsPerson {
		return "Usuário não encontrado para atualização"
	}

	s.repository.Update(person)
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
