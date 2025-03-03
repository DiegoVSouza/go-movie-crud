package web

import (
	"encoding/json"
	"net/http"

	dto "monte_clone_go/internal/DTO"
	"monte_clone_go/internal/entity"
	"monte_clone_go/internal/event"
	"monte_clone_go/internal/usecase"
	"monte_clone_go/pkg/events"
)

type WebMovieHandler struct {
	EventDispatcher   events.EventDispatcherInterface
	MovieRepository   entity.MovieRepositoryInterface
	MovieCreatedEvent event.MovieCreatedEvent
	MovieUpdatedEvent event.MovieUpdatedEvent
	MovieDeletedEvent event.MovieDeletedEvent
}

func NewWebMovieHandler(
	eventDispatcher events.EventDispatcherInterface,
	movieRepository entity.MovieRepositoryInterface,
	movieCreatedEvent event.MovieCreatedEvent,
	movieUpdatedEvent event.MovieUpdatedEvent,
	movieDeletedEvent event.MovieDeletedEvent,
) *WebMovieHandler {
	return &WebMovieHandler{
		EventDispatcher:   eventDispatcher,
		MovieRepository:   movieRepository,
		MovieCreatedEvent: movieCreatedEvent,
		MovieUpdatedEvent: movieUpdatedEvent,
		MovieDeletedEvent: movieDeletedEvent,
	}
}

func (h *WebMovieHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input dto.MovieInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createMovie := usecase.NewCreateMovieUseCase(h.MovieRepository, h.MovieCreatedEvent, h.EventDispatcher)
	output, err := createMovie.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *WebMovieHandler) Get(w http.ResponseWriter, r *http.Request) {
	// Extrair filtros da query string
	filters := make(map[string]string)
	for key, values := range r.URL.Query() {
		if len(values) > 0 {
			filters[key] = values[0]
		}
	}

	getMovie := usecase.NewGetMovieUseCase(h.MovieRepository)
	outputs, err := getMovie.Execute(filters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(outputs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *WebMovieHandler) Update(w http.ResponseWriter, r *http.Request) {
	var input dto.MovieInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updateMovie := usecase.NewUpdateMovieUseCase(h.MovieRepository, h.MovieUpdatedEvent, h.EventDispatcher)
	output, err := updateMovie.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *WebMovieHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// Extrair o ID da URL
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	deleteMovie := usecase.NewDeleteMovieUseCase(h.MovieRepository, h.MovieDeletedEvent, h.EventDispatcher)
	err := deleteMovie.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
