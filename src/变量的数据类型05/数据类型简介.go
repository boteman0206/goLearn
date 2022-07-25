package main

import (
	"fmt"
	"unsafe"
)

/**
1: 整形  分类无符号uint 和 有符号的int
	1.1：int8和uint8：
		uint8 是无符号8 代表 8个bit，能表示的数值个数有 2^8 = 256
		int8 是有符号，既可以正数，也可以负数，那怎么办？对半分呗，-128-127，也刚好 256个数。
	1.2： int16/uint16  int32/uint32 int64/uint64 同理

	1.3： 而int 没有并没有指定它的位数
		当你在32位的系统下，int 和 uint 都占用 4个字节，也就是32位。
		若你在64位的系统下，int 和 uint 都占用 8个字节，也就是64位。


2: 浮点型
	2.1: float32 :也即我们常说的单精度，存储占用4个字节，也即4*8=32位，其中1位用来符号，8位用来指数，剩下的23位表示尾数
	2.2: float64 :也即我们熟悉的双精度，存储占用8个字节，也即8*8=64位，其中1位用来符号，11位用来指数，剩下的52位表示尾数

	2.3: 那么精度是什么意思？有效位有多少位？
		对于 float32（单精度）来说，表示尾数的为23位，除去全部为0的情况以外，最小为2-23，约等于1.19*10-7，所以float小数部分只能精确到后面6位，加上小数点前的一位，即有效数字为7位。
		同理 float64（单精度）的尾数部分为 52位，最小为2-52，约为2.22*10-16，所以精确到小数点后15位，加上小数点前的一位，有效位数为16位。
	2.4: float32 和 float64 可以表示的数值很多
		常量 math.MaxFloat32 表示 float32 能取到的最大数值，大约是 3.4e38；
		常量 math.MaxFloat64 表示 float64 能取到的最大数值，大约是 1.8e308；
		float32 和 float64 能表示的最小值分别为 1.4e-45 和 4.9e-324。

3: 空结构体
	3.1：优点
		不占用内存
		地址都一样（所有的空结构变量都指向全局的同一个地址）
	3.2：作用：
		map[string]struct{} 表示key是否存在，其值不占用任何内存。Go语言没有集合类型，用这种结构就可以模拟了。
		chan struct{} 用于通道模拟信号

*/

func main() {

	// 二进制  以0b或0B为前缀
	var num01 int = 0b1100
	fmt.Println("二进制: ", num01)
	//八进制 以0o或者 0O为前缀
	var num02 int = 0o14
	fmt.Println("八进制：", num02)

	//16进制：以0x 为前缀
	var num03 int = 0xC
	fmt.Println("16进制：", num03)

	var str string = "acdfff"
	fmt.Printf("str的类型 %T 占用字节数%d \n", str, unsafe.Sizeof(str))
	var num32 int32 = 12
	fmt.Printf("num32的类型 %T 占用字节数%d \n", num32, unsafe.Sizeof(num32))

	s1 := struct{}{}
	s2 := struct{}{}
	s3 := struct{}{}
	fmt.Printf("s1的类型 %T 占用字节数%d  地址 %p\n", s1, unsafe.Sizeof(s1), &s1)
	fmt.Printf("s2的类型 %T 占用字节数%d  地址 %p\n", s2, unsafe.Sizeof(s2), &s2)
	fmt.Printf("s3的类型 %T 占用字节数%d  地址 %p \n", s3, unsafe.Sizeof(s3), &s3)

}
