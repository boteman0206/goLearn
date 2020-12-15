package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

/**
redis
*/
var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

func main() {

	get := pool.Get()
	defer get.Close()

	fmt.Println("get pool : ", get)
	reply, err := redis.String(get.Do("get", "name1"))
	fmt.Println(reply, " ", err)
}
