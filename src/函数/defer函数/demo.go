package main

import "fmt"

/**
1：defer函数回将语句放入栈中，也会将相关的值同时拷贝
2：用来延时执行，函数执行之后再执行相关的defer栈
*/

func add(n1, n2 int) int {
	defer fmt.Println("defer1 n1 : ", n1)
	defer fmt.Println("defer2 n2 : ", n2)
	n1++
	n2++
	fmt.Println("n1 : ", n1, " n2 : ", n2)
	return n1 + n2
}

func main() {

	fmt.Println("defer 函数")

	res := add(10, 100)
	fmt.Println("res : ", res)
}
