package entity

import "errors"

type Movie struct {
	ID       string
	Title    string
	Director string
	Writer   string
	Link     string
	Duration float64
}

func NewMovie(id, title, director, writer, link string, duration float64) (*Movie, error) {
	movie := &Movie{
		ID:       id,
		Title:    title,
		Director: director,
		Writer:   writer,
		Link:     link,
		Duration: duration,
	}
	if err := movie.IsValid(); err != nil {
		return nil, err
	}
	return movie, nil
}

func (m *Movie) IsValid() error {
	if m.ID == "" {
		return errors.New("invalid id")
	}
	if m.Title == "" {
		return errors.New("invalid title")
	}
	if m.Director == "" {
		return errors.New("invalid director")
	}
	if m.Writer == "" {
		return errors.New("invalid writer")
	}
	if m.Link == "" {
		return errors.New("invalid link")
	}
	if m.Duration <= 0 {
		return errors.New("invalid duration")
	}
	return nil
}
