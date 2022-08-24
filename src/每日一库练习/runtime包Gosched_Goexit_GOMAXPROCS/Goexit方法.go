package main

import (
	"fmt"
	"runtime"
	"time"
)

/**

立即终止当前协程，不会影响其它协程，且终止前会调用此协程声明的defer方法。由于Goexit不是panic，所以recover捕获的error会为nil
当main方法所在主协程调用Goexit时，Goexit不会return，所以主协程将继续等待子协程执行，当所有子协程执行完时，程序报错deadlock

*/

func main() {
	go func() {
		defer func() {
			fmt.Println("defer func executed!")
			fmt.Println("recovered error == ", recover())
		}()

		for i := 0; i < 3; i++ {
			if i == 1 {
				runtime.Goexit()
			}

			fmt.Println(i)
		}
	}()

	// main里面使用此方法，不会return，会继续等待子协程执行完 ，然后程序直接退出,引发panic, 报错deadlock
	runtime.Goexit()
	time.Sleep(10 * time.Second)
}
