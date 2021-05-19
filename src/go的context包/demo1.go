package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func test1(ctx context.Context) {

	time.Sleep(3 * time.Second)
	fmt.Println("this is ")
	fmt.Println("获取到value： ", ctx.Value("name"))
	wait.Done()
}

var wait *sync.WaitGroup

func main() {

	fmt.Println("main run ...")
	wait = new(sync.WaitGroup)
	wait.Add(1)
	background := context.WithValue(context.Background(), "name", "jack-data")
	go test1(background)

	wait.Wait()

	fmt.Println("测试 run ")

}
