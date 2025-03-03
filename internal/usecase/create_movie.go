package usecase

import (
	dto "monte_clone_go/internal/DTO"
	"monte_clone_go/internal/entity"
	"monte_clone_go/internal/event"
	"monte_clone_go/pkg/events"
)

type CreateMovieUseCase struct {
	MovieRepository entity.MovieRepositoryInterface
	MovieCreated    event.MovieCreatedEvent
	EventDispatcher events.EventDispatcherInterface
}

func NewCreateMovieUseCase(
	MovieRepository entity.MovieRepositoryInterface,
	MovieCreated event.MovieCreatedEvent,
	EventDispatcher events.EventDispatcherInterface,
) *CreateMovieUseCase {
	return &CreateMovieUseCase{
		MovieRepository: MovieRepository,
		MovieCreated:    MovieCreated,
		EventDispatcher: EventDispatcher,
	}
}

func (c *CreateMovieUseCase) Execute(input dto.MovieInputDTO) (dto.MovieOutputDTO, error) {
	movie, err := entity.NewMovie(input.ID, input.Title, input.Director, input.Writer, input.Link, input.Duration)
	if err != nil {
		return dto.MovieOutputDTO{}, err
	}
	movie.Title = input.Title

	if err := c.MovieRepository.Save(movie); err != nil {
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

	c.MovieCreated.SetPayload(output)
	c.EventDispatcher.Dispatch(c.MovieCreated)

	return output, nil
}
