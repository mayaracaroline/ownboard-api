package tests

import (
	"time"

	"github.com/mercadolibre/api/business/model"
	"github.com/stretchr/testify/mock"
)

type PersonMock struct {
	mock.Mock
}
type PersonRepositoryMock struct {
	mock.Mock
}

func (p *PersonMock) NewPersonMock() *model.Person {
	return &model.Person{
		Name:       "Mayara",
		LastName:   "Santos",
		DateOfBorn: time.Date(1994, 04, 20, 18, 0, 0, 0, time.UTC),
		Type:       "Fisica",
		Document:   "44221617845",
	}
}

func (r *PersonRepositoryMock) Save(person model.Person) {}

func (r *PersonRepositoryMock) Update(person model.Person) bool {
	args := r.Called(person)
	return args.Bool(0)
}

func (r *PersonRepositoryMock) FindAll() []model.Person {
	args := r.Called()
	return args.Get(0).([]model.Person)
}
func (r *PersonRepositoryMock) FindByDocument(id string) (bool, model.Person) {
	args := r.Called(id)
	return args.Bool(0), args.Get(1).(model.Person)
}

func (r *PersonRepositoryMock) checkForExistingPerson(id string) bool {
	args := r.Called(id)
	return args.Bool(0)
}

func (r *PersonRepositoryMock) DeleteByDocument(id string) {
	r.Called(id)
}
func (r *PersonRepositoryMock) DeleteAll() {
	r.Called()
}
