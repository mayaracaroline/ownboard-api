package service_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/mercadolibre/api/business/model"
	"github.com/mercadolibre/api/business/service"
	"github.com/mercadolibre/api/repositories"
	"github.com/stretchr/testify/assert"
)

var s service.Service

const (
	_registeredDocument   = "44221617845"
	_unregisteredDocument = "12345678900"
)

func init() {
	var repo repositories.Repository = repositories.NewPersonRepository()
	person := *model.NewPerson(
		"Mayara",
		"Santos",
		time.Date(1994, 04, 20, 18, 0, 0, 0, time.UTC),
		_registeredDocument,
		"Fisica",
	)

	repo.Save(person)
	s = service.NewPersonService(repo)
}

func TestGetPersons(t *testing.T) {

	persons := s.GetPersons()
	assert.NotEmpty(t, persons)

}
func TestGetPersonByDocument(t *testing.T) {

	person, message := s.GetPersonByDocument(_registeredDocument)
	assert.NotEmpty(t, person)
	assert.Empty(t, message)

}

func TestGetPersonByDocumentError(t *testing.T) {

	person, message := s.GetPersonByDocument(_unregisteredDocument)
	assert.Empty(t, person)
	assert.NotEmpty(t, message)

}

func TestCreatePerson(t *testing.T) {
	person := *model.NewPerson(
		"Joao",
		"Silva",
		time.Date(1997, 01, 19, 18, 0, 0, 0, time.UTC),
		"44321517843",
		"Fisica",
	)
	request, err := createRequest(person)
	message := s.CreatePerson(request)

	assert.NoError(t, err)
	assert.Equal(t, "Pessoa cadastrada com sucesso!", message)
}

func TestCreatePersonError(t *testing.T) {
	person := *model.NewPerson(
		"Mayara",
		"Santos",
		time.Date(1994, 04, 20, 18, 0, 0, 0, time.UTC),
		_registeredDocument,
		"Fisica",
	)
	request, err := createRequest(person)
	message := s.CreatePerson(request)

	assert.NoError(t, err)
	assert.Equal(t, "Pessoa já cadastrada!", message)
}

func TestUpdatePerson(t *testing.T) {
	person := *model.NewPerson(
		"Mayara",
		"de Paula",
		time.Date(1994, 04, 20, 18, 0, 0, 0, time.UTC),
		_registeredDocument,
		"Fisica",
	)
	request, err := createRequest(person)
	message := s.UpdatePerson(request)

	assert.NoError(t, err)
	assert.Equal(t, "Dados atualizados com sucesso", message)
}

func TestUpdatePerson_Data(t *testing.T) {
	p := *model.NewPerson(
		"Mayara",
		"de Paula",
		time.Date(1994, 04, 20, 18, 0, 0, 0, time.UTC),
		_registeredDocument,
		"Fisica",
	)
	request, err := createRequest(p)
	s.UpdatePerson(request)

	person, _ := s.GetPersonByDocument(_registeredDocument)
	expected := "de Paula"
	actual := person.LastName

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestUpdatePersonError(t *testing.T) {
	person := *model.NewPerson(
		"Mayara",
		"de Paula",
		time.Date(1994, 04, 20, 18, 0, 0, 0, time.UTC),
		_unregisteredDocument,
		"Fisica",
	)
	request, err := createRequest(person)
	message := s.UpdatePerson(request)

	assert.NoError(t, err)
	assert.Equal(t, "Pessoa não encontrada para atualização", message)
}

func TestDeletePersonByDocument(t *testing.T) {

	s.DeletePersonByDocument(_registeredDocument)
	person, err := s.GetPersonByDocument(_registeredDocument)

	assert.NotEmpty(t, err)
	assert.Empty(t, person)
}

func TestDeleteAllPerson(t *testing.T) {
	p := *model.NewPerson(
		"Jose",
		"Camargo",
		time.Now(),
		"33871752053",
		"Fisica",
	)
	request, err := createRequest(p)
	s.CreatePerson(request)

	s.DeleteAllPersons()
	person := s.GetPersons()

	assert.Nil(t, err)
	assert.Empty(t, person)
}

func createRequest(p model.Person) (*http.Request, error) {
	jsonBytes, _ := json.Marshal(p)
	request, err := http.NewRequest(http.MethodPost, "/", bytes.NewReader(jsonBytes))
	return request, err
}
