package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	time.Sleep(10 * time.Second)
	since := time.Since(now)
	fmt.Println(since)

}
