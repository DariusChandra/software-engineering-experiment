package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"google.golang.org/api/option"
	"log"
)

const (
	projectID    = "dh-adtech"
	topicID      = "adtech-adcredits-credit-events-stg-euw3-v1"
	emulatorHost = "localhost:8085"
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
}
