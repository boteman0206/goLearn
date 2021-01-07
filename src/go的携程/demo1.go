package main

import (
	"fmt"
	"sync"
)

var wait sync.WaitGroup

func mu(i int) {
	defer wait.Done()
	fmt.Println("mu go run ...", i)
}

func main() {

	for i := 0; i < 100; i++ {
		wait.Add(1)
		go mu(i)
	}

	wait.Wait()

	fmt.Println("所有的都执行完毕了。。。。")
}
