package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis"
	ratelimiter "github.com/teambition/ratelimiter-go"
)

//Implements RedisClient for redis.Client
type redisClient struct {
	*redis.Client
}

//
func (c *redisClient) RateDel(key string) error {
	return c.Del(key).Err()
}
func (c *redisClient) RateEvalSha(sha1 string, keys []string, args ...interface{}) (interface{}, error) {
	return c.EvalSha(sha1, keys, args...).Result()
}
func (c *redisClient) RateScriptLoad(script string) (string, error) {
	return c.ScriptLoad(script).Result()
}

func GetRedisConn() *redis.Client {

	//addr = "59.110.239.18:6379"
	//pwd = "PC6W58qYof"
	addr := "10.1.1.245:6979"
	pwd := "qlwhIPO82@#KDFQAwe"

	redisHandle := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     pwd,
		DB:           0,
		MinIdleConns: 10,
		IdleTimeout:  60,
	})
	_, err := redisHandle.Ping().Result()
	if err != nil {
		panic(err)
	}

	return redisHandle
}

var LocalRedisClient *redis.Client
var limiter *ratelimiter.Limiter

func init() {
	LocalRedisClient = GetRedisConn()

	limiter = ratelimiter.New(ratelimiter.Options{
		Max:      20,
		Duration: time.Second, // 每秒20
		Client:   &redisClient{LocalRedisClient},
	})

}

type Result struct {
	Lock sync.Mutex
	Data []int
}

func (r *Result) Add(num int) {
	r.Lock.Lock()
	defer r.Lock.Unlock()

	r.Data = append(r.Data, num)
}

//实际用法

var (
	Prefix = "获取限流异常: "
	//ele批量删除商品接口  接口限流20次/source/秒
	EleSkuDelete = "ele.sku.delete"
)

func EleDeleteProductRate() bool {

	for i := 0; i < 10; i++ {
		res, err := limiter.Get(EleSkuDelete)
		if err != nil {
			fmt.Println(Prefix, " 饿了么删除商品：", err.Error())
		}
		if res.Remaining > 0 {
			return true
		} else {
			//fmt.Println("等待时间：", res.Duration.Seconds())
			//fmt.Println("等待时间：", res.Reset.Sub(time.Now()).Seconds())
			time.Sleep(res.Duration)
		}
	}

	return false
}

func main() {
	var data = Result{
		Lock: sync.Mutex{},
		Data: make([]int, 0),
	}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer func() {
				wg.Done()
			}()
			ok := EleDeleteProductRate()
			if ok {
				data.Add(i)
				fmt.Println("id: 可以执行 ", i)
			} else {
				fmt.Println("失败的id： ", i)
			}
		}(i)
		//ok := EleDeleteProductRate()
		//if ok {
		//	data.Add(i)
		//	fmt.Println("id: 可以执行 ", i)
		//} else {
		//	fmt.Println("失败的id： ", i)
		//}
	}

	wg.Wait()
	marshal, err := json.Marshal(data)
	if err != nil {
		return
	}
	fmt.Println("len: ", len(data.Data), " shuju: ", string(marshal))
}

func main1() {
	// use memory
	// limiter := ratelimiter.New(ratelimiter.Options{
	// 	Max:      10,
	// 	Duration: time.Minute, // limit to 1000 requests in 1 minute.
	// })

	// or use redis
	conn := GetRedisConn()

	limiter := ratelimiter.New(ratelimiter.Options{
		Max:      10,
		Duration: time.Second, // limit to 1000 requests in 1 minute.
		Client:   &redisClient{conn},
	})

	var data = Result{
		Lock: sync.Mutex{},
		Data: make([]int, 0),
	}
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer func() {
				wg.Done()
			}()
			for true {
				res, err := limiter.Get("test-0001")
				marshal, _ := json.Marshal(res)
				//fmt.Println("等待时间：", res.Reset.Sub(time.Now()).Seconds())
				//time.Sleep(time.Duration(res.Reset.Sub(time.Now()).Seconds()) * time.Second)
				time.Sleep(res.Duration)
				if res.Remaining > 0 {
					data.Add(i)
					fmt.Println("id: 可以执行 ", i, "    ", string(marshal), " ", err)
					break
				} else {
					//fmt.Println("没有执行", i)
					//fmt.Println("等待时间：", res.Duration.Seconds())
					//fmt.Println("等待时间：", res.Reset.Sub(time.Now()).Seconds())
					time.Sleep(time.Duration(res.Reset.Sub(time.Now()).Seconds()) * time.Second)
				}
			}
		}(i)
	}

	//time.Sleep(20 * time.Second)
	wg.Wait()
	marshal, err := json.Marshal(data)
	if err != nil {
		return
	}
	fmt.Println("len: ", len(data.Data), " shuju: ", string(marshal))
}
