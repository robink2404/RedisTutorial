package main

import (
    "context"
    "fmt"

    "github.com/redis/go-redis/v9"
)

func main() {
    // Initialize Redis client with connection parameters
    // Connects to local Redis instance on port 6379
    client := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
        Protocol: 2,
    })

    // Create a background context for all Redis operations
    // Enables timeout management and operation cancellation
    ctx := context.Background()

    // ============================================
    // REDIS STREAMS: Append-only log data structure
    // Perfect for time-series data, event logs, and real-time feeds
    // ============================================

    // XAdd: Add entries to stream
    // Stream ID is auto-generated with "*" (timestamp-sequence format)
    // Ideal for sensor data, event logs, message queues
    res1, err := client.XAdd(ctx, &redis.XAddArgs{
        Stream: "gostream:sensorData",
        ID:     "*", // Auto-generate ID based on timestamp
        Values: map[string]interface{}{
            "temperature": 35,
            "humidity":    80,
        },
    }).Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("Added sensor entry ID:", res1)

    // Add another sensor reading to the stream
    res2, err := client.XAdd(ctx, &redis.XAddArgs{
        Stream: "gostream:sensorData",
        ID:     "*",
        Values: map[string]interface{}{
            "temperature": 30,
            "humidity":    65,
        },
    }).Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("Added sensor entry ID:", res2)

    // ============================================
    // XREAD: Read entries from stream
    // Parameters: Stream name, start ID (0 = from beginning), Count, Block timeout
    // Block: 0 = non-blocking, >0 = wait for new entries (milliseconds)
    // ============================================

    // XRead: Read last 2 entries from stream
    // Count: 2 means read maximum 2 entries
    // Block: 5000 means wait up to 5 seconds for new data
    res3, err := client.XRead(ctx, &redis.XReadArgs{
        Streams: []string{"gostream:sensorData", "0"}, // "0" = read from beginning
        Count:   2,  // Max entries to return
        Block:   5000, // Wait 5 seconds for new data
    }).Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("Stream entries read:", res3)

    // ============================================
    // XLEN: Get total number of entries in stream
    // Useful for monitoring stream size
    // ============================================

    // XLen: Get stream length
    // Returns total count of all entries in the stream
    res4, err := client.XLen(ctx, "gostream:sensorData").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("Stream total entries:", res4)

    // ============================================
    // XRANGE: Retrieve entries within ID range
    // Parameters: stream, start ID, end ID
    // "-" = earliest entry, "+" = latest entry
    // ============================================

    // XRange: Get entries within ID range
    // Useful for retrieving historical data by timestamp range
    res5, err := client.XRange(ctx, "gostream:sensorData", "-", "+").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("Stream range entries:", res5)

    // ============================================
    // XRANGENX: Get N entries starting from ID
    // More efficient for large streams
    // ============================================

    // XRangeN: Get max 2 entries from stream (alternative to XRange)
    res6, err := client.XRangeN(ctx, "gostream:sensorData", "-", "+", 2).Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("Stream first 2 entries:", res6)

    // Display Redis client instance info
    fmt.Println("\n✅ Redis client connected:", client)
}