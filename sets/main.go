package main

import (
    "context"
    "fmt"

    "github.com/redis/go-redis/v9"
)

func main() {
    // Initialize Redis client with connection parameters
    // Sets up connection to local Redis instance on default port 6379
    client := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
        Protocol: 2,
    })

    // Create a background context for all Redis operations
    // Context enables timeout management and cancellation across operations
    ctx := context.Background()

    // ============================================
    // REDIS SETS: Unordered collection of unique strings
    // Perfect for membership testing, unions, and intersections
    // ============================================

    // SAdd: Add multiple car brands to the "goset:cars" set
    // Note: Redis sets automatically eliminate duplicates
    // First operation adds: tata, enova, ford, tesla
    err := client.SAdd(ctx, "goset:cars", "tata", "enova", "ford", "tesla").Err()
    if err != nil {
        panic(err)
    }

    // Second SAdd operation on same key adds: tata (duplicate, ignored), mahindra, suzuki
    // Only mahindra and suzuki are actually added to the set
    err = client.SAdd(ctx, "goset:cars", "tata", "mahindra", "suzuki").Err()
    if err != nil {
        panic(err)
    }

    // SMembers: Retrieve all members from "goset:cars" set
    // Returns unordered collection of all unique values
    val, err := client.SMembers(ctx, "goset:cars").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("goset:cars members:", val)

    // ============================================
    // Create a separate set for Indian car brands
    // ============================================
    err = client.SAdd(ctx, "golist:indiancars", "mahindra", "tata", "ashok").Err()
    if err != nil {
        panic(err)
    }

    // Retrieve and display members of Indian cars set
    val1, err := client.SMembers(ctx, "golist:indiancars").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("golist:indiancars members:", val1)

    // ============================================
    // Set Operations: Add more brands and demonstrate removals
    // ============================================

    // SAdd: Add international brands to "goset:cars"
    err = client.SAdd(ctx, "goset:cars", "byd", "kia", "honda", "maruti").Err()
    if err != nil {
        panic(err)
    }

    // SRem: Remove specific members from the set
    // Removes "tesla" and "byd" from "goset:cars"
    err = client.SRem(ctx, "goset:cars", "tesla", "byd").Err()
    if err != nil {
        panic(err)
    }

    // ============================================
    // Membership Testing and Cardinality
    // ============================================

    // SIsMember: Check if "tata" exists in "goset:cars"
    // Returns boolean: true if member exists, false otherwise
    val2, err := client.SIsMember(ctx, "goset:cars", "tata").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("Is 'tata' member of goset:cars?", val2)

    // SCard: Get the cardinality (total count) of set members
    // Useful for determining set size in O(1) time complexity
    val3, err := client.SCard(ctx, "goset:cars").Result()
    if err != nil {
        panic(err)
    }

    // ============================================
    // Set Scanning (for large datasets)
    // ============================================

    // SScan: Incrementally iterate through set members
    // Parameters: key, cursor (start at 0), pattern ("" = all), count (0 = default)
    // Useful for iterating large sets without loading entire set into memory
    val4, cursor, err := client.SScan(ctx, "goset:cars", 0, "", 0).Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("goset:cars scan:", val4, "cursor:", cursor)
    fmt.Println("goset:cars cardinality:", val3)

    // ============================================
    // Create Bharat (India) cars set for set operations
    // ============================================
    err = client.SAdd(ctx, "goset:bharatcars", "mahindra", "tata", "ashok").Err()
    if err != nil {
        panic(err)
    }

    // ============================================
    // SET OPERATIONS: Union and Intersection
    // ============================================

    // SUnion: Get all unique members from multiple sets
    // Combines "golist:indiancars" and "goset:bharatcars"
    // Result contains all members from both sets (deduplicated)
    val6, err := client.SUnion(ctx, "golist:indiancars", "goset:bharatcars").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("Union of golist:indiancars and goset:bharatcars:", val6)

    // SInter: Get common members across sets
    // Finds members that exist in BOTH "goset:cars" AND "goset:bharatcars"
    // Common brands: tata, mahindra, ashok (if they exist in cars set)
    val5, err := client.SInter(ctx, "goset:cars", "goset:bharatcars").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("Intersection of goset:cars and goset:bharatcars:", val5)

    // Display Redis client instance info
    fmt.Println(client)
}