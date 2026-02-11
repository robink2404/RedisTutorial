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
    // REDIS HASHES: Key-Field-Value mappings
    // Perfect for storing structured objects like user profiles, products
    // More efficient than storing entire JSON strings
    // ============================================

    // HSet: Set multiple field-value pairs in hash at once
    // Parameters: context, key, field1, value1, field2, value2, ...
    // Atomically sets all fields - either all succeed or all fail
    err := client.HSet(ctx, "gohash:bike1",
        "model", "yamaha",
        "brand", "yamaha",
        "price", 200000,
    ).Err()
    if err != nil {
        panic(err)
    }

    // HSet: Add additional field to existing hash
    // yearOfPurchase field is added to the same hash
    err = client.HSet(ctx, "gohash:bike1", "yearOfPurchase", 2023).Err()
    if err != nil {
        panic(err)
    }

    // ============================================
    // HASH RETRIEVAL: Single and multiple fields
    // ============================================

    // HGet: Retrieve single field value from hash
    // Returns the value associated with the specified field
    // Returns error if field doesn't exist
    val, err := client.HGet(ctx, "gohash:bike1", "price").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("gohash:bike1 price:", val)

    // HMGet: Retrieve multiple field values at once
    // Parameters: context, key, field1, field2, ...
    // Returns values in same order as requested fields
    // More efficient than multiple HGet calls
    val1, err := client.HMGet(ctx, "gohash:bike1", "model", "brand").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("gohash:bike1 model and brand:", val1)

    // HGetAll: Retrieve all fields and values from hash
    // Returns a map[string]string with all field-value pairs
    // Useful for getting complete object data
    val2, err := client.HGetAll(ctx, "gohash:bike1").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("gohash:bike1 all fields:", val2)

    // ============================================
    // HASH ATOMIC OPERATIONS: Increment field values
    // ============================================

    // HIncrBy: Increment numeric field by specified amount
    // Atomically increases the price field by 2000
    // Perfect for counters, inventory, and balance updates
    err = client.HIncrBy(ctx, "gohash:bike1", "price", 2000).Err()
    if err != nil {
        panic(err)
    }

    // HGet: Verify the incremented value
    // Retrieves the updated price after increment
    val3, err := client.HGet(ctx, "gohash:bike1", "price").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("gohash:bike1 price after increment:", val3)

    // Display Redis client instance info
    fmt.Println("\n✅ Redis client connected:", client)
}