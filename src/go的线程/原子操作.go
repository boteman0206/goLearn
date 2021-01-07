package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var x int32
var wait1 sync.WaitGroup

func addX() {
	wait1.Done()
	//x ++  // 每次的数值都不一样
	atomic.AddInt32(&x, 1)

}

func main() {

	for i := 0; i < 1000; i++ {
		wait1.Add(1)
		go addX()
	}

	wait1.Wait()
	fmt.Println(x)

	//比较和交换
	swapped := atomic.CompareAndSwapInt32(&x, 1000, 200)
	fmt.Println(swapped, x)

}
