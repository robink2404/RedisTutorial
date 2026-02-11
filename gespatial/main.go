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

	err:=client.GeoAdd(ctx,"goGeo:myLocation",&redis.GeoLocation{
    Longitude: 2.3522,    // Paris, France
    Latitude: 48.8566,
    Name: "france",
}).Err()
if err!=nil{
    panic(err)
}
err=client.GeoAdd(ctx,"goGeo:myLocation",&redis.GeoLocation{
    Longitude: 10.4515,   // Berlin, Germany
    Latitude: 52.5200,
    Name: "germany",
}).Err()
if err!=nil{
    panic(err)
}

res,err:=client.GeoSearch(ctx,"goGeo:myLocation",&redis.GeoSearchQuery{
	Longitude: 10.4515,
	Latitude: 52.5210,
	Radius: 5000,
	RadiusUnit: "km",
}).Result()
if err!=nil{
	panic(err)
}	
fmt.Println("Locations within 30 km of Berlin:", res)

	
	fmt.Println(client)
}

