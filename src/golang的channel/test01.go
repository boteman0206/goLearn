package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan struct{})

	go func() {
		for {
			select {
			case st, ok := <-c:
				fmt.Println(" close chan  ", st, " ok : ", ok) // 这里会阻塞 关闭之后不会阻塞 ok=false
				if !ok {
					return
				}
			default:
				fmt.Println(" 默认---")

			}
		}
	}()

	time.Sleep(2 * time.Nanosecond)

	c <- struct{}{}

	time.Sleep(2 * time.Nanosecond)
	close(c)
	time.Sleep(2 * time.Second)
}
