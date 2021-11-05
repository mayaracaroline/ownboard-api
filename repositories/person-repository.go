package repositories

import (
	"github.com/mercadolibre/api/business/model"
)

type personRepository struct {
	db map[string]model.Person
}

func NewPersonRepository() Repository {
	p := make(map[string]model.Person)
	return &personRepository{
		db: p,
	}
}

func (r *personRepository) Save(p model.Person) {
	r.db[p.Document] = p
}

func (r *personRepository) FindAll() []model.Person {
	_persons := []model.Person{}

	for _, p := range r.db {
		_persons = append(_persons, p)
	}
	return _persons
}

func (r *personRepository) Update(p model.Person) {
	r.db[p.Document] = p
}

func (r *personRepository) FindById(id string) model.Person {
	return r.db[id]
}

func (r *personRepository) DeleteByDocument(id string) {
	delete(r.db, id)
}

func (r *personRepository) CheckForExistingPerson(id string) bool {
	person := r.db[id]

	return person != (model.Person{})
}

func (r *personRepository) DeleteAll() {
	r.db = make(map[string]model.Person)
}
