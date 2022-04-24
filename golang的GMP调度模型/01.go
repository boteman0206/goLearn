package main

import (
	"fmt"
	"runtime"
)

/**
   // 参考文章
	https://zhuanlan.zhihu.com/p/323271088

	// 抢占时调度和协作式调度
	https://mp.weixin.qq.com/s/2ibzE46UVnO_YNtugikpPQ
*/

func main() {

	runtime.GOMAXPROCS(1)

	fmt.Println(runtime.NumCPU())

	//SetMaxThreads()

}
