package model

import "time"

type (
	Person struct {
		Name       string
		LastName   string
		DateOfBorn time.Time
		Document   string
		Tipo       string
	}
)

func NewPerson(name string,
	lastName string,
	dateOfBorn time.Time,
	document string) *Person {
	return &Person{
		Name:       name,
		LastName:   lastName,
		DateOfBorn: dateOfBorn,
		Document:   document,
	}
}

func (person *Person) setName(name string) {
	person.Name = name
}
