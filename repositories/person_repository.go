package repositories

import (
	"errors"

	"github.com/mercadolibre/api/business/model"
)

type personRepository struct {
	db map[string]model.Person
}

const ErrNotFound = "pessoa não encontrada"

func NewPersonRepository() Repository {
	p := make(map[string]model.Person)
	return &personRepository{
		db: p,
	}
}

func (r *personRepository) Save(p model.Person) error {
	r.db[p.Document] = p
	return nil
}

func (r *personRepository) FindAll() []model.Person {
	_persons := []model.Person{}

	for _, p := range r.db {
		_persons = append(_persons, p)
	}
	return _persons
}

func (r *personRepository) Update(p model.Person) error {
	existsPerson := r.checkForExistingPerson(p.Document)
	if !existsPerson {
		return errors.New(ErrNotFound)
	}

	r.db[p.Document] = p

	return nil
}

func (r *personRepository) FindByDocument(document string) (model.Person, error) {
	person := r.db[document]

	return person, nil
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
