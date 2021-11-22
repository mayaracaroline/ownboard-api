package handler_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/mercadolibre/api/business/model"
	"github.com/mercadolibre/api/business/service"
	"github.com/mercadolibre/api/handler"
	"github.com/mercadolibre/api/repositories"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/mock"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	Client = &http.Client{}
}
func TestGetPersonByDocument(t *testing.T) {

	repositoryMock := repositories.PersonRepositoryMock{}
	s := service.NewPersonService(&repositoryMock)
	handler := handler.NewPersonHandle(s)
	repositoryMock.On("FindByDocument", mock.AnythingOfType("string")).Return(model.Person{}, errors.New(""))

	r := mux.NewRouter()
	r.HandleFunc("/person/{document}", handler.GetPersonByDocument)
	ts := httptest.NewServer(r)
	defer ts.Close()
	apitest.New().
		Handler(r).
		Get("/person/01234567890").
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetPersons(t *testing.T) {

	repositoryMock := repositories.PersonRepositoryMock{}
	s := service.NewPersonService(&repositoryMock)
	handler := handler.NewPersonHandle(s)
	repositoryMock.On("FindAll", mock.AnythingOfType("string")).Return([]model.Person{}, errors.New(""))

	r := mux.NewRouter()
	r.HandleFunc("/persons", handler.GetPersons)
	ts := httptest.NewServer(r)
	defer ts.Close()
	apitest.New().
		Handler(r).
		Get("/persons").
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestCreatePerson(t *testing.T) {

	repositoryMock := repositories.PersonRepositoryMock{}
	s := service.NewPersonService(&repositoryMock)
	handler := handler.NewPersonHandle(s)
	repositoryMock.On("FindByDocument", mock.AnythingOfType("string")).Return(model.Person{}, errors.New(""))

	r := mux.NewRouter()
	r.HandleFunc("/person", handler.CreatePerson)
	ts := httptest.NewServer(r)
	defer ts.Close()
	apitest.New().
		Handler(r).
		Post("/person").
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestUpdatePerson(t *testing.T) {

	repositoryMock := repositories.PersonRepositoryMock{}
	s := service.NewPersonService(&repositoryMock)
	handler := handler.NewPersonHandle(s)

	r := mux.NewRouter()
	r.HandleFunc("/person", handler.CreatePerson)
	ts := httptest.NewServer(r)
	defer ts.Close()
	apitest.New().
		Handler(r).
		Put("/person").
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestDeletePersonByDocument(t *testing.T) {

	repositoryMock := repositories.PersonRepositoryMock{}
	s := service.NewPersonService(&repositoryMock)
	handler := handler.NewPersonHandle(s)

	r := mux.NewRouter()
	r.HandleFunc("/person/{document}", handler.CreatePerson)
	ts := httptest.NewServer(r)
	defer ts.Close()
	apitest.New().
		Handler(r).
		Delete("/person/01234567890").
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestDeletePerson(t *testing.T) {

	repositoryMock := repositories.PersonRepositoryMock{}
	s := service.NewPersonService(&repositoryMock)
	handler := handler.NewPersonHandle(s)

	r := mux.NewRouter()
	r.HandleFunc("/person", handler.CreatePerson)
	ts := httptest.NewServer(r)
	defer ts.Close()
	apitest.New().
		Handler(r).
		Delete("/person").
		Expect(t).
		Status(http.StatusOK).
		End()
}
