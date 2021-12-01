package service_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/mercadolibre/api/business/model"
	"github.com/mercadolibre/api/business/service"
	"github.com/mercadolibre/api/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	_registeredDocument   = "44221617845"
	_unregisteredDocument = "12345678900"
)

func TestGetPersons(t *testing.T) {

	repositoryMock := repositories.PersonRepositoryMock{}
	s := service.NewPersonService(&repositoryMock)
	person := *model.NewPerson(
		"Mayara",
		"Santos",
		time.Date(1994, 04, 20, 18, 0, 0, 0, time.UTC),
		_registeredDocument,
		"Fisica",
	)

	mock := make([]model.Person, 1)
	mock = append(mock, person)
	repositoryMock.On("FindAll").Return(mock)
	persons := s.GetPersons()
	assert.NotEmpty(t, persons)

}
func TestGetPersonByDocument(t *testing.T) {

	repositoryMock := repositories.PersonRepositoryMock{}
	s := service.NewPersonService(&repositoryMock)
	person := *model.NewPerson(
		"Mayara",
		"Santos",
		time.Date(1994, 04, 20, 18, 0, 0, 0, time.UTC),
		_registeredDocument,
		"Fisica",
	)

	repositoryMock.On("FindByDocument", mock.AnythingOfType("string")).Return(person, errors.New(""))
	person, message := s.GetPersonByDocument(_registeredDocument)
	assert.NotEmpty(t, person)
	assert.Empty(t, message)

}

func TestGetPersonByDocumentError(t *testing.T) {
	repositoryMock := repositories.PersonRepositoryMock{}
	s := service.NewPersonService(&repositoryMock)
	person := *model.NewPerson(
		"Mayara",
		"Santos",
		time.Date(1994, 04, 20, 18, 0, 0, 0, time.UTC),
		_registeredDocument,
		"Fisica",
	)

	repositoryMock.On("FindByDocument", mock.AnythingOfType("string")).Return(model.Person{}, nil)
	person, message := s.GetPersonByDocument(_unregisteredDocument)
	assert.Empty(t, person)
	assert.NotEmpty(t, message)

}

func TestGetPersonByDocumentNotFound(t *testing.T) {
	repositoryMock := repositories.PersonRepositoryMock{}
	s := service.NewPersonService(&repositoryMock)
	person := *model.NewPerson(
		"Mayara",
		"Santos",
		time.Date(1994, 04, 20, 18, 0, 0, 0, time.UTC),
		_registeredDocument,
		"Fisica",
	)

	repositoryMock.On("FindByDocument", mock.AnythingOfType("string")).Return(model.Person{}, errors.New("Error"))
	person, message := s.GetPersonByDocument(_unregisteredDocument)
	assert.Empty(t, person)
	assert.NotEmpty(t, message)

}

func TestCreatePerson(t *testing.T) {
	repositoryMock := repositories.PersonRepositoryMock{}
	s := service.NewPersonService(&repositoryMock)
	person := *model.NewPerson(
		"Joao",
		"Silva",
		time.Date(1997, 01, 19, 18, 0, 0, 0, time.UTC),
		"44321517843",
		"Fisica",
	)
	repositoryMock.On("FindByDocument", mock.AnythingOfType("string")).Return(model.Person{}, nil)
	repositoryMock.On("Save", mock.AnythingOfType("model.Person")).Return(nil)

	request, err := createRequest(person)
	message := s.CreatePerson(request)

	assert.NoError(t, err)
	assert.Nil(t, message)
}

func TestCreatePersonError(t *testing.T) {
	repositoryMock := repositories.PersonRepositoryMock{}
	s := service.NewPersonService(&repositoryMock)
	person := *model.NewPerson(
		"Mayara",
		"Santos",
		time.Date(1994, 04, 20, 18, 0, 0, 0, time.UTC),
		_registeredDocument,
		"Fisica",
	)
	repositoryMock.On("FindByDocument", mock.AnythingOfType("string")).Return(person, nil)
	request, err := createRequest(person)
	message := s.CreatePerson(request)

	assert.NoError(t, err)
	assert.Error(t, message)
}

func TestCreatePersonRequestError(t *testing.T) {
	repositoryMock := repositories.PersonRepositoryMock{}
	s := service.NewPersonService(&repositoryMock)

	repositoryMock.On("FindByDocument", mock.AnythingOfType("string")).Return(model.Person{}, nil)
	request, err := createRequestError()
	message := s.CreatePerson(request)

	assert.NoError(t, err)
	assert.Error(t, message)
}

func TestUpdatePerson(t *testing.T) {
	repositoryMock := repositories.PersonRepositoryMock{}
	s := service.NewPersonService(&repositoryMock)
	person := *model.NewPerson(
		"Mayara",
		"de Paula",
		time.Date(1994, 04, 20, 18, 0, 0, 0, time.UTC),
		_registeredDocument,
		"Fisica",
	)
	repositoryMock.On("FindByDocument", mock.AnythingOfType("string")).Return(person, nil)
	repositoryMock.On("Update", mock.AnythingOfType("model.Person")).Return(nil)
	request, err := createRequest(person)
	message := s.UpdatePerson(request)

	assert.NoError(t, err)
	assert.Nil(t, message)
}

func TestUpdatePerson_Data(t *testing.T) {
	repositoryMock := repositories.PersonRepositoryMock{}
	s := service.NewPersonService(&repositoryMock)
	p := *model.NewPerson(
		"Mayara",
		"de Paula",
		time.Date(1994, 04, 20, 18, 0, 0, 0, time.UTC),
		_registeredDocument,
		"Fisica",
	)

	repositoryMock.On("FindByDocument", mock.AnythingOfType("string")).Return(p, nil)
	repositoryMock.On("Update", mock.AnythingOfType("model.Person")).Return(nil)
	request, err := createRequest(p)
	s.UpdatePerson(request)

	person, _ := s.GetPersonByDocument(_registeredDocument)
	expected := "de Paula"
	actual := person.LastName

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestUpdatePersonError(t *testing.T) {
	repositoryMock := repositories.PersonRepositoryMock{}
	s := service.NewPersonService(&repositoryMock)
	person := *model.NewPerson(
		"Mayara",
		"de Paula",
		time.Date(1994, 04, 20, 18, 0, 0, 0, time.UTC),
		_unregisteredDocument,
		"Fisica",
	)

	repositoryMock.On("FindByDocument", mock.AnythingOfType("string")).Return(model.Person{}, "Error")
	repositoryMock.On("Update", mock.AnythingOfType("model.Person")).Return(errors.New("Pessoa não encontrada para atualização"))
	request, err := createRequest(person)
	updateErr := s.UpdatePerson(request)

	assert.NoError(t, err)
	assert.Error(t, updateErr)
}

func TestUpdatePersonRequestError(t *testing.T) {
	repositoryMock := repositories.PersonRepositoryMock{}
	s := service.NewPersonService(&repositoryMock)

	request, err := createRequestError()
	updateErr := s.UpdatePerson(request)

	assert.NoError(t, err)
	assert.Error(t, updateErr)

}

func TestDeletePersonByDocument(t *testing.T) {
	repositoryMock := repositories.PersonRepositoryMock{}
	repositoryMock.On("DeleteByDocument", mock.AnythingOfType("string")).Return()
	repositoryMock.On("FindByDocument", mock.AnythingOfType("string")).Return(model.Person{}, errors.New("Error"))

	s := service.NewPersonService(&repositoryMock)
	s.DeletePersonByDocument(_registeredDocument)
	person, err := s.GetPersonByDocument(_registeredDocument)

	assert.NotEmpty(t, err)
	assert.Empty(t, person)
}

func TestDeleteAllPerson(t *testing.T) {

	repositoryMock := repositories.PersonRepositoryMock{}
	s := service.NewPersonService(&repositoryMock)

	p := *model.NewPerson(
		"Jose",
		"Camargo",
		time.Now(),
		"33871752053",
		"Fisica",
	)

	repositoryMock.On("DeleteByDocument", mock.AnythingOfType("string")).Return()
	repositoryMock.On("FindByDocument", mock.AnythingOfType("string")).Return(model.Person{}, errors.New("Error"))
	repositoryMock.On("FindAll", mock.AnythingOfType("string")).Return([]model.Person{}, nil)
	repositoryMock.On("DeleteAll", mock.AnythingOfType("string")).Return()
	repositoryMock.On("Save", mock.AnythingOfType("model.Person"))

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

func createRequestError() (*http.Request, error) {
	request, err := http.NewRequest(http.MethodPost, "/", bytes.NewBufferString("{ name:}"))
	return request, err
}
