package service

import (
	"encoding/json"
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
	person := toPerson(r)

	if s.repository.FindById(person.Document) != (model.Person{}) {
		return "Usuário já cadastrado!"
	}

	s.repository.Save(person)

	return "Usuário cadastrado com sucesso!"
}

func (s *personService) UpdatePerson(r *http.Request) string {
	person := toPerson(r)
	existsPerson := s.repository.CheckForExistingPerson(person.Document)

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

func toPerson(r *http.Request) model.Person {
	var person model.Person
	json.NewDecoder(r.Body).Decode(&person)
	return person
}
