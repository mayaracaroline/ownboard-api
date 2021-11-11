package repositories

import "github.com/mercadolibre/api/business/model"

type Repository interface {
	Save(person model.Person)
	Update(person model.Person) error
	FindAll() []model.Person
	FindByDocument(id string) (model.Person, error)
	checkForExistingPerson(id string) bool
	DeleteByDocument(id string)
	DeleteAll()
}
