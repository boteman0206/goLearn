package main

import (
	"fmt"
	"sync"
)

var smap sync.Map // 并发安全的map

func main() {

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("name%d", i)
		smap.Store(key, i)
		smap.Load(key)
	}

	fmt.Println(smap)
}
