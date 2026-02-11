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

	url := "https://jsonplaceholder.typicode.com/users"

	client:=redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
		Protocol: 2,
	})
	ctx := context.Background()
	resfromredis,err:=client.Get(ctx,"users").Result() //if found in redis cache then ret

	if err!=nil{
		fmt.Println("Data not found in Redis, calling external API")
	}
	if resfromredis!= "" {
		fmt.Println("Data found in Redis, returning cached response")
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(resfromredis))
		return
	}

	res, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to call external API", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}
	err=client.Set(ctx,"users",string(body),0).Err() //cache the response in redis with key "users"
	if err!=nil{
		fmt.Println("Failed to cache data in Redis:", err)
	}

	err=client.Expire(ctx,"users",time.Second*30).Err() //set expiration for the cached data to 30 seconds
	if err!=nil{
		fmt.Println("Failed to set expiration for cached data:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}


func main() {

	http.HandleFunc("/users", handler)

	fmt.Println("🚀 Server running on port 8083")

	err := http.ListenAndServe(":8083", nil)
	if err != nil {
		panic(err)
	}
}
