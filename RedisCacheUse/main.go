package main

import (
    "fmt"
    "io"
    "net/http"
    "time"
    "context"

    "github.com/redis/go-redis/v9"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Request received:", r.URL.Path)

    // External API endpoint to fetch user data
    url := "https://jsonplaceholder.typicode.com/users"

    // Initialize Redis client with connection configuration
    // Connects to local Redis instance on port 6379
    client := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
        Protocol: 2,
    })

    // Create background context for Redis operations
    // Enables timeout management and operation cancellation
    ctx := context.Background()

    // ============================================
    // CACHE-ASIDE PATTERN: Check cache first
    // ============================================
    // Attempt to retrieve cached user data from Redis using key "users"
    // If data exists in cache, we avoid calling the external API
    resfromredis, err := client.Get(ctx, "users").Result()

    // If key doesn't exist in Redis, err will be redis.Nil
    if err != nil {
        fmt.Println("Cache MISS: Data not found in Redis, calling external API")
    }

    // Check if we have valid cached data (non-empty response)
    // If cache hit occurs, return cached response immediately
    if resfromredis != "" {
        fmt.Println("Cache HIT: Returning cached response (faster response time)")
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(resfromredis))
        return // Exit early to avoid redundant API call
    }

    // ============================================
    // CACHE MISS: Fetch from external API
    // ============================================
    // Call external JSONPlaceholder API to fetch user data
    res, err := http.Get(url)
    if err != nil {
        http.Error(w, "Failed to call external API", http.StatusInternalServerError)
        return
    }
    defer res.Body.Close()

    // Read response body from API call
    body, err := io.ReadAll(res.Body)
    if err != nil {
        http.Error(w, "Failed to read response", http.StatusInternalServerError)
        return
    }

    // ============================================
    // POPULATE CACHE: Store API response in Redis
    // ============================================
    // Cache the API response in Redis with key "users"
    // TTL set to 0 means no expiration (we'll set it separately)
    err = client.Set(ctx, "users", string(body), 0).Err()
    if err != nil {
        fmt.Println("Failed to cache data in Redis:", err)
    }

    // ============================================
    // SET CACHE EXPIRATION (TTL)
    // ============================================
    // Set Time-To-Live (TTL) for cached data to 30 seconds
    // After 30 seconds, Redis will automatically delete this key
    // This ensures fresh data is fetched periodically
    err = client.Expire(ctx, "users", time.Second*30).Err()
    if err != nil {
        fmt.Println("Failed to set expiration for cached data:", err)
    }

    // Return API response to client
    w.Header().Set("Content-Type", "application/json")
    w.Write(body)
}

func main() {
    // Register HTTP handler function for /users endpoint
    http.HandleFunc("/users", handler)

    fmt.Println("🚀 Server running on port 8083")
    fmt.Println("📌 Endpoint: http://localhost:8083/users")
    fmt.Println("⏱️  Cache TTL: 30 seconds")

    // Start HTTP server on port 8083
    // Blocks until server error occurs
    err := http.ListenAndServe(":8083", nil)
    if err != nil {
        panic(err)
    }
}