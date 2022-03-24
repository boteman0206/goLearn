package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mapT = sync.Map{}
)

func main() {

	dt := make(chan int, 100)

	//for i:=0; i<99; i++ {
	//
	//	dt <- i
	//}

	close(dt)
	fmt.Println("----------")
	for value := range dt {
		fmt.Println(value)
		fmt.Println("================")
		if len(dt) <= 0 {
			fmt.Println("break")
			break
		}
	}

	fmt.Println("ooooooooooooo")
	time.Sleep(1 * time.Minute)
	fmt.Println("jie   shu")
}
