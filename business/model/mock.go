package model

import (
	"time"

	"github.com/stretchr/testify/mock"
)

type PersonMock struct {
	mock.Mock
}

func (p *PersonMock) NewPersonMock() *Person {
	return &Person{
		Name:       "Mayara",
		LastName:   "Santos",
		DateOfBorn: time.Date(1994, 04, 20, 18, 0, 0, 0, time.UTC),
		Type:       "Fisica",
		Document:   "44221617845",
	}
}
