package main

import (
	"fmt"
	lru "github.com/hashicorp/golang-lru"
)

func main() {
	cache, ok := lru.New(3)
	fmt.Println(cache, ok)
	for i := 0; i < 3; i++ {
		cache.Add(i, i*10)
	}
	if cache.Len() != 3 {
		panic(fmt.Sprintf("bad len: %v", cache.Len()))
	}
	fmt.Println(cache.Len())
	for i := 0; i < 3; i++ {
		value, o := cache.Get(i)
		fmt.Println("----", value, o)
	}

	cache.Add(8, 8*10) // 容量只有3，这里增加了一个8 会吧最前面的0移除掉。
	value, o := cache.Get(0)
	fmt.Println("0", value, o) // 这里久获取不到0了，  只能获取到  1，2，8， 其中8会在链表的最前面
	value8, o := cache.Get(8)
	fmt.Println(value8, o)
}
