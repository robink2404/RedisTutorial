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
    // REDIS LISTS: Ordered collection of strings
    // Perfect for queues, stacks, and activity feeds
    // ============================================

    // RPush: Push elements to the right (tail) of the list
    // Parameters: context, key, values...
    // O(1) operation - constant time regardless of list size
    err := client.RPush(ctx, "golist:queue", 2).Err()
    if err != nil {
        panic(err)
    }

    // RPush: Add multiple elements at once
    // More efficient than multiple individual pushes
    err1 := client.RPush(ctx, "golist:queue", 1, 3, 4, 5, 6, 7, 8).Err()
    if err1 != nil {
        panic(err1)
    }

    // ============================================
    // LIST OPERATIONS: Pop, Trim, and Move
    // ============================================

    // LPop: Remove and return element from left (head) of list
    // Perfect for FIFO (First-In-First-Out) queue operations
    err2 := client.LPop(ctx, "golist:queue").Err()
    if err2 != nil {
        panic(err2)
    }

    // LTrim: Keep only elements within specified range
    // Removes elements outside the range [start, stop]
    // Useful for limiting list size and removing old entries
    errmove := client.LTrim(ctx, "golist:queue", 0, 2).Err()
    if errmove != nil {
        panic(errmove)
    }
    fmt.Println("✅ List trimmed to first 3 elements")

    // LMove: Atomically move element from one list to another
    // Source: "golist:queue" (left/head), Destination: "golist:stack" (left/head)
    // Useful for transferring items between queues
    errmove = client.LMove(ctx, "golist:queue", "golist:stack", "left", "left").Err()
    if errmove != nil {
        panic(errmove)
    }
    fmt.Println("✅ Element moved from queue to stack")

    // BLPop: Blocking left pop - wait for element with timeout
    // Blocks up to 10 seconds waiting for element to be available
    // Returns immediately if element exists
    // Perfect for task queues and worker patterns
    errblockpop := client.BLPop(ctx, 10, "golist:queue").Err()
    if errblockpop != nil {
        panic(errblockpop)
    }
    fmt.Println("✅ Blocked pop operation complete")

    // ============================================
    // LIST INSPECTION: Length and Range queries
    // ============================================

    // LLen: Get the total length (count) of elements in list
    // O(1) operation - very fast regardless of list size
    val3, err3 := client.LLen(ctx, "golist:queue").Result()
    if err3 != nil {
        panic(err3)
    }
    fmt.Println("golist:queue length:", val3)

    // LRange: Retrieve elements within index range
    // Parameters: key, start index, stop index (-1 = last element)
    // Returns all elements from start to stop (inclusive)
    val1, err1 := client.LRange(ctx, "golist:queue", 0, val3-1).Result()
    if err1 != nil {
        panic(err1)
    }
    fmt.Println("golist:queue elements:", val1)

    // Get length of stack list
    val, err := client.LLen(ctx, "golist:stack").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("golist:stack length:", val)

    // Display Redis client instance info
    fmt.Println("\n✅ Redis client connected:", client)
}