package gcp

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"log"
)

type ClientWrapper struct {
	client     *pubsub.Client
	subscriber *pubsub.Subscription
	topic      *pubsub.Topic
}

func InitSubscriber(sub string, projectId string, numRoutines int, maxMessages int, event EventHandler) {
	//os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8284")
	ctx := context.Background()
	client, err := NewClient(ctx, projectId)

	if err != nil {
		log.Fatalf("Failed to create client: %v\n", err)
	}

	client.Subscriber(sub, projectId, numRoutines, maxMessages)

	log.Printf("Start receiving messages, sub=%v NumGoRutine=%v, MaxMessages=%v", sub, numRoutines, maxMessages)

	client.Receive(ctx, event)

	if err != nil {
		log.Fatalf("Failed to create client: %v\n", err)
	}
}

func NewClient(ctx context.Context, projectId string) (*ClientWrapper, error) {
	client, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		log.Printf("Failed to create client: %v\n", err)
		return nil, err
	}
	return &ClientWrapper{client: client}, nil
}

func (cw *ClientWrapper) CloseClient() error {
	err := cw.client.Close()
	if err != nil {
		log.Printf("Failed to close client: %v\n", err)
	}
	return err
}

func (cw *ClientWrapper) Subscriber(name string, projectId string, numRoutines int, maxMessages int) {
	sub := cw.client.SubscriptionInProject(name, projectId)
	sub.ReceiveSettings.NumGoroutines = numRoutines
	sub.ReceiveSettings.Synchronous = false
	sub.ReceiveSettings.MaxOutstandingMessages = maxMessages
	cw.subscriber = sub
}

func (cw *ClientWrapper) Receive(ctx context.Context, event EventHandler) {
	cw.subscriber.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		event.Handle(msg)
	})
}

func (cw *ClientWrapper) ReceiveAndHandle(ctx context.Context, receivingFunc func(ctx context.Context, msg *pubsub.Message)) error {
	return cw.subscriber.Receive(ctx, receivingFunc)
}

func (cw *ClientWrapper) Topic(name string, projectId string) {
	cw.topic = cw.client.TopicInProject(name, projectId)
}

func (cw *ClientWrapper) PublishAndStop() {
	defer cw.topic.Stop()
	return
}

func (cw *ClientWrapper) Publish(payload []byte, attr map[string]string, ctx context.Context) error {
	msg := &pubsub.Message{
		Data:       payload,
		Attributes: attr,
	}
	result := cw.topic.Publish(ctx, msg)
	_, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("Failed to publish message: %v\n", err)
	}
	return nil
}
