package main

import "fmt"

var (
	func1 = func(n1, n2 int) int {
		return n1 * n2
	}
)

func main() {

	res := func(n1 int) int {
		fmt.Println("n1 ", n1)
		return n1
	}(10)

	fmt.Println(res)

	a := func(n1, n2 int) int {
		return n1 + n2
	}

	i := a(2, 4)
	fmt.Println("i = ", i)

	//全局匿名函数
	i2 := func1(6, 5)
	fmt.Println("全局匿名函数 ： ", i2)

}
