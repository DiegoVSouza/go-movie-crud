package service

import (
	"context"

	dto "monte_clone_go/internal/DTO"
	"monte_clone_go/internal/infra/grpc/pb"
	"monte_clone_go/internal/usecase"
)

type MovieService struct {
	pb.UnimplementedMovieServiceServer
	createMovieUseCase usecase.CreateMovieUseCase
}

func NewMovieService(createMovieUseCase usecase.CreateMovieUseCase) *MovieService {
	return &MovieService{
		createMovieUseCase: createMovieUseCase,
	}
}

func (s *MovieService) CreateMovie(ctx context.Context, in *pb.CreateMovieRequest) (*pb.CreateMovieResponse, error) {
	dto := dto.MovieInputDTO{
		ID:       in.Id,
		Title:    in.Title,
		Director: in.Director,
		Writer:   in.Writer,
		Link:     in.Link,
		Duration: float64(in.Duration),
	}

	output, err := s.createMovieUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}

	return &pb.CreateMovieResponse{
		Id:       output.ID,
		Title:    output.Title,
		Director: output.Director,
		Writer:   output.Writer,
		Link:     output.Link,
		Duration: float32(output.Duration),
	}, nil
}
