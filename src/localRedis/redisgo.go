package localRedis

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var (
	RedisHandle *redis.Client
)

func SetupDB() {
	RedisHandle = GetRedisConn()

	fmt.Printf("%p", RedisHandle)
}

//获取redis集群客户端
func GetRedisConn() *redis.Client {
	if RedisHandle != nil {
		_, err := RedisHandle.Ping().Result()
		//glog.Info("redis connections: ", redisHandle.PoolStats().TotalConns)
		if err == nil {
			return RedisHandle
		}
	}

	var db = 0
	var addr = "10.1.1.245:6979"
	var pwd = "qlwhIPO82@#KDFQAwe"

	//glog.Info("redis connections:" + addr + ",paw:" + pwd)

	RedisHandle = redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     pwd,
		DB:           db,
		MinIdleConns: 28,
		IdleTimeout:  30,
		PoolSize:     512,
		MaxConnAge:   30 * time.Second,
	})
	_, err := RedisHandle.Ping().Result()
	if err != nil {
		panic(err)
	}

	return RedisHandle
}
