package main

import (
	"fmt"
	"reflect"
)

/**
切片相当于动态的数组，切片是值类型
切片可以动态的变化长度
切片定义 var a []int
*/

func main() {

	var intArr = [5]int{12, 34, 6, 7, 8}

	//申明切片
	//方式一
	slice := intArr[1:3]
	fmt.Println("slice : ", slice, " len ", len(slice), " cap: ", cap(slice))
	slice[1] = 9000
	fmt.Println("修改slice之后 : ", intArr)

	// 方式二
	var t1 []int
	//t1[0] = 9  // 不能直接使用需要make初始化
	//fmt.Println(t1)
	t1 = make([]int, 3, 6)
	fmt.Println(t1, t1[0])
	// 方式三
	var num = []int{12, 2334, 56, 7}
	fmt.Println(reflect.TypeOf(num), num)

	//切片的遍历
	for index, value := range num {
		fmt.Println("index :", index, " value : ", value)
	}

	ints := append(num, 122)
	fmt.Println("ints :", ints)

	i := append(num, num...) // 自增加数组
	fmt.Println("num 自增加 ：", i)

	i2 := copy(num, slice)
	fmt.Println(i2)
	fmt.Println("num : ", num)

}
