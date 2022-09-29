package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/golang/glog"
	"github.com/limitedlee/microservice/common/config"
	"github.com/spf13/cast"
	"strings"
	"sync"
	"time"
)

//获取redis集群客户端
func GetRedisConn() *redis.Client {
	//if redisHandle != nil {
	//	_, err := redisHandle.Ping().Result()
	//	//glog.Info("redis connections: ", redisHandle.PoolStats().TotalConns)
	//	if err == nil {
	//		return redisHandle
	//	}
	//}

	var db = cast.ToInt(config.GetString("redis.DB"))
	var addr = config.GetString("redis.Addr")
	var pwd = config.GetString("redis.Password")

	//glog.Info("redis connections:" + addr + ",paw:" + pwd)

	redisHandle := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       db,
	})
	_, err := redisHandle.Ping().Result()
	if err != nil {
		glog.Error("redis error: ", err.Error())
		if strings.Contains(err.Error(), "timeout") { // 超时重试，之前超时的时候直接panic了
			for i := 0; i < 5; i++ {
				time.Sleep(1 * time.Second)
				redisHandle = redis.NewClient(&redis.Options{
					Addr:     addr,
					Password: pwd,
					DB:       db,
				})
				_, err := redisHandle.Ping().Result()
				glog.Error("循环五次：", err)
				if err == nil {
					return redisHandle
				}
			}
		}

		panic(err) // 循环五次还不行的话直接panic
	}

	return redisHandle
}

func main() {

	var wait sync.WaitGroup

	for i := 0; i < 100; i++ {
		wait.Add(1)
		go func() {
			defer func() {
				wait.Done()
			}()
			conn := GetRedisConn()
			get := conn.Get("name")
			fmt.Println(get.String())
		}()
	}

	wait.Wait()

}
