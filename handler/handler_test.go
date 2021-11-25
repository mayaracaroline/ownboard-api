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
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/mock"
)

func TestGetPersonByDocumentError(t *testing.T) {

	s := service.PersonServiceMock{}
	handler := handler.NewPersonHandle(&s)
	s.On("GetPersonByDocument", mock.AnythingOfType("string")).Return(model.Person{}, errors.New(""))

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

func TestGetPersonByDocument(t *testing.T) {

	s := service.PersonServiceMock{}
	handler := handler.NewPersonHandle(&s)
	s.On("GetPersonByDocument", mock.AnythingOfType("string")).Return(model.Person{}, nil)

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

	s := service.PersonServiceMock{}
	handler := handler.NewPersonHandle(&s)
	s.On("GetPersons").Return([]model.Person{})

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

	s := service.PersonServiceMock{}
	handler := handler.NewPersonHandle(&s)
	s.On("CreatePerson", mock.AnythingOfType("*http.Request")).Return(nil)

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

func TestCreatePersonError(t *testing.T) {

	s := service.PersonServiceMock{}
	handler := handler.NewPersonHandle(&s)
	s.On("CreatePerson", mock.AnythingOfType("*http.Request")).Return(errors.New(""))

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

	s := service.PersonServiceMock{}
	handler := handler.NewPersonHandle(&s)
	s.On("UpdatePerson", mock.AnythingOfType("*http.Request")).Return(nil)

	r := mux.NewRouter()
	r.HandleFunc("/person", handler.UpdatePerson)
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

	s := service.PersonServiceMock{}
	handler := handler.NewPersonHandle(&s)
	s.On("DeletePersonByDocument", mock.AnythingOfType("string"))

	r := mux.NewRouter()
	r.HandleFunc("/person/{document}", handler.DeletePersonByDocument)
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

	s := service.PersonServiceMock{}
	handler := handler.NewPersonHandle(&s)
	s.On("DeleteAllPersons")

	r := mux.NewRouter()
	r.HandleFunc("/person", handler.DeletePerson)
	ts := httptest.NewServer(r)
	defer ts.Close()
	apitest.New().
		Handler(r).
		Delete("/person").
		Expect(t).
		Status(http.StatusOK).
		End()
}
