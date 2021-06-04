package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		s := v
		go func() {
			time.Sleep(3 * time.Second)
			fmt.Println(s)
			done <- true
		}()
	}

	// wait for all goroutines to complete before exiting
	for _ = range values {
		<-done
	}
}
