package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	go fmt.Println(<-ch1) // todo 这里会产生死锁

	/**
	改成这样即可
	go func() {
		fmt.Println(<- ch1)
	}()

	*/
	ch1 <- 5
	time.Sleep(1 * time.Second)
}
