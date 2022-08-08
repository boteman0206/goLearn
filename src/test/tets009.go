package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := make(map[string]int)
	var rwLock *sync.RWMutex = new(sync.RWMutex)
	go func() { //开一个goroutine写map
		for j := 0; j < 1000000; j++ {
			rwLock.Lock()
			c[fmt.Sprintf("%d", j)] = j
			rwLock.Unlock()
		}
	}()
	go func() { //开一个goroutine读map
		for j := 0; j < 1000000; j++ {
			rwLock.RLock()
			fmt.Println(c[fmt.Sprintf("%d", j)])
			rwLock.RUnlock()
		}
	}()
	time.Sleep(time.Second * 20)
}
