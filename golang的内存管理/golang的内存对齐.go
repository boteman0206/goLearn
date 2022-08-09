package main

import (
	"fmt"
	"unsafe"
)

/*
 unsafe.Pointer 结合 unsafe 包提供的函数直接读写内存外，还引出了一个内存对齐的问题，其中 unsafe 包提供的 Alignof函数可以返回变量值在内存中的对齐字节数，
https://mp.weixin.qq.com/s/l3T5p_iw3S9nm635ezflGQ
*/
type ST1 struct {
	A byte
	B int64
	C byte
}

type ST2 struct {
	A byte
	C byte
	B int64
}

type ST3 struct {
	A uint32
	B uint64
	C struct{}
}

func main() {
	fmt.Println("ST1.A 占用的字节数是：" + fmt.Sprint(unsafe.Sizeof(ST1{}.A)))  //1
	fmt.Println("ST1.A 对齐的字节数是：" + fmt.Sprint(unsafe.Alignof(ST1{}.A))) // 1
	fmt.Println("ST1.B 占用的字节数是：" + fmt.Sprint(unsafe.Sizeof(ST1{}.B)))  // 8
	fmt.Println("ST1.B 对齐的字节数是：" + fmt.Sprint(unsafe.Alignof(ST1{}.B))) // 8
	fmt.Println("ST1.C 占用的字节数是：" + fmt.Sprint(unsafe.Sizeof(ST1{}.C)))  // 1
	fmt.Println("ST1.C 对齐的字节数是：" + fmt.Sprint(unsafe.Alignof(ST1{}.C))) // 1
	fmt.Println("ST1结构体 占用的字节数是：" + fmt.Sprint(unsafe.Sizeof(ST1{})))   // 24
	fmt.Println("ST1结构体 对齐的字节数是：" + fmt.Sprint(unsafe.Alignof(ST1{})))  // 8

	// 仅仅只是调换了一下顺序，结构体 ST1 就减少了三分之一的内存占用空间。在实际编程应用时大部分时候我们不用太过于注意内存对齐对数据结构空间的影响，
	//不过作为工程师了解内存对齐这个知识还是很重要的，它实际上是一种典型的以空间换时间的策略。
	fmt.Println("ST2.A 占用的字节数是：" + fmt.Sprint(unsafe.Sizeof(ST2{}.A)))
	fmt.Println("ST2.A 对齐的字节数是：" + fmt.Sprint(unsafe.Alignof(ST2{}.A)))
	fmt.Println("ST2.B 占用的字节数是：" + fmt.Sprint(unsafe.Sizeof(ST2{}.B)))
	fmt.Println("ST2.B 对齐的字节数是：" + fmt.Sprint(unsafe.Alignof(ST2{}.B)))
	fmt.Println("ST2.C 占用的字节数是：" + fmt.Sprint(unsafe.Sizeof(ST2{}.C)))
	fmt.Println("ST2.C 对齐的字节数是：" + fmt.Sprint(unsafe.Alignof(ST2{}.C)))
	fmt.Println("ST2结构体 占用的字节数是：" + fmt.Sprint(unsafe.Sizeof(ST2{})))
	fmt.Println("ST2结构体 对齐的字节数是：" + fmt.Sprint(unsafe.Alignof(ST2{})))

	fmt.Println("ST3.C 占用的字节数是：" + fmt.Sprint(unsafe.Sizeof(ST3{}.C)))  // 0
	fmt.Println("ST3.C 对齐的字节数是：" + fmt.Sprint(unsafe.Alignof(ST3{}.C))) // 1
	fmt.Println("ST3 结构体占用的字节数是：" + fmt.Sprint(unsafe.Sizeof(ST3{})))   // 24
}

/**
Golang 告诉我们 ST1 结构体占用的字节数是24。但是每个字段占用的字节数总共加起来确实是只有10个字节，这是怎么回事呢？
	因为字段B占用的字节数是8，内存对齐的字节数也是8，A字段所在的8个字节里不足以存放字段B，所以只好留下7个字节的空洞，在下一个 8 字节存放字段B。
又因为结构体ST1是8字节对齐的（可以理解为占的内存空间必须是8字节的倍数，且起始地址能够整除8），所以 C 字段占据了下一个8字节，但是又留下了7个字节的空洞

既然知道了 Go 编译器在对结构体进行内存对齐的时候会在字段之间留下内存空洞，那么我们把只需要 1 个字节对齐的字段 C 放在需要 8 个字节内存对齐的字段
B 前面就能让结构体 ST1 少占 8 个字节。下面我们把 ST1 的 C 字段放在 B 的前面再观察一下 ST1 结构体的大

内存对齐
操作系统在读取数据的时候并非按照我们想象的那样一个字节一个字节的去读取，而是一个字一个字的去读取。
字是用于表示其自然的数据单位，也叫machine word。字是系统用来一次性处理事务的一个固定长度。字长 / 步长 就是一个字可容纳的字节数，一般 N 位系统的字长是 (N / 8) 个字节。

当 CPU 从存储器读数据到寄存器，或者从寄存器写数据到存储器，每次 IO 的数据长度是字长。如 32 位系统访问粒度是 4 字节（bytes），64 位系统的就是 8 字节。当被访问的数据长度为 n 字节且该数据的内存地址为 n 字节对齐，那么操作系统就可以高效地一次定位到数据，无需多次读取、处理对齐运算等额外操作。



零字节类型的对齐
我们都知道 struct{} 类型占用的字节数是 0，但其实它的内存对齐数是 1，这么设定的原因为了保证当它作为结构体的末尾字段时，不会访问到其他数据结构的地址。比如像下面这个结构体 ST2
type ST2 struct {
 A uint32
 B uint64
 C struct{}
}
虽然字段 C 占用的字节数为0，但是编译器会为它补 8 个字节，这样就能保证访问字段 C 的时候不会访问到其他数据结构的内存地址。
*/
