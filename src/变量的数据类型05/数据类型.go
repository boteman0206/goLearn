package main

import "fmt"

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

func main() {

	/**
	有符号的整数类型
	int8   1字节  -128 ~ 127
	int16  2字节  -2^16  ~ 2^16-1
	int32  3字节  -2^31 ~ 2^31-1
	int64  4字节  -2^64 ~ 2^64-1

	无符号的整数类型
	uint8    1字节	0 ~255  = 2^8-1
	uint16   2字节	0 ~ 2^16-1
	uint32   3字节	0 ~ 2^32-1
	uint64   4字节	0 ~2^64-1

	默认：
	int 有符号  32系统
	uint 无符号  32

	*/

	var unum1 uint = 31
	fmt.Println(unum1)

}
