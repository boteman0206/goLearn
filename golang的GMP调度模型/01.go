package main

import (
	"fmt"
	"godemo/src/model"
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

	brand := &model.Brand{}

	brand.Name = "abc"
	brand.Id = 90

	fmt.Println(brand)

}
