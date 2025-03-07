// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Movie struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Director string  `json:"director"`
	Writer   string  `json:"writer"`
	Link     string  `json:"link"`
	Duration float64 `json:"duration"`
}

type MovieInput struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Director string  `json:"director"`
	Writer   string  `json:"writer"`
	Link     string  `json:"link"`
	Duration float64 `json:"duration"`
}

type Mutation struct {
}

type Query struct {
}
