package usecase

import (
	dto "monte_clone_go/internal/DTO"
	"monte_clone_go/internal/entity"
)

type GetMovieUseCase struct {
	MovieRepository entity.MovieRepositoryInterface
}

func NewGetMovieUseCase(MovieRepository entity.MovieRepositoryInterface) *GetMovieUseCase {
	return &GetMovieUseCase{MovieRepository: MovieRepository}
}

func (g *GetMovieUseCase) Execute(filters map[string]string) ([]dto.MovieOutputDTO, error) {
	movies, err := g.MovieRepository.Get(filters)
	if err != nil {
		return nil, err
	}

	var outputs []dto.MovieOutputDTO
	for _, movie := range movies {
		outputs = append(outputs, dto.MovieOutputDTO{
			ID:       movie.ID,
			Title:    movie.Title,
			Director: movie.Director,
			Writer:   movie.Writer,
			Link:     movie.Link,
			Duration: movie.Duration,
		})
	}

	return outputs, nil
}
