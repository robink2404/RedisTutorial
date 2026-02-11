
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

	err := client.SAdd(ctx, "goset:cars", "tata","enova","ford","tesla").Err()
	if err != nil {
		panic(err)
	}
   err=client.SAdd(ctx,"goset:cars","tata","mahindra","suzuki").Err()

	if err != nil {
		panic(err)
	}
	val,err:=client.SMembers(ctx,"goset:cars").Result()
	if err!=nil{
		panic(err)
	}
	fmt.Println("goset:cars members",val)

	err=client.SAdd(ctx,"golist:indiancars","mahindra","tata","ashok").Err()
	if err!=nil{
		panic(err)
	}

	val1,err:=client.SMembers(ctx,"golist:indiancars").Result()
	if err!=nil{
		panic(err)
	}
	fmt.Println("golist:indiancars members",val1)

	err=client.SAdd(ctx,"goset:cars","byd","kia","honda","maruti").Err()
	if err!=nil{
		panic(err)
	}
	err=client.SRem(ctx,"goset:cars","tesla","byd").Err()
	if err!=nil{
		panic(err)
	}
	val2,err:=client.SIsMember(ctx,"goset:cars","tata").Result()
	if err!=nil{
		panic(err)
	}
	fmt.Println("is tesla member of goset:cars?",val2)
	val3,err:=client.SCard(ctx,"goset:cars").Result()
	if err!=nil{
		panic(err)
	}
	val4,cursor,err:=client.SScan(ctx,"goset:cars",0,"",0).Result()
	if err!=nil{
		panic(err)
	}
	fmt.Println("goset:cars scan",val4,cursor)
	fmt.Println("goset:cars cardinality",val3)
	err=client.SAdd(ctx,"goset:bharatcars","mahindra","tata","ashok").Err()
	if err!=nil{
		panic(err)
	}
	val6,err:=client.SUnion(ctx,"goset:indiancars","golist:bharatcars").Result()
	if err!=nil{
		panic(err)
	}
	fmt.Println("union of golist:indiancars and golist:bharatcars",val6)



	val5,err:=client.SInter(ctx,"goset:cars","goset:bharatcars").Result()
	if err!=nil{
		panic(err)
	}
	fmt.Println("intersection of goset:cars and goset:bharatcars",val5)

	fmt.Println(client)
}
