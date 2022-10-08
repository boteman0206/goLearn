package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	go func() {
		Test01(ctx)
	}()

	fmt.Println("main end ....")

	time.Sleep(10 * time.Second)
}

func Test01(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("============", err)
		}
	}()

	fmt.Println("test01 start.... ")
	select {
	case <-ctx.Done():
		fmt.Println(" 已经关闭了。。。。。ctx ")
	}

	time.Sleep(5 * time.Second)
	fmt.Println("test01 end.... ")
}
