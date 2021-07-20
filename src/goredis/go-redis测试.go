package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
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

	//var db = cast.ToInt(config.GetString("redis.DB"))
	//var addr = config.GetString("redis.Addr")

	addr := "10.1.1.242:6679"
	pwd := "MkdGH*3ldf"
	db := 0
	//var pwd = config.GetString("redis.Password")

	//glog.Info("redis connections:" + addr + ",paw:" + pwd)

	redisHandle1 := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     pwd,
		DB:           db,
		MinIdleConns: 28,
		IdleTimeout:  30,
		PoolSize:     512,
		MaxConnAge:   30 * time.Second,
	})

	//redisHandle1.
	_, err := redisHandle1.Ping().Result()
	if err != nil {
		panic(err)
	}

	return redisHandle1
}

func main() {
	conn := GetRedisConn()
	defer conn.Close()

	get := conn.Get("testpw")
	//conn.Do()
	fmt.Println("get : ", get)

	return
	// 使用pipeline管道
	//pipeline := conn.Pipeline()
	//defer pipeline.Close()
	//i := pipeline.HLen("warehouse:store:relation").Val()
	//fmt.Println(i)
	//result, _ := pipeline.Exec()
	//cmder := result[0]
	//intData := cmder.(*redis.IntCmd)
	//
	//fmt.Println(len(result), cap(result), intData.Val())

	// 直接获取的hlen长度
	val := conn.HLen("warehouse:store:relation").Val()
	fmt.Println("val: ", val)

	scan := conn.HScan("warehouse:store:relation", 0, "", val*2)
	keys, cursor, err := scan.Result()
	fmt.Println("data : ", keys, cursor, err)

	allStrings := make(map[string]string, 0)

	lenData := len(keys)
	key := ""
	for i := 0; i < lenData; i++ {
		if i%2 == 0 {
			key = keys[i]
			fmt.Println("keys: ", key)
		} else {
			fmt.Println("value: ", keys[i])
			allStrings[key] = keys[i]
		}

	}

	fmt.Println("mapData: ", len(allStrings), len(keys), allStrings)
	bytes, err := json.Marshal(allStrings)
	fmt.Println(string(bytes))
}
