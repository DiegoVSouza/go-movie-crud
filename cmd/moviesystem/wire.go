//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"monte_clone_go/internal/entity"
	"monte_clone_go/internal/event"
	"monte_clone_go/internal/infra/database"
	"monte_clone_go/internal/infra/web"
	"monte_clone_go/internal/usecase"
	"monte_clone_go/pkg/events"

	"github.com/google/wire"
)

var setMovieRepositoryDependency = wire.NewSet(
	database.NewMovieRepository,
	wire.Bind(new(entity.MovieRepositoryInterface), new(*database.MovieRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setMovieCreatedEvent = wire.NewSet(
	event.NewMovieCreated,
	wire.Bind(new(event.MovieCreatedEvent), new(*event.MovieCreated)),
)

var setMovieUpdatedEvent = wire.NewSet(
	event.NewMovieUpdated,
	wire.Bind(new(event.MovieUpdatedEvent), new(*event.MovieUpdated)),
)

var setMovieDeletedEvent = wire.NewSet(
	event.NewMovieDeleted,
	wire.Bind(new(event.MovieDeletedEvent), new(*event.MovieDeleted)),
)

// Função para criar o use case de criação de filme
func NewCreateMovieUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateMovieUseCase {
	wire.Build(
		setMovieRepositoryDependency,
		setMovieCreatedEvent,
		usecase.NewCreateMovieUseCase,
	)
	return &usecase.CreateMovieUseCase{}
}

func NewGetMovieUseCase(db *sql.DB) *usecase.GetMovieUseCase {
	wire.Build(
		setMovieRepositoryDependency,
		usecase.NewGetMovieUseCase,
	)
	return &usecase.GetMovieUseCase{}
}

func NewUpdateMovieUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.UpdateMovieUseCase {
	wire.Build(
		setMovieRepositoryDependency,
		setMovieUpdatedEvent,
		usecase.NewUpdateMovieUseCase,
	)
	return &usecase.UpdateMovieUseCase{}
}

func NewDeleteMovieUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.DeleteMovieUseCase {
	wire.Build(
		setMovieRepositoryDependency,
		setMovieDeletedEvent,
		usecase.NewDeleteMovieUseCase,
	)
	return &usecase.DeleteMovieUseCase{}
}

func NewWebMovieHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebMovieHandler {
	wire.Build(
		setMovieRepositoryDependency,
		setMovieCreatedEvent,
		setMovieUpdatedEvent,
		setMovieDeletedEvent,
		web.NewWebMovieHandler,
	)
	return &web.WebMovieHandler{}
}
