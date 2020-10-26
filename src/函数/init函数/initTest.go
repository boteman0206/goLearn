package main

import "fmt"
import "../demo"

/**
1：init函数实在main函数之前开始执行的
2：init函数的主要作用是完成一些初始化的工作
3： 被引入包的变量 --> 被引入包的init --> 当前全局变量 --> 当前包的init函数 --> mian函数
*/

var num = test()

func test() int {
	fmt.Println("变量初始化 。。。。")
	return 90
}

func init() {
	fmt.Println("init 执行")
}

func main() {

	fmt.Println("main  函数执行")
	fmt.Println(num)
	cal := demo2.Cal(8, 9)
	fmt.Println("cal :", cal)
}
