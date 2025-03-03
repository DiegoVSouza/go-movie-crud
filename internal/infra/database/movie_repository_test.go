package database

import (
	"database/sql"
	"testing"

	"monte_clone_go/internal/entity"

	"github.com/stretchr/testify/suite"

	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

type MovieRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *MovieRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)

	_, err = db.Exec(`CREATE TABLE movies (
		id VARCHAR(255) NOT NULL, 
		title VARCHAR(255) NOT NULL,
		director VARCHAR(255) NOT NULL, 
		writer VARCHAR(255) NOT NULL, 
		link VARCHAR(255) NOT NULL, 
		duration FLOAT NOT NULL,
		PRIMARY KEY (id)
	)`)
	suite.NoError(err)

	suite.Db = db
}

func (suite *MovieRepositoryTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(MovieRepositoryTestSuite))
}

func (suite *MovieRepositoryTestSuite) TestGivenAMovie_WhenSave_ThenShouldSaveMovie() {
	movie, err := entity.NewMovie("123", "Title", "Some Director", "Some Writer", "Some Link", 120)
	suite.NoError(err)

	repo := NewMovieRepository(suite.Db)

	err = repo.Save(movie)
	suite.NoError(err)

	var movieResult entity.Movie
	err = suite.Db.QueryRow(`
		SELECT id, title, director, writer, link, duration 
		FROM movies 
		WHERE id = ?`, movie.ID).
		Scan(&movieResult.ID, &movieResult.Title, &movieResult.Director, &movieResult.Writer, &movieResult.Link, &movieResult.Duration)

	suite.NoError(err)
	suite.Equal(movie.ID, movieResult.ID)
	suite.Equal(movie.Title, movieResult.Title)
	suite.Equal(movie.Director, movieResult.Director)
	suite.Equal(movie.Writer, movieResult.Writer)
	suite.Equal(movie.Link, movieResult.Link)
	suite.Equal(movie.Duration, movieResult.Duration)
}
