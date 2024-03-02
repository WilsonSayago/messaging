package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type EventHandler interface {
	Handle(msg *amqp.Delivery)
}
