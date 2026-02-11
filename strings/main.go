package main

import (
	"context"
	"fmt"
	//  "time"

	"github.com/redis/go-redis/v9"
)

func main() {    
    client := redis.NewClient(&redis.Options{
        Addr:	  "localhost:6379",
        Password: "", // No password set
        DB:		  0,  // Use default DB
        Protocol: 2,  // Connection protocol
    })

	ctx := context.Background()

err := client.Set(ctx, "cars:1", "Toyota", 0).Err()
if err != nil {
    panic(err)
}
err1 := client.Set(ctx, "cars:2", "Honda", 0).Err()
if err1 != nil {
    panic(err)
}

val, err := client.Get(ctx, "cars:1").Result()
if err != nil {
    panic(err)
}
fmt.Println("cars:1", val)

// 	val1, err1 := client.Get(ctx, "cars:2").Result()
// if err1 != nil {
//     panic(err1)
// }

// fmt.Println("cars:2", val1)

// err2:=client.MSet(ctx,"cars:3","ford","cars:4","kia","cars:5","suzuki").Err()
// if err2!=nil{
// 	panic(err2)
// }


// vals, err := client.MGet(ctx, "cars:3", "cars:4", "cars:5").Result()
// if err != nil {
// 	panic(err)
// }

// fmt.Println("cars:4","cars:3","cars:5",vals)


// ok, err := client.MSetNX(ctx,
// 	"cars:10", "bmw",
// 	"cars:11", "nissan",
// ).Result()

// if err != nil {
// 	panic(err)
// }

// // er:=client.SetXX(ctx,"cars:10","byd",0).Err()
// // if er!=nil{
// // 	panic(er)
// // }
// // val10, err := client.Get(ctx, "cars:10").Result()
// // if err != nil {
// // 	panic(err)
// // }
// // fmt.Println("cars:10", val10)

// fmt.Println("Were keys set?", ok)

// vals2, err := client.MGet(ctx, "cars:11", "cars:10").Result()
// if err != nil {
// 	panic(err)
// }

// fmt.Println(vals2)

// errn:=client.MSetNX(ctx,"cars:12","mercedes","cars:13","audi").Err()
// if errn!=nil{
// 	panic(errn)
// }
// errn2:=client.MSetNX(ctx,"cars:12","figo","cars:13","audi").Err()
// if errn2!=nil{
// 	panic(errn2)
// }

// vals3, errn := client.MGet(ctx, "cars:12", "cars:13").Result()
// if errn != nil {
// 	panic(err)
// }

// // errnew:=client.IncrBy(ctx,"counter",10).Err()
// // if errnew!=nil{
// // 	panic(errnew)
// // }
// valnew, errnew := client.Get(ctx, "counter").Result()
// if errnew != nil {
// 	panic(errnew)
// }
// fmt.Println("counter", valnew)

// fmt.Println(vals3)

// errlatest:=client.MSet(ctx,"sims:1","Airtel","sims:2","Jio").Err()
// if errlatest!=nil{
// 	panic(errlatest)
// }	
// iter := client.Scan(ctx, 0, "sims:*", 0).Iterator()
// for iter.Next(ctx) {
//     err := client.Del(ctx, iter.Val()).Err()
//     if err != nil {
//         panic(err)
//     }
// }
// if err = iter.Err(); err != nil {
//     panic(err)
// }
// //expiration
// errEXP:=client.Set(ctx,"tempkey","tempvalue",10*time.Second).Err()
// if errEXP!=nil{	
// 	panic(errEXP)}
// // expcars:=client.Expire(ctx,"cars:10",30*time.Second).Err()
// // if expcars!=nil{
// // 	panic(expcars)
// // }
// // 	expcars=client.PExpireAt(ctx,"cars:7",time.Now().Add(3000*time.Millisecond)).Err()
// // 	if expcars!=nil{
// // 		panic(expcars)
// // 	}

// valcar,err:=client.TTL(ctx,"cars:10").Result()
// if err!=nil{
// 	panic(err)
// }
// fmt.Println("TTL of cars:10",valcar)



	fmt.Println(client)

	
}




