package main

import (
	"fmt"
	"golang.org/x/net/context"
	"time"
)

func main() {
	d := time.Now().Add(20 * time.Second)
	//设置超时控制WithDeadline，超时时间2
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	for true {
		select {
		case <-time.After(30 * time.Second):
			fmt.Println("timeout")
		case <-ctx.Done():
			//2到了到了，执行该代码
			fmt.Println(ctx.Err())
			return
		}

	}
}
