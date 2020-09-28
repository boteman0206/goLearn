package main

import (
	"fmt"
	"unsafe"
)

/**
	一： 基本数据类型
	1： 数值型  1.1: 整数类型（int，int8，int16, int32,int64
  				uint, uint8, uint16, uint32, uint64, byte）
               1.2： 浮点类型 （float32, float64）

	2： 字符型（没有专门的字符类型，使用byte来保存单个字母字符）
	3： 布尔值（bool类型）
	4： 字符串类型（string 官方分类到基本数据类型）

*/

/**
派生/复杂的数据类型
1： 指针
2： 数组
3： 结构体
4： 管道
5： 函数（也是一种数据类型）
6： 切片
7： 接口
8： map
*/

/**
基本数据类型的默认值： 格式化 %v，按照变量的值来输出
	int 0，
	float 0,
	bool false,
	string ""
*/

func main() {

	/**
	有符号的整数类型
	int8   1字节  -128 ~ 127
	int16  2字节  -2^16  ~ 2^16-1
	int32  4字节  -2^31 ~ 2^31-1
	int64  8字节  -2^64 ~ 2^64-1

	无符号的整数类型
	uint8    1字节	0 ~255  = 2^8-1
	uint16   2字节	0 ~ 2^16-1
	uint32   4字节	0 ~ 2^31-1
	uint64   8字节	0 ~2^64-1

	默认：
	int 有符号  32系统 -2^31 ~ 2^31-1  64位系统 -2^64 ~ 2^64-1
	uint 无符号  -2^31 ~ 2^31-1  64位系统 -2^64 ~ 2^64-1
	rune 有符号  与int32一样    -2^31 ~ 2^32 - 1
	byte 无符号  与uint8等价    0-255  当要存储字符时选用byte
	*/

	var unum1 uint = 31
	fmt.Println(unum1)

	var t1 rune = -1
	fmt.Println("t1 : ", t1)

	//var t2 byte  = -1 // 报错
	var t2 byte = 5
	fmt.Println("t2 :", t2)

	/**
	整形使用细节
	*/
	var n1 = 100                // 默认为int类型
	fmt.Printf("n1 的类型 %T", n1) // 查看某一个数据的类型

	// 查看某个变量的占用细节的大小和数据类型
	var y1 int64 = 9
	fmt.Printf("y1 的数据类型是 %T 占用的字节是 %d", y1, unsafe.Sizeof(y1))

	// 满足开发要求的情况，尽量使用数据类型小的数据
	// bit  1byte=8bit
}
