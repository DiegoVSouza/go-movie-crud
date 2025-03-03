package database

import (
	"database/sql"
	"errors"
	"fmt"
	"monte_clone_go/internal/entity"
	"strings"
)

type MovieRepository struct {
	Db *sql.DB
}

func NewMovieRepository(db *sql.DB) *MovieRepository {
	return &MovieRepository{Db: db}
}

func (r *MovieRepository) Save(movie *entity.Movie) error {
	stmt, err := r.Db.Prepare(`
		INSERT INTO movies (id, title, director, writer, link, duration) 
		VALUES (?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(movie.ID, movie.Title, movie.Director, movie.Writer, movie.Link, movie.Duration)
	return err
}

func (r *MovieRepository) Get(filters map[string]string) ([]*entity.Movie, error) {
	var conditions []string
	var args []interface{}

	for key, value := range filters {
		conditions = append(conditions, fmt.Sprintf("%s LIKE ?", key))
		args = append(args, "%"+value+"%")
	}

	query := "SELECT id, title, director, writer, link, duration FROM movies"
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	rows, err := r.Db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*entity.Movie
	for rows.Next() {
		movie := &entity.Movie{}
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Director,
			&movie.Writer,
			&movie.Link,
			&movie.Duration,
		)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return movies, nil
}

func (r *MovieRepository) FindByID(id string) (*entity.Movie, error) {
	row := r.Db.QueryRow(`
		SELECT id, title, director, writer, link, duration 
		FROM movies 
		WHERE id = ?
	`, id)

	var movie entity.Movie
	err := row.Scan(&movie.ID, &movie.Title, &movie.Director, &movie.Writer, &movie.Link, &movie.Duration)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &movie, nil
}

func (r *MovieRepository) Update(movie *entity.Movie) error {
	stmt, err := r.Db.Prepare(`
		UPDATE movies 
		SET title = ?, director = ?, writer = ?, link = ?, duration = ?
		WHERE id = ?
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(movie.Title, movie.Director, movie.Writer, movie.Link, movie.Duration, movie.ID)
	return err
}

func (r *MovieRepository) Delete(id string) error {
	stmt, err := r.Db.Prepare(`DELETE FROM movies WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}

func (r *MovieRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("SELECT count(*) FROM movies").Scan(&total)
	return total, err
}
