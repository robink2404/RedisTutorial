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

	err := client.ZAdd(ctx, "goSortedSet:list1", redis.Z{5,"kareena"},redis.Z{2,"priyanka"},redis.Z{3,"deepika"}).Err()
	if err != nil {
		panic(err)
	}
	err=client.ZAdd(ctx,"goSortedSet:list1",redis.Z{10,"alia"}).Err()
	if err!=nil{
		panic(err)
	}
	val,err:=client.ZRange(ctx,"goSortedSet:list1",0,-1).Result()

	if err!=nil{
		panic(err)
	}
	fmt.Println("goSortedSet:list1 members",val)
	val,err=client.ZRevRange(ctx,"goSortedSet:list",0,-1).Result()
	if err!=nil{
		panic(err)
	}
	fmt.Println("goSortedSet:list1 members in reverse order",val)
	err=client.ZRem(ctx,"goSortedSet:list1","priyanka").Err()
	if err!=nil{
		panic(err)
	}
	val,err=client.ZRange(ctx,"goSortedSet:list1",0,-1).Result()
	if err!=nil{
		panic(err)
	}
	rank, err := client.ZRank(ctx, "goSortedSet:list1", "alia").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("goSortedSet:list1 rank of alia", rank)
	fmt.Println(client)
}
