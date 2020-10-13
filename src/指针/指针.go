package main

import "fmt"

/**
	1：基本数据类型，变量存储的就是指，也叫值类型
	2：获取变量的地址用 &符号
    3: 通过指针变量获取值 *ptr
*/

func main() {

	var i int = 10
	fmt.Println("i的内存内存地址 ： ", &i) // i的内存内存地址 ：  0xc00000a0b0

	var ptr *int = &i            // 将变量i的地址赋值给ptr指针变量
	fmt.Println("ptr变量 ： ", ptr) // 0xc00000a0b0
	//获取值
	fmt.Println("获取ptr存储的值：", *ptr) // 获取ptr存储的值： 10

	//指针也有地址
	fmt.Println("ptr指针的地址：", &ptr) // 0xc000006030

	//案例：
	var num int = 9
	var pptr *int = &num
	fmt.Println("修改之前：", num)
	*pptr = 100
	fmt.Println("修改之后： ", num)

}
