package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"monte_clone_go/pkg/events"

	"github.com/streadway/amqp"
)

type MovieDeletedHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewMovieDeletedHandler(rabbitMQChannel *amqp.Channel) *MovieDeletedHandler {
	return &MovieDeletedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *MovieDeletedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Movie deleted: %v\n", event.GetPayload())

	jsonOutput, err := json.Marshal(event.GetPayload())
	if err != nil {
		fmt.Printf("error marshalling deleted event: %v\n", err)
		return
	}

	msg := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	if err := h.RabbitMQChannel.Publish(
		"amq.direct",
		"movie.deleted",
		false,
		false,
		msg,
	); err != nil {
		fmt.Printf("error publishing deleted event: %v\n", err)
	}
}
