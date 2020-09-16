package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

var ctx = context.Background()

type User struct {
	Name string
}

func main() {
	opt, err := redis.ParseURL("")
	if err != nil {
		log.Fatal(err)
	}
	opt.Username = ""
	client := redis.NewClient(opt)
	pong, err := client.Ping(ctx).Result()
	fmt.Println(err)
	fmt.Println(pong)

	user := User{Name: "Joan"}

	userJson, _ := json.Marshal(user)

	client.Set(ctx, "key", "value", 0)       // 0 means never
	client.Set(ctx, "userData", userJson, 0) // 0 means never

	fmt.Println(client.Get(ctx, "key"))
	fmt.Println(client.Get(ctx, "userData").Result())

}
