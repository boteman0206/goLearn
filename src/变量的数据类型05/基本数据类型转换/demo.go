package main

import (
	"fmt"
)

/**
基本数据类型的相互转换
1： 只能显示转换， 不能显式的转换
2: 大的类行转成小的类型会溢出，需要注意
3： 被转换的是值，变量数据的类型并不会变化

*/
func main() {

	var a int = 988
	fmt.Printf("a类型 %T, a = %v \n", a, a)

	f := float32(a)
	fmt.Printf("f类型 %T, f = %v \n", f, f)

	var b float64 = 0.1111
	fmt.Printf("b类型 %T, b = %v \n", b, b)

	i := float32(b)
	fmt.Printf("i类型 %T, i = %v \n", i, i)

	var n3 float64
	var n4 float64

	//n3 = b + a  // 不同的类型不能直接相加，必须显示的转换
	n4 = b + 30 // 正确  ?? 为啥这个没错呢

	var name1 = 30
	fmt.Printf("name1 %T \n", name1)
	fmt.Println(n3)
	fmt.Println(n4)

}
