package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
1：数组是内存连续的
2：数组是值类型， 切片是引用类型
*/

func main() {
	// 四种初始化数组的方式
	var num1 [3]int = [3]int{1, 2, 3}
	var num2 = [3]int{1, 2, 4}
	var num3 = [...]int{1, 2, 3, 4}
	// 指定下标的方式
	var num4 = [...]int{1: 90, 0: 100, 2: 67, 6: 901}
	num5 := [...]int{123, 34, 56}

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)
	fmt.Println(num5)

	//1: 遍历方式常规的for循环
	//2: for - range遍历
	for index, value := range num1 {
		fmt.Println("index :", index, "value : ", value)
	}

	for _, value := range num1 {
		fmt.Println("value : ", value)
	}

	// 使用细节和注意事项
	/**
	1:数组的长度，类型是固定的不能动态的变化
	2：var num []int，数切片不是数组
	3：可以:是基本数据类型也可以是之类类型
	4：数组有默认值，就是（零值）
	5：数组是值类型，默认情况下是值拷贝，修改数组不会影响原来的数组的值
	6： 如果想改变原来的数组的值没使用引用传递的方式（指针的方式）
	*/

	var intArr [5]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		intArr[i] = rand.Intn(100)
	}
	fmt.Println(intArr)
}
