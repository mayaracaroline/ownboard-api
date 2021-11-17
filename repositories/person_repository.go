package repositories

import (
	"errors"

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

func NewPersonRepositoryDb(db map[string]model.Person) Repository {
	return &personRepository{
		db: db,
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

func (r *personRepository) Update(p model.Person) error {
	existsPerson := r.checkForExistingPerson(p.Document)

	if !existsPerson {
		return errors.New("Pessoa não encontrada para atualização")
	}

	r.db[p.Document] = p

	return nil
}

func (r *personRepository) FindByDocument(document string) (model.Person, error) {

	person, ok := r.db[document]

	if !ok {
		return model.Person{}, errors.New("Pessoa não encontrada para o documento: " + document)
	}
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
