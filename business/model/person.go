package model

import "time"

type (
	Person struct {
		Name       string    `json:"name"`
		LastName   string    `json:"lastName"`
		DateOfBorn time.Time `json:"dateOfBorn"`
		Document   string    `json:"document"`
		Tipo       string    `json:"tipo"`
	}
)

func NewPerson(name string,
	lastName string,
	dateOfBorn time.Time,
	document string,
	tipo string) *Person {
	return &Person{
		Name:       name,
		LastName:   lastName,
		DateOfBorn: dateOfBorn,
		Document:   document,
		Tipo:       tipo,
	}
}
