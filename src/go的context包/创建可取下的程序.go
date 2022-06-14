package main

import (
	"fmt"
	"golang.org/x/net/context"
	"time"
)

func main() {
	//创建一个可取消子context,context.Background():返回一个空的Context，这个空的Context一般用于整个Context树的根节点。
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			//使用select调用<-ctx.Done()判断是否要结束
			case <-ctx.Done():
				fmt.Println("goroutine exit")
				return
			default:
				fmt.Println("goroutine running.")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(10 * time.Second)
	fmt.Println("main fun exit")
	//取消context
	cancel()
	time.Sleep(5 * time.Second)
}
