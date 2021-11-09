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

func (r *personRepository) Update(p model.Person) bool {
	existsPerson := r.checkForExistingPerson(p.Document)

	if existsPerson {
		r.db[p.Document] = p
	}
	return existsPerson
}

func (r *personRepository) FindByDocument(id string) (bool, model.Person) {

	person, ok := r.db[id]

	return ok, person
}

func (r *personRepository) DeleteByDocument(id string) {
	delete(r.db, id)
}

func (r *personRepository) DeleteAll() {
	r.db = make(map[string]model.Person)
}

func (r *personRepository) checkForExistingPerson(id string) bool {
	_, ok := r.db[id]

	return ok
}
