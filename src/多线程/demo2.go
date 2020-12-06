package main

import "fmt"
import "runtime"

var (
	myMap = make(map[int]int)
)

/**
计算某个n的阶乘 n！
*/
func test1(n int) {
	res := 1
	for i := 1; i < n; i++ {
		res *= i
	}

	myMap[n] = res
}

func main() {

	//runtime的使用
	fmt.Println("cpu数量： ", runtime.NumCPU())
	runtime.GOMAXPROCS(2) // 设置运行cpu数

	//channel管道的使用

	for i := 1; i <= 200; i++ {
		go test1(i) // 开启两百个协程, 使用map会发生并发修改异常
	}

	// 休眠十秒
	time.Sleep(10 * time.Second)

	for e, v := range myMap {
		fmt.Println(e, " v: ", v)
	}

	/**
	todo 检验是否并发异常
	E:\goProject\goLearn\src\多线程>go build -race demo2.go
	E:\goProject\goLearn\src\多线程>demo2.exe
	25  v:  -7835185981329244160
	29  v:  -5968160532966932480
	63  v:  7638104968020361216
	Found 2 data race(s)  -- 发现有并发的异常


	*/
}
