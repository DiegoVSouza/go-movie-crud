package usecase

import (
	"monte_clone_go/internal/entity"
	"monte_clone_go/internal/event"
	"monte_clone_go/pkg/events"
)

type DeleteMovieUseCase struct {
	MovieRepository entity.MovieRepositoryInterface
	MovieDeleted    event.MovieDeletedEvent
	EventDispatcher events.EventDispatcherInterface
}

func NewDeleteMovieUseCase(
	MovieRepository entity.MovieRepositoryInterface,
	MovieDeleted event.MovieDeletedEvent,
	EventDispatcher events.EventDispatcherInterface,
) *DeleteMovieUseCase {
	return &DeleteMovieUseCase{
		MovieRepository: MovieRepository,
		MovieDeleted:    MovieDeleted,
		EventDispatcher: EventDispatcher,
	}
}

func (d *DeleteMovieUseCase) Execute(id string) error {
	err := d.MovieRepository.Delete(id)
	if err != nil {
		return err
	}

	d.MovieDeleted.SetPayload(map[string]string{"id": id})
	d.EventDispatcher.Dispatch(d.MovieDeleted)

	return nil
}
