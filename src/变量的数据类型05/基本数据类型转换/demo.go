package main

import (
	"fmt"
	"strconv"
)

/**
基本数据类型的相互转换
1： 只能显示转换， 不能显式的转换
2: 大的类行转成小的类型会溢出，需要注意
3： 被转换的是值，变量数据的类型并不会变化

*/
func main() {
	/**
	数字类型的转换
	*/
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
	n4 = b + 30 // todo 正确  ?? 为啥这个没错呢 默认是int啊

	var name1 = 30
	fmt.Printf("name1 %T \n", name1)
	fmt.Println(n3)
	fmt.Println(n4)

	/**
	基本类型转string数据
	*/
	//方式一  fmt.Sprintf
	var h1 bool
	var h2 byte = 'h'
	var s1 string
	s1 = fmt.Sprintf("%v  \n", n4)
	s1 = fmt.Sprintf("%v \n", a)
	s1 = fmt.Sprintf("%t \n", h1) // bool输出转换
	s1 = fmt.Sprintf("%c \n", h2) // 字符输出转换
	fmt.Println("s1  ", s1)

	//方式二 strconv包下面的格式
	formatBool := strconv.FormatBool(h1)
	formatInt := strconv.FormatInt(int64(a), 10) // 后面的10代表十进制
	formatInt2 := strconv.FormatInt(int64(a), 2) // 后面的10代表十进制
	fmt.Println(formatBool)
	fmt.Println(formatInt)
	fmt.Println(formatInt2)
	itoa := strconv.Itoa(988)
	fmt.Println("直接转换int ", itoa)

	//string数据类型的转基本类型
	parseBool, e := strconv.ParseBool("true")
	parseBool1, e := strconv.ParseBool("1") // 转换的时候居然能识别 1 0
	fmt.Println("parsebool = ", parseBool, "e = ", e)
	fmt.Println("parsebool1 = ", parseBool1, "e = ", e)
	i2, err := strconv.ParseInt("131", 10, 64) // bitsize 0 8 16 32 64
	fmt.Println(i2, err)
	if err == nil {
		fmt.Println("err is ", err)
	}

}
