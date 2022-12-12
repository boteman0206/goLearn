package main

import (
	"container/list"
	"sync"
)

/**
https://juejin.cn/post/7043603199705481247

LRU 是 Least Recently Used 的简写，字面意思则是最近最少使用。

实现原理没什么好说的，和 Java 的一样：
	一个双向链表存储数据的顺序
	一个 map 存储最终的数据
	当数据达到上限时移除链表尾部数据
	将使用到的 Node 移动到链表的头结点

*/

type LruCache struct {
	size     int
	values   *list.List
	cacheMap map[interface{}]*list.Element
	lock     sync.Mutex
}

func NewLruList(size int) *LruCache {
	values := list.New()
	return &LruCache{
		size:     size,                                      // size 存储缓存大小。
		values:   values,                                    // 链表存储数据顺序。
		cacheMap: make(map[interface{}]*list.Element, size), // map 存储数据。
		lock:     sync.Mutex{},                              // lock 用于控制并发安全。
	}
}

func (c *LruCache) Put(k, v interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.values.Len() == c.size { // 超过容量
		back := c.values.Back()
		c.values.Remove(back)

		delete(c.cacheMap, k)

	}

	front := c.values.PushFront(v)
	c.cacheMap[k] = front

}

func (c *LruCache) Get(k interface{}) (interface{}, bool) {

	element, ok := c.cacheMap[k]
	if ok {
		c.values.MoveToFront(element)
		return element, ok
	} else {
		return nil, false
	}
}

func main() {

}
