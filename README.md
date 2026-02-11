# Redis Cache Use Case - Go Tutorial

A practical tutorial demonstrating how to implement caching using Redis in Go. This project shows how to cache external API responses to improve performance and reduce API calls.

## Project Overview

This tutorial demonstrates:
- Connecting to Redis from Go
- Caching API responses in Redis
- Setting expiration times for cached data
- Implementing a simple cache-aside pattern

## Prerequisites

- **Go** (version 1.16 or higher)
- **Redis** (running locally on port 6379)
- **Basic knowledge** of Go and HTTP

## Setup Instructions

### 1. Install Go

If you don't have Go installed, download it from [golang.org](https://golang.org/dl/)

Verify installation:
```bash
go version
```

### 2. Install Redis

On macOS (using Homebrew):
```bash
brew install redis
```

Start Redis server:
```bash
redis-server
```

Or run in the background:
```bash
brew services start redis
```

### 3. Clone/Navigate to Project

```bash
cd /Users/robin/Desktop/redisGo/RedisCacheUse
```

### 4. Initialize Go Module (if not done)

```bash
go mod init rediscache
```

### 5. Install Dependencies

```bash
go get github.com/redis/go-redis/v9
```

### 6. Run the Application

```bash
go run main.go
```

You should see:
```
🚀 Server running on port 8083
```

## Testing the Application

### First Request (Cache Miss)
```bash
curl http://localhost:8083/users
```
Output in terminal:
```
Request received: /users
Data not found in Redis, calling external API
```
The response is fetched from the external API and cached in Redis.

### Second Request (Cache Hit)
```bash
curl http://localhost:8083/users
```
Output in terminal:
```
Request received: /users
Data found in Redis, returning cached response
```
The response is served directly from Redis cache (faster!).

### After 30 Seconds
Wait 30 seconds and make another request - the cache will expire and fetch fresh data from the API.

## How It Works

### Cache-Aside Pattern

```
Request → Check Redis → Found? 
  ├─ Yes → Return cached data
  └─ No  → Call API → Cache result → Return data
```

### Code Flow

1. **Request arrives** at `/users` endpoint
2. **Check Redis** for cached data with key `"users"`
3. **If found:** Return cached response immediately
4. **If not found:** 
   - Call external JSONPlaceholder API
   - Cache the response in Redis
   - Set 30-second expiration
   - Return response to client

## Key Redis Operations Used

| Operation | Purpose |
|-----------|---------|
| `client.Get()` | Retrieve data from cache |
| `client.Set()` | Store data in cache |
| `client.Expire()` | Set expiration time (TTL) |

## Configuration

### Redis Connection
```go
client := redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",  // Redis server address
    Password: "",                 // No password by default
    DB:       0,                  // Database 0
    Protocol: 2,                  // Protocol version
})
```

### Cache Expiration
Currently set to **30 seconds**. Modify in `main.go`:
```go
client.Expire(ctx, "users", time.Second*30)  // Change 30 to desired seconds
```

## Performance Benefits

- **First request:** ~200-500ms (API call)
- **Subsequent requests (within 30s):** ~5-10ms (cached)
- **Reduction:** 95%+ faster for cached requests

## Troubleshooting

### Error: `connection refused`
- Ensure Redis is running: `redis-server`
- Check port 6379 is available

### Error: `cannot find package`
- Run: `go mod tidy`
- Re-run: `go run main.go`

### Cache not working
- Verify Redis is accessible: `redis-cli ping` (should return `PONG`)
- Check cache expiration hasn't passed

## Next Steps

Explore these advanced topics:
- Implement different cache invalidation strategies
- Add cache warming
- Use Redis Lua scripts for atomic operations
- Implement distributed caching
- Add cache statistics and monitoring

## Project Structure

```
RedisCacheUse/
├── main.go              # Application code
├── go.mod              # Module definition
└── go.sum              # Dependency checksums
```

## Resources

- [Go Redis Documentation](https://github.com/redis/go-redis)
- [Redis Commands](https://redis.io/commands)
- [Cache-Aside Pattern](https://docs.microsoft.com/en-us/azure/architecture/patterns/cache-aside)

## License

This is a tutorial project for learning purposes.
