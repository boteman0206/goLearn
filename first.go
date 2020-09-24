package main

import (
	"fmt" // 导入 fmt 包，打印字符串是需要用到
)

func a() {
	fmt.Println("this is the Hello World!") // 打印 Hello World!

}

func main() { // 声明 main 主函数
	fmt.Println("a")
	fmt.Println("bc") // todo 同一行需要加分号 否则报错
	a()
}

//go run  path    直接运行， 不会生成可执行文件    速度慢

// go build path   会进行编译成  .exe的可执行文件  速度快

// go build进行编译的时候回直接打包运行环境， 在没有sdk的环境下面依然可以直接执行
