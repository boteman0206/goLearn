package main

import (
	"fmt"
	"runtime"
)

/*
	设置可同时执行的逻辑Cpu数量，默认和硬件的线程数一致而不是核心数，可以通过调用GOMAXPROCS(-1)来获取当前逻辑Cpu数
	最好在main函数之前设置它，GOMAXPROCS同时也是go的环境变量之一
	 当n<1：非法数字，方法不作修改
	 当n==1：单核心，多协程并发执行，并发只是看起来是同时执行的，实际上是同一时刻只有一个协程在跑，只是由于cpu的任务调度算法，让多个协程在效果上同时执行

	 当n>1：多核心，多协程并行执行，并行一定是并发，不同的核心同时地跑不同的协程

*/
func main() {

	gomaxprocs := runtime.GOMAXPROCS(-1)
	fmt.Println("逻辑Cpu数： ", gomaxprocs)

	cpu := runtime.NumCPU()
	fmt.Println(" 逻辑cpu数： ", cpu)

	//n := runtime.GOMAXPROCS(4) //指定以4核运算
	//fmt.Println("n = ", n)

	go task()
	go task()
	go task()
	go task()

	select {}
}

func task() {
	for {

	}
}
