package repositories

import "github.com/mercadolibre/api/business/model"

type Repository interface {
	Save(person model.Person) error
	Update(person model.Person) error
	FindAll() []model.Person
	FindByDocument(id string) (model.Person, error)
	DeleteByDocument(id string)
	DeleteAll()
}
