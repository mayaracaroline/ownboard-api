package repositories

import (
	"github.com/mercadolibre/api/business/model"
	"github.com/stretchr/testify/mock"
)

type PersonRepositoryMock struct {
	mock.Mock
}

func (r *PersonRepositoryMock) Save(person model.Person) error {
	args := r.Called(person)
	return args.Error(0)
}

func (r *PersonRepositoryMock) Update(person model.Person) error {
	args := r.Called(person)
	return args.Error(0)
}

func (r *PersonRepositoryMock) FindAll() []model.Person {
	args := r.Called()
	return args.Get(0).([]model.Person)
}
func (r *PersonRepositoryMock) FindByDocument(id string) (model.Person, error) {
	args := r.Called(id)
	return args.Get(0).(model.Person), args.Error(1)
}

func (r *PersonRepositoryMock) DeleteByDocument(id string) {
	r.Called(id)
}
func (r *PersonRepositoryMock) DeleteAll() {
	r.Called()
}
