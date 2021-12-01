package model

import "time"

type (
	Person struct {
		Name       string    `json:"name"`
		LastName   string    `json:"lastName"`
		DateOfBorn time.Time `json:"dateOfBorn"`
		Document   string    `json:"document"`
		Type       string    `json:"type"`
	}
)

func NewPerson(name string,
	lastName string,
	dateOfBorn time.Time,
	document string,
	t string) *Person {
	return &Person{
		Name:       name,
		LastName:   lastName,
		DateOfBorn: dateOfBorn,
		Document:   document,
		Type:       t,
	}
}
