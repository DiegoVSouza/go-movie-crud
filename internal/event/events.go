package event

import "monte_clone_go/pkg/events"

// MovieCreatedEvent é um tipo específico para o evento de criação de filme
type MovieCreatedEvent events.EventInterface

// MovieUpdatedEvent é um tipo específico para o evento de atualização de filme
type MovieUpdatedEvent events.EventInterface

// MovieDeletedEvent é um tipo específico para o evento de exclusão de filme
type MovieDeletedEvent events.EventInterface
