package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"monte_clone_go/pkg/events"

	"github.com/streadway/amqp"
)

type MovieCreatedHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewMovieCreatedHandler(rabbitMQChannel *amqp.Channel) *MovieCreatedHandler {
	return &MovieCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *MovieCreatedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Movie created: %v", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.Publish(
		"amq.direct",
		"",
		false,
		false,
		msgRabbitmq,
	)
}
