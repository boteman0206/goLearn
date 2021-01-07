package main

import (
	"fmt"
	"sync"
	"time"
)

/**
sync.Once
*/

var one sync.Once // 只需要一次操作的时候

func test() {
	fmt.Println("test run once ...")
}

func goRun() {
	one.Do(test) // 只能接受没有参数的函数， 可以使用闭包来传递参数
	// 只会执行一次test函数
}
func main() {
	for i := 0; i < 10; i++ {
		go goRun()
	}

	time.Sleep(3 * time.Second)
}
