package main

import "fmt"

// 全局变量和局部变量的定义类似
// 方式1:
var x1, x2 = 90, "全局变量"

// 方拾2：
var (
	y1 = 78
	y2 = "pop"
)

func main() {

	/**
	一： 单个变量的基本使用
	*/
	//1: 基本使用
	var i int
	i = 10
	fmt.Println("变量的测试", i)

	//2: 默认值
	var num int
	fmt.Println("int的默认值 ： ", num)

	// 3: 类型推到方式
	var num1 = 10
	fmt.Println("默认推到类型 ： ", num1)

	// 4: name:='tom' 冒号不能省略 并且name之前没有申明过，否则报错
	name := "tom"
	fmt.Println("name : ", name)

	/**
	二： 多个变量的申明
	*/

	// 1 : 相同的类型
	var n1, n2, n3 int
	fmt.Println("多变量： ", n1, n2, n3)

	// 2: 不同的类型
	var m1, m2, m3 = 100, "tom", 103.3
	fmt.Println("不同类型的多变量：", m1, m2, m3)

	// 类型推到
	t1, t2, t3 := 90, 87, "jack"
	fmt.Println("类型推导形式 ： ", t1, t2, t3)

	var x1 = 829 // todo 可以覆盖全局变量不会报错
	//var t1 = 99  // todo 覆盖局部变量报错 只能修改
	//全局变量的输出
	fmt.Println("x1, x2 ", x1, x2)
	fmt.Println("y1, y2 ", y1, y2)

	/**
	三： 变量的使用细节
		1: 可以在同一个类型范围内改变数值， 即重新赋值但不能改变数据的类型
		2：变量在同一个作用域内不能重名， 在一个函数中或者代码块之内
		3：int，float 默认值是 0， string 默认是 空串
	*/

	var a float32
	fmt.Println("小数的默认值 ： ", a)

	var name1 string
	fmt.Println("string初始值： ", name1)

}
