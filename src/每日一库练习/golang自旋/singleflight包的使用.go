package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/singleflight"
	"log"
	"strconv"
	"sync"
	"time"
)

/**
参考文档  https://www.lixueduan.com/posts/go/singleflight/
Go语言之防缓存穿透利器Singleflight
1. 缓存击穿
	某个热点数据缓存失效后，同一时间的大量请求直接被打到的了DB，会给DB造成极大压力，甚至直接打崩DB。

2： 常见的解决方案是加锁
	虽然加锁能解决问题，但是太重了，而且逻辑比较复杂，又是加锁又是等待的。相比之下 singleflight 就是一个轻量级的解决方案。


*/

var (
	g            singleflight.Group
	ErrCacheMiss = errors.New("cache miss")
)

func main() {

	var wg sync.WaitGroup
	wg.Add(10)

	// 模拟10个并发
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			data, err := load("key")
			if err != nil {
				log.Print(err)
				return
			}
			log.Println(data)
		}()
	}

	/*
		上面执行完成之后会删除key， 后面如果有类似的请求，还是会重新查询数据库
	*/
	time.Sleep(3 * time.Second)
	data, err := load("key")
	fmt.Println("测试key有没有删除： ", data, err)
	wg.Wait()
}

// 获取数据
func load(key string) (string, error) {
	data, err := loadFromCache(key)
	if err != nil && err == ErrCacheMiss {
		// 利用 singleflight 来归并请求
		v, err, _ := g.Do(key, func() (interface{}, error) {
			data, err := loadFromDB(key)
			if err != nil {
				return nil, err
			}
			setCache(key, data)
			return data, nil
		})
		if err != nil {
			log.Println(err)
			return "", err
		}
		data = v.(string)
	}
	return data, nil
}

// getDataFromCache 模拟从cache中获取值 cache miss
func loadFromCache(key string) (string, error) {
	return "", ErrCacheMiss
}

// setCache 写入缓存
func setCache(key, data string) {}

// getDataFromDB 模拟从数据库中获取值
func loadFromDB(key string) (string, error) {
	fmt.Println("query db")
	unix := strconv.Itoa(int(time.Now().UnixNano()))
	return unix, nil
}

/**
todo 改进
singleflight 内部使用 waitGroup 来让同一个 key 的除了第一个请求的后续所有请求都阻塞。直到第一个请求执行 fn 返回后，其他请求才会返回。
这意味着，如果 fn 执行需要很长时间，那么后面的所有请求都会被一直阻塞。
这时候我们可以使用 DoChan 结合 ctx + select 做超时控制

*/
func loadChan(ctx context.Context, key string) (string, error) {
	data, err := loadFromCache(key)
	if err != nil && err == ErrCacheMiss {
		// 使用 DoChan 结合 select 做超时控制
		result := g.DoChan(key, func() (interface{}, error) {
			data, err := loadFromDB(key)
			if err != nil {
				return nil, err
			}
			setCache(key, data)
			return data, nil
		})
		select {
		case r := <-result:
			return r.Val.(string), r.Err
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}
	return data, nil
}
