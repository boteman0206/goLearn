package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"sync"
	"time"
)

type MyMap struct {
	store map[string]string
	lock  sync.RWMutex
}

func (m *MyMap) Get(key string) string {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.store[key]
}

func (m *MyMap) Set(key, value string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.store[key] = value
}

func main() {

	//myMap := MyMap{
	//	make(map[string]string, 0),
	//	sync.RWMutex{},
	//}
	//myMap.Set("name", "jack")
	//
	//fmt.Println(myMap)
	//
	//
	//
	//
	//type data struct {
	//	Name string `json:"name"`
	//	Age int32 `json:"age"`
	//}
	//
	//d := data{"jack", 19}
	//
	//fmt.Printf("%+v ", d)
	//fmt.Printf("%v ", d)
	//fmt.Printf("%#v ", d)

	limiter := rate.NewLimiter(20, 1)
	fmt.Println(limiter)

	start := time.Now()
	for i := 0; i <= 100; i++ {
		r := limiter.Reserve()
		if !r.OK() {
			fmt.Println("---------------------限流退出----------")
		}
		time.Sleep(r.Delay())

	}
	end := time.Now()
	fmt.Println(end.Sub(start).Seconds())

}
