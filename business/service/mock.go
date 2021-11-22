package service

import (
	"net/http"

	"github.com/mercadolibre/api/business/model"
	"github.com/stretchr/testify/mock"
)

type PersonServiceMock struct {
	mock.Mock
}

func (s *PersonServiceMock) GetPersons() []model.Person {
	args := s.Called()
	return args.Get(0).([]model.Person)
}
func (s *PersonServiceMock) GetPersonByDocument(document string) (model.Person, string) {
	args := s.Called(document)
	return args.Get(0).(model.Person), args.String(1)
}
func (s *PersonServiceMock) CreatePerson(r *http.Request) string {
	args := s.Called(r)
	return args.String(0)
}
func (s *PersonServiceMock) UpdatePerson(r *http.Request) string {
	args := s.Called(r)
	return args.String(0)
}
func (s *PersonServiceMock) DeletePersonByDocument(documento string) {
	s.Called(documento)
}
func (s *PersonServiceMock) DeleteAllPersons() {
	s.Called()
}
