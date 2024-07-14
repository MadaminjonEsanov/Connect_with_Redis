package main

import (
	"fmt"
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {

	redisCli := createRedisCli()

	_, err := redisCli.Set(ctx, "salom", "Hello,World", 1*time.Minute).Result()
	if err != nil {
	 log.Println("error in set data to redis!", err)
	 return
	}
   
	data, err := redisCli.Get(ctx, "salom").Result()
	if err != nil {
	 log.Println("error in get data to redis!", err)
	 return
	}
   
	fmt.Println(data)
   }



func createRedisCli() *redis.Client {
	
	option := &redis.Options{

		Addr:     "localhost:6379",
        Password: "",                // no password set
		DB:       0,                 // use default DB
	}  

	redisCli := redis.NewClient(option)

	msg, err := redisCli.Ping(context.Background()).Result()

	if err!= nil {
		log.Println("error on connecting to redis: ",err)
		return nil
	}

	log.Println("succesfully connected to redis: ",msg)

	return redisCli
}