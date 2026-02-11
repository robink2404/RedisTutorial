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
    // REDIS SORTED SETS: Ordered collection with scores
    // Members are ordered by score (numeric value)
    // Perfect for leaderboards, rankings, and priority queues
    // ============================================

    // ZAdd: Add members with scores to sorted set
    // Format: redis.Z{score, member}
    // Members are automatically ordered by score (ascending)
    err := client.ZAdd(ctx, "goSortedSet:list1",
        redis.Z{5, "kareena"},
        redis.Z{2, "priyanka"},
        redis.Z{3, "deepika"},
    ).Err()
    if err != nil {
        panic(err)
    }

    // Add another member to the sorted set
    // Score 10 places "alia" at the highest position
    err = client.ZAdd(ctx, "goSortedSet:list1", redis.Z{10, "alia"}).Err()
    if err != nil {
        panic(err)
    }

    // ============================================
    // ZRANGE: Retrieve members in score order (ascending)
    // Parameters: key, start index, stop index
    // 0 = first element, -1 = last element
    // ============================================

    // ZRange: Get all members in ascending order by score
    // Range: 0 to -1 (all members)
    val, err := client.ZRange(ctx, "goSortedSet:list1", 0, -1).Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("goSortedSet:list1 members (ascending):", val)
    // Output: [priyanka deepika kareena alia]

    // ============================================
    // ZREVRANGE: Retrieve members in reverse score order (descending)
    // Returns highest-scored members first
    // ============================================

    // ZRevRange: Get all members in descending order by score
    val, err = client.ZRevRange(ctx, "goSortedSet:list1", 0, -1).Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("goSortedSet:list1 members (descending):", val)
    // Output: [alia kareena deepika priyanka]

    // ============================================
    // ZREM: Remove member from sorted set
    // Useful for removing outdated or invalidated entries
    // ============================================

    // ZRem: Remove "priyanka" from sorted set
    err = client.ZRem(ctx, "goSortedSet:list1", "priyanka").Err()
    if err != nil {
        panic(err)
    }

    // Verify removal by retrieving all members
    val, err = client.ZRange(ctx, "goSortedSet:list1", 0, -1).Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("After removal:", val)

    // ============================================
    // ZRANK: Get position/rank of member
    // Returns 0-based index (0 = lowest score)
    // ZRevRank: Returns reverse position (0 = highest score)
    // ============================================

    // ZRank: Get rank of "alia" in ascending order
    // rank = 2 means alia is at index 2 (3rd position from bottom)
    rank, err := client.ZRank(ctx, "goSortedSet:list1", "alia").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("Rank of alia (ascending):", rank) // Output: 2

    // ============================================
    // ZCARD: Get total number of members in sorted set
    // ============================================

    // ZCard: Get count of all members
    count, err := client.ZCard(ctx, "goSortedSet:list1").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("Total members in sorted set:", count)

    // ============================================
    // ZSCORE: Get score of a member
    // ============================================

    // ZScore: Get score of "deepika"
    score, err := client.ZScore(ctx, "goSortedSet:list1", "deepika").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("Score of deepika:", score) // Output: 3

    // Display Redis client instance info
    fmt.Println("\n✅ Redis client connected:", client)
}