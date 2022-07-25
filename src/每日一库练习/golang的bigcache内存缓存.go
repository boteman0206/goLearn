package main

import (
	"fmt"
	"github.com/allegro/bigcache/v3"
	"time"
)

/**
LifeWindow & CleanWindow

1: LifeWindow is a time. After that time, an entry can be called dead but not deleted.  // todo 可以被驱逐但是不会被删除
2: CleanWindow is a time. After that time, all the dead entries will be deleted, but not the entries that still have life.


*/
func main() {

	DefaultConfig()

	//MyConfig()
}

//自定义cache

func MyConfig() {
	config := bigcache.Config{
		// number of shards (must be a power of 2)
		Shards: 1024,

		// time after which entry can be evicted
		LifeWindow: 6 * time.Second,

		// Interval between removing expired entries (clean up).
		// If set to <= 0 then no action is performed.
		// Setting to < 1 second is counterproductive — bigcache has a one second resolution.
		CleanWindow: 0, // todo 设置为0则不执行任何清除key的操作 如果设置为<=0，则不执行任何操作。

		// rps * lifeWindow, used only in initial memory allocation
		MaxEntriesInWindow: 1000 * 10 * 60,

		// max entry size in bytes, used only in initial memory allocation
		MaxEntrySize: 500,

		// prints information about additional memory allocation
		Verbose: true,

		// cache will not allocate more memory than this limit, value in MB
		// if value is reached then the oldest entries can be overridden for the new ones
		// 0 value means no size limit
		HardMaxCacheSize: 8192,

		// callback fired when the oldest entry is removed because of its expiration time or no space left
		// for the new entry, or because delete was called. A bitmask representing the reason will be returned.
		// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
		OnRemove: nil,

		// OnRemoveWithReason is a callback fired when the oldest entry is removed because of its expiration time or no space left
		// for the new entry, or because delete was called. A constant representing the reason will be passed through.
		// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
		// Ignored if OnRemove is specified.
		OnRemoveWithReason: nil,
	}

	cache, _ := bigcache.NewBigCache(config)

	cache.Set("name", []byte("jack"))

	get, err := cache.Get("name")
	if err != nil {
		return
	}
	fmt.Println(string(get))

	time.Sleep(7 * time.Second)

	get, err = cache.Get("name")
	if err != nil {
		return
	}
	fmt.Println("sleep之后：", string(get))

	cache.Delete("name")

	get, err = cache.Get("name")
	if err != nil {
		fmt.Println("删除之后：", err.Error())
		return
	}

	fmt.Println("删除之后：", string(get))

}

func DefaultConfig() {

	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(1 * time.Minute))

	cache.Set("my-unique-key", []byte("value"))

	entry, _ := cache.Get("my-unique-key")
	fmt.Println(string(entry))

	//err := cache.Delete("my-unique-key")
	//if err != nil {
	//	return
	//}

	cache.Set("my-unique-key", []byte("my new value"))
	entry2, _ := cache.Get("my-unique-key")
	fmt.Println("覆盖之后的value: ", string(entry2))

	cache.Reset()
	get, err := cache.Get("my-unique-key")
	if err != nil {
		return
	}
	fmt.Println("清空之后：", get)

	time.Sleep(1 * time.Minute)

	entry1, _ := cache.Get("my-unique-key")
	fmt.Println("sleep之后： ", entry1)

	//err := cache.Delete("my-unique-key")
	//if err != nil {
	//	return
	//}

}
