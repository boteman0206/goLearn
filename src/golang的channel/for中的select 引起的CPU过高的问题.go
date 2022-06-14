package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	quit := make(chan bool)
	for i := 0; i != runtime.NumCPU(); i++ {
		go func() {
			for {
				select {
				case <-quit:
					fmt.Println("获取到quit")
					break // todo 换成return可以解决
				default:
				}
			}
		}()
	}

	time.Sleep(time.Second * 5)
	for i := 0; i != runtime.NumCPU(); i++ {
		quit <- true
	}
	time.Sleep(30 * time.Second)
}

/**
上面这段代码会把所有CPU都跑满，原因就就在select的用法上。
一般来说，我们用select监听各个case的IO事件，每个case都是阻塞的。上面的例子中，我们希望select在获取到quit通道里面的数据时立即退出循环，
但由于他在for{}里面，在第一次读取quit后，仅仅退出了select{}，并未退出for，所以下次还会继续执行select{}逻辑，
此时永远是执行default，直到quit通道里读到数据，否则会一直在一个死循环中运行，即使放到一个goroutine里运行，也是会占满所有的CPU。

解决方法就是把default去掉即可，这样select就会一直阻塞在quit通道的IO上， 当quit有数据时，就能够随时响应通道中的信息。

*/
