// 展示内存增长和pprof，并不是泄露

// 上面这个demo会不断的申请内存，把它编译运行起来，然后执行：

package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
)

// 运行一段时间：fatal error: runtime: out of memory
func main() {
	// 开启pprof
	go func() {
		ip := "0.0.0.0:6060"
		if err := http.ListenAndServe(ip, nil); err != nil {
			fmt.Printf("start pprof failed on %s\n", ip)
			os.Exit(1)
		}
	}()

	go doWhile()

	_ = http.ListenAndServe("0.0.0.0:8880", nil)
}

func doWhile() {
	loop := 10000000000
	for i := 0; i < loop; i++ {
		fmt.Println("----", i)
	}
}

/**
https://graphviz.org/download/


1: go tool pprof http://localhost:6060/debug/pprof/heap

2: top命令 按指标大小列出前10个函数，比如内存是按内存占用多少，CPU是按执行时间多少。
top会列出5个统计数据：
flat: 本函数占用的内存量。
flat%: 本函数内存占使用中内存总量的百分比。
sum%: 前面每一行flat百分比的和，比如第2行虽然的100% 是 100% + 0%。
cum: 是累计量，加入main函数调用了函数f，函数f占用的内存量，也会记进来。
cum%: 是累计量占总量的百分比。
3：list 查看某个函数的代码，以及该函数每行代码的指标信息，如果函数名不明确，会进行模糊匹配，比如list main会列出main.main和runtime.main。

4： traces  打印所有调用栈，以及调用栈的指标信息。



需要安装 brew install graphviz

todo go tool pprof -http=:8000 http://127.0.0.1:6060/debug/pprof/profile  // 可以查看分析内存
*/
