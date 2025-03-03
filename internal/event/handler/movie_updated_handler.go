package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"monte_clone_go/pkg/events"

	"github.com/streadway/amqp"
)

type MovieUpdatedHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewMovieUpdatedHandler(rabbitMQChannel *amqp.Channel) *MovieUpdatedHandler {
	return &MovieUpdatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *MovieUpdatedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Movie updated: %v\n", event.GetPayload())

	jsonOutput, err := json.Marshal(event.GetPayload())
	if err != nil {
		fmt.Printf("error marshalling updated event: %v\n", err)
		return
	}

	msg := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	if err := h.RabbitMQChannel.Publish(
		"amq.direct",
		"movie.updated",
		false,
		false,
		msg,
	); err != nil {
		fmt.Printf("error publishing updated event: %v\n", err)
	}
}
