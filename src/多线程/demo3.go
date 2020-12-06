package main

import (
	"sync"
	"time"
)
import "fmt"
import "runtime"

/**
在demo2的基础上使用全局变量myMap加锁限制并发修改异常
sync加锁，同步基本单元，低水平的使用

*/

var (
	myMap1 = make(map[int]int)
	lock   sync.Mutex
)

/**
计算某个n的阶乘 n！
*/
func test2(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	lock.Lock() // 加锁
	myMap1[n] = res
	lock.Unlock() // 解锁
}

func main() {

	//runtime的使用
	fmt.Println("cpu数量： ", runtime.NumCPU())
	runtime.GOMAXPROCS(2) // 设置运行cpu数

	//channel管道的使用

	for i := 1; i <= 20; i++ {
		go test2(i) // 开启两百个协程, 使用map会发生并发修改异常
	}

	// 休眠十秒
	time.Sleep(10 * time.Second)

	lock.Lock()
	for e, v := range myMap1 {
		fmt.Println(e, " v: ", v)
	}
	lock.Unlock()

}
