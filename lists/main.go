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

	err := client.RPush(ctx, "golist:queue", 2).Err()
	if err != nil {
		panic(err)
	}
	// err1 := client.RPush(ctx, "golist:queue", 1,3,4,5,6,7,8).Err()
	// if err1 != nil {
	// 	panic(err1)
	// }
	// err2:=client.LPop(ctx,"golist:queue").Err()
	// if err2!=nil{
	// 	panic(err2)
	// }

	// errmove:=client.LTrim(ctx,"golist:queue",0,2).Err()
	// if errmove!=nil{
	// 	panic(errmove)
	// }
	// fmt.Println("golist:queue trim",valmove)

	// errmove:=client.LMove(ctx,"golist:queue","golist:stack","left","left").Err()
	// if errmove!=nil{
	// 	panic(errmove)
	// }
	// errblockpop:=client.BLPop(ctx,10,"golist:queue").Err()
	// if errblockpop!=nil{
	// 	panic(errblockpop)
	// }

	val3,err3:=client.LLen(ctx,"golist:queue").Result()	
	if err3!=nil{
		panic(err3)
	}
	fmt.Println("golist:queue length",val3)
	val1,err1:=client.LRange(ctx,"golist:queue",0,val3-1).Result()
	if err1!=nil{
		panic(err1)
	}
	fmt.Println("golist:queue range",val1)
	val,err:=client.LLen(ctx,"golist:stack").Result()
	if err!=nil{
		panic(err)
	}
	fmt.Println("golist:stack",val)






	

	fmt.Println(client)

	// This is a placeholder for the main function in the lists package.
	// You can implement your Redis list operations here.
}
