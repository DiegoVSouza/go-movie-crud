package graph

import "monte_clone_go/internal/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CreateMovieUseCase usecase.CreateMovieUseCase
}
