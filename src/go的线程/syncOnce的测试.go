package main

import (
	"sync"
	"sync/atomic"
)
import "fmt"

// todo  不会打印任何东西，sync.Once.Do 方法中传入的函数只会被执行一次，哪怕函数中发生了 panic；
func panicDo() {
	once := &sync.Once{}
	defer func() {
		if err := recover(); err != nil {
			once.Do(func() {
				fmt.Println("run in recover")
			})
		}
	}()
	once.Do(func() {
		panic("panic i=0")
	})

}

//todo 自己实现的sync.once
type MyOnce struct {
	flag uint32
	lock sync.Mutex
}

func (m *MyOnce) Do(f func()) {
	if atomic.LoadUint32(&m.flag) == 0 {
		m.lock.Lock()
		defer m.lock.Unlock()
		if atomic.CompareAndSwapUint32(&m.flag, 0, 1) {
			f()
		}

	}

}

func main() {

	panicDo()

	// 测试自己实现的sync.Once
	once := MyOnce{}

	for i := 0; i < 10; i++ {
		once.Do(func() {
			fmt.Println("this is my sync.Once")
		})
	}

}
