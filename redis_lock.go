package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6380",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("db init successful")

	if BoolCmd := rdb.SetNX(ctx, "key01", "val01", 0); !BoolCmd.Val() {
		log.Fatal("setnx false")
	} else {
		fmt.Println("setnx ok")
	}

	result, err := rdb.Get(ctx, "key01").Result()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(result)
	}
}

/*
PS D:\gopro\wbsTest\v1\redis> go run .\redis_lock.go
db init successful
setnx ok
val01
PS D:\gopro\wbsTest\v1\redis> go run .\redis_lock.go
db init successful
2024/03/18 09:37:19 setnx false
exit status 1
PS D:\gopro\wbsTest\v1\redis>
*/
