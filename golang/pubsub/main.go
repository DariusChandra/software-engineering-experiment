package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

const (
	projectID      = "your-project-id"
	topicID        = "your-topic-id"
	subscriptionID = "your-subscription-id"
	emulatorHost   = "localhost:8085"
)

func main() {
	ctx := context.Background()

	// Create Pub/Sub client with emulator option
	client, err := pubsub.NewClient(ctx, projectID, option.WithEndpoint(fmt.Sprintf(emulatorHost)))
	if err != nil {
		log.Fatalf("Error creating Pub/Sub client: %v", err)
	}

	// Create a topic
	topic, err := client.CreateTopic(ctx, topicID)
	if err != nil {
		log.Fatalf("Error creating topic: %v", err)
	}

	fmt.Printf("Topic created: %v\n", topic)

	// Create a subscription
	sub, err := client.CreateSubscription(ctx, subscriptionID, pubsub.SubscriptionConfig{
		Topic: topic,
	})
	if err != nil {
		log.Fatalf("Error creating subscription: %v", err)
	}

	fmt.Printf("Subscription created: %v\n", sub)

	// Publish a message
	message := &pubsub.Message{
		Data: []byte("Hello, Pub/Sub!"),
	}
	result := topic.Publish(ctx, message)
	_, err = result.Get(ctx)
	if err != nil {
		log.Fatalf("Error publishing message: %v", err)
	}

	// Pull messages from the subscription
	cctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var messages []*pubsub.Message
	err = sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		messages = append(messages, msg)
		msg.Ack()
	})
	if err != nil {
		log.Fatalf("Error receiving messages: %v", err)
	}

	// Print received messages
	for _, msg := range messages {
		fmt.Printf("Received message: %s\n", string(msg.Data))
	}
}
