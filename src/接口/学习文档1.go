package main

import "fmt"

/**


文档： https://golang.design/go-questions/interface/iface-eface/

iface 和 eface 都是 Go 中描述接口的底层结构体，区别在于 iface 描述的接口包含方法，而 eface 则是不包含任何方法的空接口：interface{}。


iface源码：
type iface struct {
	tab  *itab
	data unsafe.Pointer
}

type itab struct {
	inter  *interfacetype
	_type  *_type
	link   *itab
	hash   uint32 // copy of _type.hash. Used for type switches.
	bad    bool   // type does not implement interface
	inhash bool   // has this itab been added to hash?
	unused [2]byte
	fun    [1]uintptr // variable sized
}

一个table是40字节大小
type itab struct {
	inter  *interfacetype // 8字节
	_type  *_type // 8字节
	link   *itab // 8字节
	hash   uint32 // 4字节
	bad    bool   // 1字节
	inhash bool   // 1字节
	unused [2]byte // 2字节
	fun    [1]uintptr // variable sized // 8字节
}

iface 内部维护两个指针，tab 指向一个 itab 实体， 它表示接口的类型以及赋给这个接口的实体类型。data 则指向接口具体的值，一般而言是一个指向堆内存的指针。




eface源码
type eface struct {
    _type *_type
    data  unsafe.Pointer
}


todo _type结构体
type _type struct {
    // 类型大小
	size       uintptr
    ptrdata    uintptr
    // 类型的 hash 值
    hash       uint32
    // 类型的 flag，和反射相关
    tflag      tflag
    // 内存对齐相关
    align      uint8
    fieldalign uint8
    // 类型的编号，有bool, slice, struct 等等等等
	kind       uint8
	alg        *typeAlg
	// gc 相关
	gcdata    *byte
	str       nameOff
	ptrToThis typeOff
}







从源码里可以看到：iface包含两个字段：tab 是接口表指针，指向类型信息；data 是数据指针，则指向具体的数据。它们分别被称为动态类型和动态值。



*/

/**
todo 容易忽略的点
【引申1】接口类型和 nil 作比较
接口值的零值是指动态类型和动态值都为 nil。当仅且当这两部分的值都为 nil 的情况下，这个接口值就才会被认为 接口值 == nil。


*/
type Coder interface {
	code()
}

type Gopher struct {
	name string
}

func (g Gopher) code() {
	fmt.Printf("%s is coding\n", g.name)
}

func main() {
	var c Coder
	fmt.Println(c == nil)
	fmt.Printf("c: %T, %v\n", c, c)

	var g *Gopher
	fmt.Println(g == nil)

	c = g
	fmt.Println(c == nil)
	fmt.Printf("c: %T, %v\n", c, c)

	//一开始，c 的 动态类型和动态值都为 nil，g 也为 nil，当把 g 赋值给 c 后，c 的动态类型变成了
	//*main.Gopher，仅管 c 的动态值仍为 nil，但是当 c 和 nil 作比较的时候，结果就是 false 了。

}
