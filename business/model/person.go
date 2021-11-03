package model

type (
	Person struct {
		Name       string
		LastName   string
		DateOfBorn string
		Document   string
	}
)

func NewPerson(name string,
	lastName string,
	dateOfBorn string,
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
