package repositories

import "github.com/mercadolibre/api/business/model"

type Repository interface {
	Save(person model.Person)
	Update(person model.Person)
	FindAll() []model.Person
	FindById(id string) model.Person
	CheckForExistingPerson(id string) bool
	DeleteById(id string)
	DeleteAll()
}
