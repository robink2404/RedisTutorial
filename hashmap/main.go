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

	err := client.HSet(ctx, "gohash:bike1","model","yamaha","brand","yamaha","price",200000).Err()
	if err != nil {
		panic(err)
	}
	err=client.HSet(ctx,"gohash:bike1","yearOfPurchase",2023).Err()

	if err != nil {
		panic(err)
	}
	val,err:=client.HGet(ctx,"gohash:bike1","price").Result()

	if err!=nil{
		panic(err)
	}
	fmt.Println("gohash:bike1 price",val)
	val1,err:=client.HMGet(ctx,"gohash:bike1","model","brand").Result()
	if err!=nil{
		panic(err)
	}
	fmt.Println("gohash:bike1 model and brand",val1)
	val2,err:=client.HGetAll(ctx,"gohash:bike1").Result()
	if err!=nil{
		panic(err)
	}
	fmt.Println("gohash:bike1 all fields",val2)
	err=client.HIncrBy(ctx,"gohash:bike1","price",2000).Err()

	if err!=nil{
		panic(err)
	}
	val3,err:=client.HGet(ctx,"gohash:bike1","price").Result()
	if err!=nil{
		panic(err)
	}
	fmt.Println("gohash:bike1 price after increment",val3)
	fmt.Println(client)
}

