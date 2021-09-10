package main

import (
	"github.com/go-redis/redis"
	"log"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "139.224.1.155:6379",
		Password: "comma-redis", // no password set
		DB:       15,            // use default DB
		PoolSize: 10,
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		log.Fatalf("redis comm error", err)
	}

}

func main() {
	rdb.Set("aa", 1, 0)
}
