package repositories

import "github.com/mercadolibre/api/business/model"

type Repository interface {
	Save(person model.Person)
	Update(person model.Person) bool
	FindAll() []model.Person
	FindByDocument(id string) (bool, model.Person)
	checkForExistingPerson(id string) bool
	DeleteByDocument(id string)
	DeleteAll()
}
