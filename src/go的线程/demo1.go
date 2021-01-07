package main

import (
	"fmt"
	"sync"
)

var wait sync.WaitGroup

func mu(i int) {
	defer wait.Done() // 计数器减一
	fmt.Println("mu go run ...", i)
}

func main() {

	for i := 0; i < 100; i++ {
		wait.Add(1) // 计数器加一
		go mu(i)
	}

	wait.Wait() // 等待所有的线程结束任务

	fmt.Println("所有的都执行完毕了。。。。")
}
