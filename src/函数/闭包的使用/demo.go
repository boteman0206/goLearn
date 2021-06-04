package main

/**

 */

import "fmt"

func main() {

	one := nextOne()
	fmt.Println(one())
	fmt.Println(one())
	fmt.Println(one())

}

/**
返回的是匿名函数，但是匿名函数引用到函数外部的i1，因此这个匿名函数就和i1
形成一个整体，就是闭包。i1并不会重新初始化
*/
func nextOne() func() int {
	i1 := 3
	return func() int {
		i1 += 1
		return i1
	}
}
