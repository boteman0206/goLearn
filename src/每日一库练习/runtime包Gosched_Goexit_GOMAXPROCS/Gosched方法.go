package main

/**
Gosched
暂停当前goroutine，使其他goroutine先行运算。只是暂停，不是挂起，当时间片轮转到该协程时，Gosched()后面的操作将自动恢复
*/

import (
	"fmt"
	"runtime"
)

func main() {
	go output("goroutine 2")

	/**
	**结论：**在打印goroutine 1之前，主协程调用了runtime.Gosched()方法，暂停了主协程。子协程获得了调度，从而先行打印了goroutine 2。
		主协程不是一定要等其他协程执行完才会继续执行，而是一定时间。如果这个时间内其他协程没有执行完，那么主协程将继续执行，
	*/
	runtime.Gosched()
	output("goroutine 1")

}

func output(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
	}
}
