package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func(c chan int) {
		// 修改时间后,再查看执行结果
		time.Sleep(time.Second * 1)
		ch <- 1
	}(ch)

	select {
	case v := <-ch:
		fmt.Print(v)
	case <-time.After(2 * time.Second): // 等待 2s
		fmt.Println("no case ok")
	}

	time.Sleep(time.Second * 10)
}

//我们通过修改上面的时等待时间可以看到，如果等待时间超出<2秒，则输出1，否则打印“no case ok”
