package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyID_WhenCreateANewMovie_ThenShouldReceiveAnError(t *testing.T) {
	movie := Movie{}
	assert.Error(t, movie.IsValid(), "invalid id")
}

func TestGivenAValidParams_WhenICallNewMMovie_ThenIShouldReceiveCreateMovieWithAllParams(t *testing.T) {
	movie := Movie{
		ID:       "123",
		Title:    "Ainda estou aqui",
		Director: "Walter Salles",
		Writer:   "Marcelo Rubens Paiva",
		Link:     "link",
		Duration: 135,
	}

	assert.Equal(t, "123", movie.ID)
	assert.Equal(t, "Ainda estou aqui", movie.Title)
	assert.Equal(t, "Walter Salles", movie.Director)
	assert.Equal(t, "Marcelo Rubens Paiva", movie.Writer)
	assert.Equal(t, "link", movie.Link)
	assert.Nil(t, movie.IsValid())
}
