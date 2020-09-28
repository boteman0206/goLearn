package main

import "fmt"

/**
单精度float32  4字节  尾数部分可能丢失，造成精度损失

双精度float63  8字节  精度大一点

golang的默认浮点类型为float64
*/

func main() {
	fmt.Println("小数使用")

	var price float32 = 0.12
	fmt.Println("price : ", price)

	var a = 1.2
	fmt.Printf("默认数据类型 %T", a)

	//科学计数法
	num0 := 5.1212e2 //  * 10的2次方
	fmt.Println("num0 = ", num0)
}
