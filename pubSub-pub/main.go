package main
import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})

	ctx := context.Background()

	err:=client.Publish(ctx,"GoMessage","go to elsewhere").Err()
	if err!=nil{
		panic(err)
	}
	fmt.Println(client)
}


