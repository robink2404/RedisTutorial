package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Subscribe to channel
	pubsub := client.Subscribe(ctx, "GoMessage")

	// Ensure subscription is active
	_, err := pubsub.Receive(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("Subscribed to channel: GoMessage")

	// Channel to receive messages
	ch := pubsub.Channel()

	// Listen for messages
	for msg := range ch {
		fmt.Printf("📩 Received message from %s: %s\n", msg.Channel, msg.Payload)
	}
}
