package main

import "fmt"

/**
当通过通道发送有限的数据时，我们可以通过close函数关闭通道来告知从该通道接收值的goroutine停止等待。
当通道被关闭时，往该通道发送值会引发panic，从该通道里接收的值一直都是类型零值。那如何判断一个通道是否被关闭了呢？

*/
// channel 练习
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1 //todo 通道关闭后再取值ok=false
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()

	// 在主goroutine中从ch2中接收值打印
	//todo 在遍历时，如果channel 没有关闭，那么会一直等待下去，出现 deadlock 的错误；如果在遍历时channel已经关闭，那么在遍历完数据后自动退出遍历。也就是说，for range 的遍历方式时阻塞型的遍历方式。
	for i := range ch2 { // 通道关闭后会退出for range循环
		fmt.Println(i)
	}
}
