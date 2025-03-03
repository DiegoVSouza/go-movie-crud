package event

import "time"

type MovieDeleted struct {
	Name    string
	Payload interface{}
}

func NewMovieDeleted() *MovieDeleted {
	return &MovieDeleted{
		Name: "MovieDeleted",
	}
}

func (e *MovieDeleted) GetName() string {
	return e.Name
}

func (e *MovieDeleted) GetPayload() interface{} {
	return e.Payload
}

func (e *MovieDeleted) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *MovieDeleted) GetDateTime() time.Time {
	return time.Now()
}
