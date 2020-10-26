package main

import (
	"fmt"
)
import "../demo"     // 引用方式一： 包名demo和文件包的名demo2不一致
import "../test"     // 引用方式二： 包名test和文件包名test一致
import abc "../test" // 别名的方式导入，使用时候用abc去调用

/**
包的快速入门
1： 包名和文件夹的名称通常保持一致
2： 使用首字母大写的函数或者变量才可以导出， 类似于java的public关键字
3： import包时， 路径是从$GOPATH的src下面开始的，
	这个需要配置，可以使用相对路径来解决
*/
func main() {

	cal := demo2.Cal(16, 18) // 使用方式： 包名和文件名不一致
	fmt.Println("cal : ", cal)

	s := test.Test("hello ,", "world!")
	fmt.Println("s : ", s) // 使用方式： 包名和文件名一致

	s1 := abc.Test("hello ,", "world!")
	fmt.Println("别名: ", s1)

	// cd到main包的文件目录下
	//go build -o my.exe 变编译成二进制文件 取别名
	//go build ./main.go 变编译成二进制文件 名称为main.exe

	//返回多个值得在计算方式, 可以使用_下划线来站位进行忽略
	i, i2 := demo2.MultiNum(30, 4)
	fmt.Println("第一个： ", i, " 第二个值： ", i2)

	digui1(4)
	fmt.Println("============")
	digui2(4)
	fmt.Println("fib========")
	fmt.Println(fib(5))

	type myInt int // 定义类型
	var num1 myInt
	num1 = 900
	fmt.Println(num1)

	sum, sub := getSumAndSub(12, 99)
	fmt.Println("sum ： ", sum, " sub : ", sub)

	tt := sumTT(1, 2, 32, 3, 4)
	fmt.Println("TT : ", tt)
}

func digui1(n int) {
	if n > 2 {
		n--
		digui1(n)
	}
	fmt.Println("n = ", n)
}

func digui2(n int) {
	if n > 2 {
		n--
		digui2(n)
	} else {
		fmt.Println("n = ", n)
	}

}

//斐波那契数列
func fib(n int) int {
	// 1,1,2,3,5,8
	if n == 2 || n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

/**
函数注意事项
1： 基本数据类型和数组都是值传递
2: 如果希望函数内的变量能够修改函数外的变量，可以使用指针的方式来操作
3： 没有函数的重载
4： 函数也是数据类型， 可以赋值给一个变量，
5： int取别名 type myInt int;
6： 支持对函数返回值的命名
7: go语言支持可变参数
*/

func getSumAndSub(n1, n2 int) (sum, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return // 直接通过名称返回 (sum, sub int)
}

//可变参数
func sumTT(n1 int, args ...int) int {
	var sum1 int
	for e := range args {
		fmt.Println("e: ", args[e])
		sum1 += args[e]
	}
	sum1 += n1
	return sum1
}
