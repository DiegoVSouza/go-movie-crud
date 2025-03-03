package usecase

import (
	dto "monte_clone_go/internal/DTO"
	"monte_clone_go/internal/entity"
	"monte_clone_go/internal/event"
	"monte_clone_go/pkg/events"
)

type UpdateMovieUseCase struct {
	MovieRepository entity.MovieRepositoryInterface
	MovieUpdated    event.MovieUpdatedEvent
	EventDispatcher events.EventDispatcherInterface
}

func NewUpdateMovieUseCase(
	MovieRepository entity.MovieRepositoryInterface,
	MovieUpdated event.MovieUpdatedEvent,
	EventDispatcher events.EventDispatcherInterface,
) *UpdateMovieUseCase {
	return &UpdateMovieUseCase{
		MovieRepository: MovieRepository,
		MovieUpdated:    MovieUpdated,
		EventDispatcher: EventDispatcher,
	}
}

func (u *UpdateMovieUseCase) Execute(input dto.MovieInputDTO) (dto.MovieOutputDTO, error) {
	movie, err := entity.NewMovie(input.ID, input.Title, input.Director, input.Writer, input.Link, input.Duration)
	if err != nil {
		return dto.MovieOutputDTO{}, err
	}

	if err := u.MovieRepository.Update(movie); err != nil {
		return dto.MovieOutputDTO{}, err
	}

	output := dto.MovieOutputDTO{
		ID:       movie.ID,
		Title:    movie.Title,
		Director: movie.Director,
		Writer:   movie.Writer,
		Link:     movie.Link,
		Duration: movie.Duration,
	}

	u.MovieUpdated.SetPayload(output)
	u.EventDispatcher.Dispatch(u.MovieUpdated)

	return output, nil
}
