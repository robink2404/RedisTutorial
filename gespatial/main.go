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
    // REDIS GEO: Geographic spatial indexing
    // Perfect for location-based services and proximity searches
    // Uses geohashing algorithm for efficient spatial queries
    // ============================================

    // GeoAdd: Add location data with coordinates to geospatial index
    // Parameters: context, key, GeoLocation{Longitude, Latitude, Name}
    // Longitude range: -180 to 180, Latitude range: -90 to 90
    // Stores location as a sorted set with geohash score
    err := client.GeoAdd(ctx, "goGeo:myLocation", &redis.GeoLocation{
        Longitude: 2.3522,  // Paris, France
        Latitude:  48.8566,
        Name:      "france",
    }).Err()
    if err != nil {
        panic(err)
    }
    fmt.Println("✅ Added Paris location")

    // GeoAdd: Add another location to the same geospatial index
    // Multiple locations can be stored under the same key
    err = client.GeoAdd(ctx, "goGeo:myLocation", &redis.GeoLocation{
        Longitude: 10.4515,  // Berlin, Germany
        Latitude:  52.5200,
        Name:      "germany",
    }).Err()
    if err != nil {
        panic(err)
    }
    fmt.Println("✅ Added Berlin location")

    // ============================================
    // GEOSEARCH: Find locations within radius
    // Proximity-based queries for location services
    // ============================================

    // GeoSearch: Query locations within specified radius
    // Parameters:
    //   - Context
    //   - Key (geospatial index)
    //   - GeoSearchQuery with center point and radius
    //   - Radius: search distance
    //   - RadiusUnit: "km", "m", "mi", "ft"
    // Returns all locations within the specified radius from center point
    res, err := client.GeoSearch(ctx, "goGeo:myLocation", &redis.GeoSearchQuery{
        Longitude:  10.4515,   // Center point: Berlin coordinates
        Latitude:   52.5210,
        Radius:     5000,      // Search radius: 5000 km
        RadiusUnit: "km",      // Units in kilometers
    }).Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("📍 Locations within 5000 km of Berlin:", res)

    // Display Redis client instance info
    fmt.Println("\n✅ Redis client connected:", client)
}