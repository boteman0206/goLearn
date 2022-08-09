package main

import (
	"fmt"
	"sync"
	"time"
)

type TokenBucket struct {
	rate         int64 //固定的token放入速率, r/s
	capacity     int64 //桶的容量
	tokens       int64 //桶中当前token数量
	lastTokenSec int64 //上次向桶中放令牌的时间的时间戳，单位为秒

	lock sync.Mutex
}

func (bucket *TokenBucket) Take() bool {
	bucket.lock.Lock()
	defer bucket.lock.Unlock()

	now := time.Now().Unix()                                              // 当前时间
	bucket.tokens = bucket.tokens + (now-bucket.lastTokenSec)*bucket.rate // 先添加令牌
	if bucket.tokens > bucket.capacity {
		bucket.tokens = bucket.capacity
	}
	bucket.lastTokenSec = now
	if bucket.tokens > 0 {
		// 还有令牌，领取令牌
		bucket.tokens--
		return true
	} else {
		// 没有令牌,则拒绝
		return false
	}
}

/**
  New生成新的令牌桶
*/
func NewBucket(rate, cap int64) *TokenBucket {

	bucket := &TokenBucket{
		rate:         rate,
		capacity:     cap,
		tokens:       0,
		lastTokenSec: time.Now().Unix(),
		lock:         sync.Mutex{},
	}
	return bucket
}

func main() {

	//测试使用
	Test01()
	//fmt.Println("time.Now().Unix(): ", time.Now().Unix())

}

func Test01() {

	newBucket := NewBucket(10, 10)

	for i := 0; i < 100; i++ {
		take := newBucket.Take()
		fmt.Println("take: ", take, " 数量: ", i)
	}
	unix := time.Unix(newBucket.lastTokenSec, 0)
	format := unix.Format("2006-01-02 03:04:05")
	fmt.Println("newBucket.lastTokenSec: ", format)
}
