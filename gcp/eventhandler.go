package gcp

import (
	"cloud.google.com/go/pubsub"
)

type EventHandler interface {
	Handle(msg *pubsub.Message)
}
