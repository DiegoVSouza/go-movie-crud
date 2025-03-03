package event

import "time"

type MovieCreated struct {
	Name    string
	Payload interface{}
}

func NewMovieCreated() *MovieCreated {
	return &MovieCreated{
		Name: "MovieCreated",
	}
}

func (e *MovieCreated) GetName() string {
	return e.Name
}

func (e *MovieCreated) GetPayload() interface{} {
	return e.Payload
}

func (e *MovieCreated) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *MovieCreated) GetDateTime() time.Time {
	return time.Now()
}
