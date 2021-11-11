package tests

import (
	"testing"
	"time"

	"github.com/mercadolibre/api/business/model"
	"github.com/mercadolibre/api/repositories"
	"github.com/stretchr/testify/assert"
)

var repo repositories.Repository = repositories.NewPersonRepository()

func init() {
	repo.Save(*model.NewPerson("Mayara", "Santos", time.Date(1994, 04, 20, 18, 0, 0, 0, time.UTC), "44221617845", "Fisica"))
}

func TestFindPersonByDocument(t *testing.T) {
	personMock := &PersonMock{}

	person, err := repo.FindByDocument("44221617845")

	assert.Equal(t, personMock.NewPersonMock(), person)
	assert.NotEmpty(t, person)
	assert.NoError(t, err)
}

func TestFindPersonByDocumentError(t *testing.T) {
	personMock := &PersonMock{}

	person, err := repo.FindByDocument("44221617840")

	assert.NotEqual(t, personMock.NewPersonMock(), person)
	assert.Empty(t, person)
	assert.Error(t, err)
}

func TestFindPersonAll(t *testing.T) {
	repo.Save(*model.NewPerson("Mayara", "Santos", time.Now(), "76858631038", "Fisica"))
	persons := repo.FindAll()

	assert.NotEmpty(t, persons)
}

func TestSavePerson(t *testing.T) {
	repo.Save(*model.NewPerson("Mayara", "Santos", time.Now(), "76858631038", "Fisica"))

	person, err := repo.FindByDocument("76858631038")
	assert.NotEmpty(t, person)
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	err := repo.Update(*model.NewPerson("Mayara", "de Paula", time.Now(), "44221617845", "Fisica"))
	assert.NoError(t, err)
}

func TestUpdate_Data(t *testing.T) {
	repo.Update(*model.NewPerson("Mayara", "de Paula", time.Now(), "44221617845", "Fisica"))

	person, err := repo.FindByDocument("44221617845")
	expected := "de Paula"
	actual := person.LastName

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestUpdatError(t *testing.T) {
	err := repo.Update(*model.NewPerson("Mayara", "Santos", time.Now(), "1234567890", "Fisica"))

	assert.Error(t, err)
}
func TestDeletePerson(t *testing.T) {
	repo.DeleteByDocument("76858631038")
	person, err := repo.FindByDocument("76858631038")
	assert.Empty(t, person)
	assert.Error(t, err)
}

func TestDeleteAll(t *testing.T) {
	repo.DeleteAll()

	persons := repo.FindAll()
	assert.Empty(t, persons)
}
