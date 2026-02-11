package main

import (
    "context"
    "fmt"
    "time"

    "github.com/redis/go-redis/v9"
)

func main() {
    ctx := context.Background()

    // Initialize Redis client for publishing
    client := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })

    // ============================================
    // PUBLISH: Send messages to subscribers
    // ============================================

    // Publish messages to "GoMessage" channel
    // Returns number of subscribers that received the message
    for i := 1; i <= 5; i++ {
        message := fmt.Sprintf("Hello from Publisher #%d", i)
        
        // Publish message to channel
        // Parameters: context, channel name, message payload
        numSubscribers, err := client.Publish(ctx, "GoMessage", message).Result()
        if err != nil {
            panic(err)
        }

        fmt.Printf("📤 Published message to %d subscribers: %s\n", numSubscribers, message)
        time.Sleep(1 * time.Second)
    }

    fmt.Println("✅ All messages published!")
}