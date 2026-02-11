package main

import (
    "context"
    "fmt"
    "time"

    "github.com/redis/go-redis/v9"
)

func main() {
    // Initialize Redis client with connection parameters
    // Connects to local Redis instance on port 6379
    client := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",     // No password set
        DB:       0,      // Use default DB
        Protocol: 2,      // Connection protocol
    })

    // Create a background context for all Redis operations
    // Enables timeout management and operation cancellation
    ctx := context.Background()

    // ============================================
    // REDIS STRINGS: Key-Value pairs
    // Simplest Redis data type, perfect for storing text/numbers
    // ============================================

    // Set: Store single key-value pairs
    // Parameters: context, key, value, expiration (0 = no expiration)
    err := client.Set(ctx, "cars:1", "Toyota", 0).Err()
    if err != nil {
        panic(err)
    }

    err1 := client.Set(ctx, "cars:2", "Honda", 0).Err()
    if err1 != nil {
        panic(err1)
    }

    // Get: Retrieve value by key
    // Returns the string value associated with the key
    val, err := client.Get(ctx, "cars:1").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("cars:1:", val)

    // ============================================
    // MULTI SET/GET: Batch operations
    // More efficient than individual Set/Get calls
    // ============================================

    // MSet: Set multiple key-value pairs atomically
    // Parameters: context, key1, value1, key2, value2, ...
    err2 := client.MSet(ctx, "cars:3", "ford", "cars:4", "kia", "cars:5", "suzuki").Err()
    if err2 != nil {
        panic(err2)
    }

    // MGet: Retrieve multiple values by keys
    // Returns values in same order as requested keys
    vals, err := client.MGet(ctx, "cars:3", "cars:4", "cars:5").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("Multiple cars:", vals)

    // ============================================
    // CONDITIONAL SET: MSetNX (Set if Not eXists)
    // Only sets keys if ALL of them don't exist
    // Useful for preventing key overwrites
    // ============================================

    // MSetNX: Set multiple keys only if NONE of them exist
    // Returns true if all keys were set, false if any key already exists
    ok, err := client.MSetNX(ctx,
        "cars:10", "bmw",
        "cars:11", "nissan",
    ).Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("Were keys set?", ok)

    // Retrieve the newly set keys
    vals2, err := client.MGet(ctx, "cars:11", "cars:10").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("Retrieved cars:", vals2)

    // ============================================
    // ATOMIC INCREMENT: Increment counter values
    // Performs increment operation atomically on server
    // ============================================

    // IncrBy: Increment a numeric value by specified amount
    // Useful for counters, page views, likes, etc.
    err3 := client.IncrBy(ctx, "counter", 10).Err()
    if err3 != nil {
        panic(err3)
    }

    // Get the incremented counter value
    valnew, errnew := client.Get(ctx, "counter").Result()
    if errnew != nil {
        panic(errnew)
    }
    fmt.Println("Counter value:", valnew)

    // ============================================
    // SCAN & DELETE: Efficiently find and delete keys
    // Scan is better than KEYS for production use
    // ============================================

    // MSet: Create sample SIM keys for deletion
    errlatest := client.MSet(ctx, "sims:1", "Airtel", "sims:2", "Jio").Err()
    if errlatest != nil {
        panic(errlatest)
    }

    // Scan: Iterate through keys matching pattern "sims:*"
    // Cursor-based iteration - more efficient for large keysets
    iter := client.Scan(ctx, 0, "sims:*", 0).Iterator()
    for iter.Next(ctx) {
        // Delete each matching key
        err = client.Del(ctx, iter.Val()).Err()
        if err != nil {
            panic(err)
        }
    }
    if err = iter.Err(); err != nil {
        panic(err)
    }
    fmt.Println("Deleted all sims:* keys")

    // ============================================
    // EXPIRATION: Time-To-Live (TTL) management
    // Automatically delete keys after specified duration
    // ============================================

    // Set with expiration: Store value with 10-second TTL
    // After 10 seconds, Redis automatically deletes this key
    errEXP := client.Set(ctx, "tempkey", "tempvalue", 10*time.Second).Err()
    if errEXP != nil {
        panic(errEXP)
    }

    // Expire: Set expiration on existing key
    // Parameters: context, key, duration
    expcars := client.Expire(ctx, "cars:10", 30*time.Second).Err()
    if expcars != nil {
        panic(expcars)
    }

    // TTL: Get remaining time to live for a key
    // Returns duration in seconds, -1 if no expiration, -2 if key doesn't exist
    valcar, err := client.TTL(ctx, "cars:10").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("TTL of cars:10:", valcar)

    // Display Redis client instance info
    fmt.Println("\n✅ Redis client connected:", client)
}