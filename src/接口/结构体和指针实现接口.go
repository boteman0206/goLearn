package main

import "fmt"

type Cat struct{}
type Duck interface {
	Quack()
}

//func (c  Cat) Quack(){}  // 使用结构体实现接口
func (c *Cat) Quack() {} // 使用结构体指针实现接口

//var d Duck = Cat{}      // 使用结构体初始化变量
var d Duck = &Cat{} // 使用结构体指针初始化变量

func NilOrNot(v interface{}) bool {

	return v == nil
}

/**

https://draveness.me/golang/docs/part2-foundation/ch04-basic/golang-interface/#nil-%E5%92%8C-non-nil
使用 runtime.iface 结构体表示包含方法的接口
使用 runtime.eface 结构体表示不包含任何方法的 interface{} 类型；

1:空接口eface的结构
	type eface struct { // 16 字节
		_type *_type
		data  unsafe.Pointer}
2:iface的结构
	type iface struct { // 16 字节
		tab  *itab
		data unsafe.Pointer
	}

*/

func main() {

	var s *Cat
	fmt.Println(s == nil)    // ==> true 定义一个变量初始化为nil
	fmt.Println(NilOrNot(s)) // false 将s传给interface变量，此时不为nil
	//todo 调用 NilOrNot 函数时发生了隐式的类型转换 在类型转换时，*TestStruct 类型会转换成 interface{} 类型，转换后的变量不仅包含转换前的变量，
	// 还包含变量的类型信息 TestStruct，所以转换后的变量与 nil 不相等。

	// 从源码里可以看到：iface包含两个字段：
	// 	1: tab 是接口表指针，指向类型信息；
	// 	2: data 是数据指针，则指向具体的数据。它们分别被称为动态类型和动态值。而接口值包括动态类型和动态值。
	// 接口值的零值是指动态类型和动态值都为 nil

	var _ error = (*MyError)(nil) // 检查是否时间了error接口

	//var _ error = MyError{}
}

type MyError struct {
}

func (s *MyError) Error() string {
	return "my error"
}

//
//func (s MyError) Error() string {
//	return "my error"
//}
