package main

import (
	"fmt"
	"strings"
	"time"
)

/**
// https://mp.weixin.qq.com/s/ZXYpfLNGyej0df2zXqfnHQ
todo 当前 channel 不存在缓冲区，也就是元素大小为 0 的情况下，就会调用 mallocgc 方法分配一段连续的内存空间。
	当前 channel 存储的类型存在指针引用，就会连同 hchan 和底层数组同时分配一段连续的内存空间。
	通用情况，默认分配相匹配的连续内存空间。
	需要注意到一块特殊点，那就是 channel 的创建都是调用的 mallocgc 方法，也就是 channel 都是创建在堆上的。因此 channel 是会被 GC 回收的，自然也不总是需要 close 方法来进行显示关闭了。
*/

func testChan(chan1 chan int) {
	for true {
		time.Sleep(1 * time.Second)
		fmt.Println("zu se zhu le 0")
		chan1 <- 12
		fmt.Println("zu se zhu le 1")
	}

	for true {
		time.Sleep(3 * time.Second)
		fmt.Println(" test chan run ")
	}

	//time.Sleep(1*time.Second)
	//fmt.Println("zu se zhu le 0")
	//chan1 <- 12
	//fmt.Println("zu se zhu le 1")
	//chan1 <- 13
}

func main() {

	//
	chan1 := make(chan int)
	go testChan(chan1)

	//
	//for {
	//	time.Sleep(10)
	//}

	//for true {
	//	select {
	//	case num := <- chan1:
	//		fmt.Println("----", num)
	//	default:
	//		//fmt.Println("default ...")
	//	}
	//}

	//select {
	//case num := <- chan1:
	//	fmt.Println("----", num)
	//default:
	//	//fmt.Println("default ...")
	//}

	//select {
	//
	//}

	code := " c "
	i := len(strings.TrimSpace(code))
	fmt.Println(i)

}

//
//func main(){
//	for i := 0; i < 20; i++ { //启动20个协程处理消息队列中的消息
//		go thrind(i)
//	}
//	select {} // 阻塞
//}
//func thrind( i int){
//	for range time.Tick(1000 * time.Millisecond) {
//		fmt.Println("\n 线程：",i)
//	}
//}
