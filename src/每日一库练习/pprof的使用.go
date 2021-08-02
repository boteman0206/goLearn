package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

/**
 添加  _ "net/http/pprof"
 使用  http://127.0.0.1:6060/debug/pprof/ 访问
	allocs：查看过去所有内存分配的样本，访问路径为 $HOST/debug/pprof/allocs。
	block：查看导致阻塞同步的堆栈跟踪，访问路径为 $HOST/debug/pprof/block。
	cmdline： 当前程序的命令行的完整调用路径。
	goroutine：查看当前所有运行的 goroutines 堆栈跟踪，访问路径为 $HOST/debug/pprof/goroutine。
	heap：查看活动对象的内存分配情况， 访问路径为 $HOST/debug/pprof/heap。
	mutex：查看导致互斥锁的竞争持有者的堆栈跟踪，访问路径为 $HOST/debug/pprof/mutex。
	profile： 默认进行 30s 的 CPU Profiling，得到一个分析用的 profile 文件，访问路径为 $HOST/debug/pprof/profile。
	threadcreate：查看创建新 OS 线程的堆栈跟踪，访问路径为 $HOST/debug/pprof/threadcreate。
	如果你在对应的访问路径上新增 ?debug=1 的话，就可以直接在浏览器访问，


 3：使用终端交互 go tool pprof http://localhost:6060/debug/pprof/profile?seconds=60 执行该命令后，需等待 60 秒（可调整 seconds 的值），pprof 会进行 CPU Profiling，结束后将默认进入 pprof 的命令行交互式模式，
	flat：函数自身的运行耗时。
	flat%：函数自身在 CPU 运行耗时总比例。
	sum%：函数自身累积使用 CPU 总比例。
	cum：函数自身及其调用函数的运行总耗时。
	cum%：函数自身及其调用函数的运行耗时总比例。
	Name：函数名。
*/

var datas []string

func main() {
	go func() {
		for {
			log.Printf("len: %d", Add("go-programming-tour-book"))
			time.Sleep(time.Millisecond * 10)
		}
	}()

	_ = http.ListenAndServe("0.0.0.0:6060", nil)
}

func Add(str string) int {
	data := []byte(str)
	datas = append(datas, string(data))
	return len(datas)
}
