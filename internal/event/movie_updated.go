package event

import "time"

type MovieUpdated struct {
	Name    string
	Payload interface{}
}

func NewMovieUpdated() *MovieUpdated {
	return &MovieUpdated{
		Name: "MovieUpdated",
	}
}

func (e *MovieUpdated) GetName() string {
	return e.Name
}

func (e *MovieUpdated) GetPayload() interface{} {
	return e.Payload
}

func (e *MovieUpdated) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *MovieUpdated) GetDateTime() time.Time {
	return time.Now()
}
