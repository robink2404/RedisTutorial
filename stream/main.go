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

	// res1,err:=client.XAdd(ctx,&redis.XAddArgs{
	// 	Stream:"gostream:sensorData",
	// 	ID:"*",
	// 	Values:map[string]interface{}{
	// 		"temperature":35,
	// 		"humidity":80,
	// 	},
	// }).Result()
	// if err!=nil{
	// 	panic(err)
	// }
	// fmt.Println("gostream:sensorData added entry ID",res1)

	// res2,err:=client.XAdd(ctx,"gostream:sensorData",&redis.XAddArgs{
	// 	Values: map[string]interface{}{
	// 		"temperature":30,
	// 		"humidity":65,
	// 	},
	// }).Result()
	// if err!=nil{
	// 	panic(err)
	// }
	// fmt.Println("gostream:sensorData added entry ID",res2)
	// res2, err := client.XAdd(ctx, &redis.XAddArgs{
	// 	Stream: "race:france",
	// 	Values: map[string]interface{}{
	// 		"rider":       "Castilla",
	// 		"speed":       30.2,
	// 		"position":    1,
	// 		"location_id": 1,
	// 	},
	// }).Result()

	// if err != nil {
	// 	panic(err)
	// }
	// 	fmt.Println(res2)

	// res, err := client.XRangeN(ctx, "gostream:sensorData", "1770768785525-0", "+", 2).Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("gostream:sensorData range query (max 2)", res)

	res1,err:=client.XRead(ctx,&redis.XReadArgs{
		Streams: []string{"gostream:sensorData","0"},
		Count: 2,
		Block: 5000,

	}).Result()
	if err!=nil{
		panic(err)
	}
	fmt.Println("gostream:sensorData read",res1)

	res3,err:=client.XLen(ctx,"gostream:sensorData").Result()
	if err!=nil{
		panic(err)
	}
	fmt.Println("gostream:sensorData length",res3)
	res4, err := client.XRange(ctx, "gostream:sensorData", "1770768785525-0","770768977008-0").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("gostream:sensorData range entries",res4)

	fmt.Println(client)
}
